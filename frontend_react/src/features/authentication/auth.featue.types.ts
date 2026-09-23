type UserRole = "user" | "admin";

export type User = {
  id: number;
  full_name: string;
  email: string;
  role: UserRole;
};

export type LoginRequest = {
    email: string;
    password: string;
    remember_me: boolean;
}

export type RegisterRequest = {
    full_name: string;
    email: string;
    password: string;
    confirm_password: string;
}

export type LoginResponse = {
  access_token: string;
  user: User;
};

export type GeneralResponse = {
    message: string
}

export type AuthContextValue = {
    user: User | null;
    loading: boolean;
    isAuthenticated: boolean;
    login: (formData: LoginRequest) => Promise<GeneralResponse>;
    register: (formData: RegisterRequest) => Promise<GeneralResponse>;
    refresh: () => Promise<GeneralResponse | undefined>;
    logout: () => Promise<void>;
}