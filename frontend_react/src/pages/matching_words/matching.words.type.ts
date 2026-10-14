export type Sense = {
    translations: string[];
    synonyms: string[];
    type: string;
    description: string;
}

export type BankWords = {
    id: string;
    word_uuid: string;
    word: string;
    senses: Sense[]
}

export type MatchingWords = {
    word_id: string;
    word_uuid: string;
    word: string;
    translations: string[];
};

export type MatchFeedback = {
    type: "correct" | "wrong";
    leftIds: string[];
    rightIds: string[];
}

export type MatchingGameState = {
    words: MatchingWords[];
    rightWords: MatchingWords[];
    queue: MatchingWords[];
    totalPairs: number;
}