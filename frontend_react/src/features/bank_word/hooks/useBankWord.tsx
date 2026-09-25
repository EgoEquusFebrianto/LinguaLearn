import { useContext } from 'react'
import { BankWordContext } from '../context/BankWordContext';

export const useBankWord = () => {
    const context = useContext(BankWordContext);

    if (!context) {
        throw new Error("useBankWord should be used in BankWordContext!");
    }

    return context;
}