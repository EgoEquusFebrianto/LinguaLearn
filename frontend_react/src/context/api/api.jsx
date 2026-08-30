import axios from 'axios';
import { TokenStorage } from '../../utils/auth/token-storage';

const API = axios.create({
    baseURL: import.meta.env.VITE_API_URL,
    timeout: 10000,
    headers: {

    },
    withCredentials: true,
});

/**
 * Request Interceptor
 * 
 * @import {InternalAxiosRequestConfig} from 'axios
 * 
 * @throws {Error}
 */
API.interceptors.request.use(
    /**
     * 
     * @param {InternalAxiosRequestConfig} config 
     * @returns {InternalAxiosRequestConfig}
     */
    (config) => {
        const token = TokenStorage.getToken();

        if (token) {
            config.headers.Authorization = `Bearer ${token}`;
        }

        return config;
    },

    (error) => {
        return Promise.reject(error);
    },
);

/**
 * Response Interceptor
 * 
 * @import { AxiosResponse, AxiosError } from 'axios
 * 
 * @throws {Error}
 */
API.interceptors.response.use(
    /**
     * 
     * @param {AxiosResponse} response 
     * @returns {AxiosResponse}
     */
    (response) => {
        return response
    },

    /**
     * 
     * @param {AxiosError} error 
     * @returns {Promise<AxiosResponse>}
     */
    async (error) => {
        const originalRequest = error.config;

        const publicEndpoints = ['/auth/login', '/auth/register'];
        const isPublicEndpoint = publicEndpoints.some(url => 
            originalRequest.url?.includes(url)
        );

        if (
            error.response?.status === 401 &&
            !originalRequest._retry &&
            !originalRequest.url.includes("/auth/refresh") &&
            !isPublicEndpoint
        ) {
            originalRequest._retry = true;

            try {
                const response = await API.post("/auth/refresh");
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
    }
);

/**
 * @type {import('axios').AxiosInstance}
 */
export default API