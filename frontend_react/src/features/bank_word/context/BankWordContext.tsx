import { createContext, useCallback, useMemo, useState } from 'react'
import type { 
    BankWordContextValue, 
    FunctionProps,
    DictionaryContentQueryParams,
    DictionaryItem
} from '../bank.word.feature.type'
import { BankWordService } from '../services/BankWordService';

export const BankWordContext = createContext<BankWordContextValue | null>(null)

export const BankWordContextProvider = ({children}: FunctionProps) => {
    const [dictionaryContent, setDictionaryContent] = useState<DictionaryItem[]>([]);
    const [loading, setLoading] = useState(false);

    const fetchDictionaryContent = useCallback(
        async (query: DictionaryContentQueryParams) => {
            setLoading(true);

            try {
                const response = await BankWordService.getDictionaryContent(query);

                setDictionaryContent(prev => [...prev, ...response.data]);
            } catch (error) {
                console.error(error);

                throw error;
            } finally {
                setLoading(false);
            }
        }, []
    );

    const value = useMemo<BankWordContextValue>(
        () => ({
            dictionaryContent,
            loading,
            fetchDictionaryContent,
        }), [dictionaryContent, loading, fetchDictionaryContent]
    );

    return (
        <BankWordContext.Provider value={value}>
            {children}
        </BankWordContext.Provider>
    )
}