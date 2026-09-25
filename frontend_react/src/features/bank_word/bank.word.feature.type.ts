
export type Sense = {
    Translations: string[];
    Synonyms: string[];
    Type: string;
    Description: string
}

export type BankWord = {
    ID: string;
    WordUUID: string;
    Word: string;
    Senses: Sense[];
}

export type BankWordSearchResponse = {
    Data: BankWord[];
    Page: number;
    Limit: number;
    Total: number;
    TotalPages: number;
}

export type FunctionProps = {
    children: React.ReactNode;
}

export type BankWordContextValue = {
    bankWords: BankWord[];
    loading: boolean;
}