import { createContext, useCallback, useEffect, useState } from 'react'
import API from '../api/api';
import { TokenStorage } from '../../utils/auth/token-storage';

export const AuthContext = createContext(null);

export const AuthContextProvider = ({children}) => {
    const [user, setUser] = useState(null);
    const [loading, setLoading] = useState(false);

    const login = async (formData) => {
        try {
            const response = await API.post(
                "/auth/login",
                formData,
            );

            const {
                access_token,
                user,
            } = response.data;

            TokenStorage.setToken(access_token);
            setUser(user);

        return {
            success: true,
            status: response.status,
            message: "Login Successful.",
        };

        } catch (error) {
            console.log("Failed Login: ", error)

            throw error;
        }
    };

    const refresh = useCallback(
        async () => {
            setLoading(true);

            try {
                const response = await API.post("/auth/refresh");
                const {access_token, user } = response.data;

                TokenStorage.setToken(access_token);
                setUser(user);

                return {
                    success: true,
                    status: response.status,
                    message: "Refresh Access Successful.",
                };
            } catch (error) {
                TokenStorage.clear();
                setUser(null);

                return;
            }
        }, []
    );

    const register = async (formData) => {
        try {
            await API.post("/auth/register", formData);

        } catch (error) {
            console.error("Failed Register", error)

            throw error;
        }
    };

    const logout = async () => {
        try {
            await API.post("/auth/logout");
        } finally {
            TokenStorage.clear();
            setUser(null);
        }
    }

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

    const value = {
        user,
        loading,
        login,
        refresh,
        register,
        logout,
        isAuthenticated: !!user,
    };

    return (
        <AuthContext.Provider value={value}>
            {children}
        </AuthContext.Provider>
    );
}