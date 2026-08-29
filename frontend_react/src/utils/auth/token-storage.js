let accessToken = null;

export const TokenStorage = {
    getToken() {
        return accessToken;
    },

    setToken(token) {
        accessToken = token;
    },

    clear() {
        accessToken = null;
    },
};