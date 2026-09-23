import API from "../../api/api";
import type { 
    LoginRequest, 
    LoginResponse, 
    RegisterRequest, 
    GeneralResponse,
    User
} from "../auth.featue.types";

export const AuthService = {
    async login(request: LoginRequest): Promise<LoginResponse> {
        const response = await API.post("/auth/login", request);

        return response.data;
    },

    async register(request: RegisterRequest): Promise<GeneralResponse> {
        const response = await API.post("/auth/register", request)

        return response.data;
    },

    async getMe(): Promise<User> {
        const response = await API.get("/auth/me");

        return response.data;
    },

    async refresh(): Promise<LoginResponse> {
        const response = await API.post("/auth/refresh");
        
        return response.data;
    },

    async logout(): Promise<GeneralResponse>{
        const response = await API.post("/auth/logout");

        return response.data;
    },
}