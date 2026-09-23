let accessToken: string | null = null;

export const TokenStorage = {
    getToken() {
        return accessToken;
    },

    setToken(token: string) {
        accessToken = token;
    },

    clear() {
        accessToken = null;
    },
};