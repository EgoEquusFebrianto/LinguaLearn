import type { 
    GeneralResponse, 
    LoginRequest, 
    RegisterRequest 
} from "../../features/authentication/auth.featue.types"

export type LoginPageProps = {
    login: (FormData: LoginRequest) => Promise<GeneralResponse>;
}

export type RegisterPageProps = {
    register: (FormData: RegisterRequest) => Promise<GeneralResponse>;
}