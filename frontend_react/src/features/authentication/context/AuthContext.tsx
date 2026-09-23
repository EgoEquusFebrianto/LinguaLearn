import { createContext, useCallback, useEffect, useMemo, useState } from 'react'
import { TokenStorage } from '../../../utils/auth/tokenStorage';
import type { 
    User, 
    LoginRequest, 
    RegisterRequest, 
    AuthContextValue,
    GeneralResponse
} from '../auth.featue.types';
import { AuthService } from '../services/AuthService';

export const AuthContext = createContext<AuthContextValue | null>(null);

export const AuthContextProvider = ({children}) => {
    const [user, setUser] = useState<User | null>(null);
    const [loading, setLoading] = useState(false);

    const login = useCallback(
        async (formData: LoginRequest) => {
            try {
                const response = await AuthService.login(formData);

                const {
                    access_token,
                    user,
                } = response;

                TokenStorage.setToken(access_token);
                setUser(user);

            return {
                message: "Login Successful.",
            };

            } catch (error) {
                console.log("Failed Login: ", error)

                throw error;
            }
        }, []
    );

    const refresh = useCallback(
        async () => {
            setLoading(true);

            try {
                const response = await AuthService.refresh()
                ;
                const {access_token, user } = response;

                TokenStorage.setToken(access_token);
                setUser(user);

                return {
                    message: "Refresh Access Successful.",
                };
            } catch (error) {
                TokenStorage.clear();
                setUser(null);

                throw error;
            }
        }, []
    );

    const register = useCallback(
        async (formData: RegisterRequest): Promise<GeneralResponse> => {
            try {
                const response = await AuthService.register(formData);
                
                return response;
            } catch (error) {
                console.error("Failed Register:", error);
                throw error;
            }
        }, []
    );

    const logout = useCallback(
        async () => {
            try {
                await AuthService.logout();
            } finally {
                TokenStorage.clear();
                setUser(null);
            }
        }, []
    );

    useEffect(() => {
        const initializeAuth = async () => {
            try {
                await refresh();
            } finally{
                setLoading(false);
            }        
        };

        initializeAuth();

    }, [refresh]);

    const value = useMemo<AuthContextValue>(
        () => ({
            user,
            loading,
            isAuthenticated: !!user,
            login,
            register,
            refresh,
            logout,
        }), [user, loading, login, register, refresh, logout]
    );

    return (
        <AuthContext.Provider value={value}>
            {children}
        </AuthContext.Provider>
    );
}