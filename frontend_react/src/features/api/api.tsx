import axios, {
    AxiosError,
    type AxiosResponse,
    type InternalAxiosRequestConfig,
} from 'axios';
import { TokenStorage } from '../../utils/auth/tokenStorage';

type RetryableRequest = InternalAxiosRequestConfig & {
    _retry?: boolean;
};

const API = axios.create({
    baseURL: import.meta.env.VITE_API_URL,
    timeout: 10000,
    withCredentials: true,
});

/**
 * Request interceptor
 */
API.interceptors.request.use(
    (config: InternalAxiosRequestConfig) => {
        const token = TokenStorage.getToken();
        if (token) {
            config.headers.Authorization = `Bearer ${token}`;
        }
        return config;
    },
    (error) => Promise.reject(error),
);

/**
 * Response interceptor
 */
API.interceptors.response.use(
    (response: AxiosResponse) => response,

    async (error: AxiosError) => {
        const originalRequest = error.config as RetryableRequest | undefined;

        if (!originalRequest) {
            return Promise.reject(error);
        }

        const publicEndpoints = ['/auth/login', '/auth/register'];
        const isPublicEndpoint = publicEndpoints.some((url) =>
            originalRequest.url?.includes(url),
        );

        if (
            error.response?.status === 401 &&
            !originalRequest._retry &&
            !originalRequest.url?.includes('/auth/refresh') &&
            !isPublicEndpoint
        ) {
            originalRequest._retry = true;

            try {
                const token = TokenStorage.getToken();
                if (token) return Promise.reject(error);

                const response = await API.post('/auth/refresh');
                const newAccessToken = response.data.access_token;

                TokenStorage.setToken(newAccessToken);
                originalRequest.headers.Authorization = `Bearer ${newAccessToken}`;

                return API(originalRequest);
            } catch (refreshError) {
                TokenStorage.clear();
                return Promise.reject(refreshError);
            }
        }

        return Promise.reject(error);
    },
);

export default API;