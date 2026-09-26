import { usePreferefences } from "../../features/preferences/hooks/usePreferefences";
import type { Language } from "../../features/preferences/preferences.types";
import type { PrefComponentProps } from "./preferences.component.types";

export const LanguageSelector = ({ showLabel }: PrefComponentProps) => {
    const {language, setLanguage} = usePreferefences();

    const handleChange = (
        event: React.ChangeEvent<HTMLSelectElement>
    ) => {
        setLanguage(event.target.value as Language);
    };

    return (
        <label className="language-selector">
            {showLabel && (
                <span className="language-selector-label">
                    Language
                </span>
            )}
            <select
                value={language}
                onChange={handleChange}
                aria-label="Select language"
            >
                <option value={"id"}>
                    🇮🇩 Indonesia
                </option>

                <option value={"en"}>
                    🇬🇧 English
                </option>
            </select>
        </label>
    );
};