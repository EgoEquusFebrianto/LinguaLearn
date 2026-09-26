export type Language = "id" | "en";

export type PreferencesContextValue = {
    isDarkMode: boolean;
    language: Language;
    toggleDarkMode: () => void;
    setLanguage: (language: Language) => void;
}

export type PreferencesContextProps = {
    children: React.ReactNode;
}