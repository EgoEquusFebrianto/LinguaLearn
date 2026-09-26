import { usePreferefences } from '../../features/preferences/hooks/usePreferefences';
import type { PrefComponentProps } from './preferences.component.types';

export const ThemeToggle = ({ showLabel }: PrefComponentProps) => {
    const {isDarkMode, toggleDarkMode} = usePreferefences();

    return (
        <label className='theme-toggle'>
            {showLabel && (
                <span className='theme-toggle-label'>
                    Dark Mode
                </span>
            )}

            <input 
                type='checkbox'
                checked={isDarkMode}
                onChange={toggleDarkMode}
                aria-label='Toggle dark mode'
            />

            <span className='theme-toggle-slider'/>
        </label>
    );
};