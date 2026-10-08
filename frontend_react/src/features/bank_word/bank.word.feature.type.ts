export type Sense = {
    translations: string[];
    synonyms: string[];
    type: string;
    description: string;
};

export type BankWord = {
    id: string;
    word_uuid: string;
    word: string;
    senses: Sense[];
};

export type BankWordSearchResponse = {
    data: BankWord[];
    page: number;
    limit: number;
    total: number;
    total_pages: number;
};

export type FunctionProps = {
    children: React.ReactNode;
};

export type DictionaryItem = {
    id: string;
    senseId: string;
    word: string;
    type: string;
    translations: string[];
    synonyms: string[];
    description: string;
};

export type DictionaryContentQueryParams = {
    page?: number
}

export type DictionaryContentResponse = { 
    data: DictionaryItem[];
    page: number;
    limit: number;
    total: number;
    total_pages: number;
};

export type BankWordContextValue = {
    dictionaryContent: DictionaryItem[];
    loading: boolean;
    fetchDictionaryContent: (query: DictionaryContentQueryParams) => Promise<void>;
};