import { createContext, useCallback, useEffect, useMemo, useState } from 'react'
import type { 
    BankWord,
    BankWordContextValue, 
    FunctionProps
} from '../bank.word.feature.type'
import { BankWordService } from '../services/BankWordService';

export const BankWordContext = createContext<BankWordContextValue | null>(null)

export const BankWordContextProvider = ({children}: FunctionProps) => {
    const [bankWords, setBankWords] = useState<BankWord[]>([]);
    const [loading, setLoading] = useState(false);

    const fetchBankWordList = useCallback(
        async () => {
            setLoading(true);

            try {
                const response = await BankWordService.getAllWords();

                setBankWords(response.Data);
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
            bankWords,
            loading,
        }), [bankWords, loading]
    );

    useEffect(() => {
        const callFetchBankWord = async () => {
            await fetchBankWordList();
        }

        callFetchBankWord()
    }, [fetchBankWordList]);
    
    return (
        <BankWordContext.Provider value={value}>
            {children}
        </BankWordContext.Provider>
    )
}