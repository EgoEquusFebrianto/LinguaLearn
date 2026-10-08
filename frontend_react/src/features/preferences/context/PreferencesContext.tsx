import { createContext, useEffect, useState } from "react"
import type { 
    Language,  
    PreferencesContextProps,  
    PreferencesContextValue 
} from "../preferences.types";

const DARK_MODE_STORAGE_KEY = "lingualearn-dark-mode";
const LANGUAGE_STORAGE_KEY = "lingualearn-language"

export const PreferencesContext = createContext<PreferencesContextValue | null>(null);

export const PreferencesContextProvider = (
    {children}: PreferencesContextProps
) => {
    const [isDarkMode, setIsDarkMode] = useState<boolean>(() => {
        const storedValue = localStorage.getItem(DARK_MODE_STORAGE_KEY);

        return storedValue === "true";
    });

    const [language, setLanguageState] = useState<Language>(() => {
        const storedValue = localStorage.getItem(LANGUAGE_STORAGE_KEY);

        if (storedValue === "id" || storedValue === "en") {
            return storedValue;
        }

        return "id";
    });

    useEffect(() => {
        localStorage.setItem(LANGUAGE_STORAGE_KEY, language);
    }, [language]);

    useEffect(() => {
        document.documentElement.classList.toggle("dark", isDarkMode);

        localStorage.setItem(DARK_MODE_STORAGE_KEY, String(isDarkMode));
    }, [isDarkMode]);

    const toggleDarkMode = () => {
        setIsDarkMode((current) => !current);
    };

    const setLanguage = (language: Language) => {
        setLanguageState(language);
    };

    const contextValue: PreferencesContextValue = {
        isDarkMode,
        toggleDarkMode,
        language,
        setLanguage,
    };

    return (
        <PreferencesContext.Provider value={contextValue}>
            {children}
        </PreferencesContext.Provider>
    )
}