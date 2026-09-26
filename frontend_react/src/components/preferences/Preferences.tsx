import { ThemeToggle } from './ThemeToggle';
import { LanguageSelector } from './LanguageSelector';
import type { PreferencesProps } from './preferences.component.types';
import "./Preferences.css"

export const Preferences = ({
    showThemeLabel = true,
    showLanguageLabel = true,
}: PreferencesProps) => {
    return (
        <div className='preferences'>
            <ThemeToggle showLabel={showThemeLabel}/>
            <LanguageSelector showLabel={showLanguageLabel}/>
        </div>
    );
};
