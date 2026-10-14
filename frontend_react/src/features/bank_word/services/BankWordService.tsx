import API from "../../api/api";
import type {
  BankWords,
  MatchingWords,
  BankWordSearchResponse, 
  DictionaryContentQueryParams, 
  DictionaryItem,
  DictionaryContentResponse
} from "../bank.word.feature.type";

export const BankWordService = {
  // async getAllWords(): Promise<BankWordSearchResponse> {
  //   const response = await API.get("/dictionary");

  //   return response.data;
  // },

  async getDictionaryContent(
    query: DictionaryContentQueryParams
  ): Promise<DictionaryContentResponse> {
    if (query.page != null && query.page < 1) {
      return {
        data: [],
        page: query.page,
        limit: 0,
        total: 0,
        total_pages:0,
      };
    }
    const response = await API.get("/dictionary", {params: query});
    const payload: BankWordSearchResponse = response.data;
    
    const res: DictionaryItem[] = payload.data.flatMap(
      ({id, word, senses}) => senses.map(
          (sense, idx) => ({
              senseId: `${id}-${idx}`,
              id,
              word,
              type: sense.type,
              translations: sense.translations ?? [],
              synonyms: sense.synonyms ?? [],
              description: sense.description ?? "",
          })
      )
    );

    // eslint-disable-next-line @typescript-eslint/no-unused-vars
    const {data, ...paginationInfo} = payload;

    return {...paginationInfo, data: res};
  },

  /**
   * Mengubah data RESPONSE API menjadi pasangan kata permainan.
   * Satu sense menghasilkan satu pasangan.
   */
  async getMatchingWordsContent(): Promise<MatchingWords[]> {
    const response = await API.get("dictionary/random");
    const payload: BankWords[] = response.data;

    return payload.flatMap(
      ({id, word_uuid, word, senses}) => senses.map(
        (sense, idx) => ({
          word_id: `${id}-${idx}`,
          word_uuid,
          word,
          translations: sense.translations ?? [],
        }
      )
    ));
  },
}