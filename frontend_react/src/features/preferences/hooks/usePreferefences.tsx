import { useContext } from 'react'
import { PreferencesContext } from '../context/PreferencesContext';

export const usePreferefences = () => {
    const context = useContext(PreferencesContext);

    if (!context) {
        throw new Error("usePreferences should be used in PreferencesContextProvider!")
    }

    return context;
};