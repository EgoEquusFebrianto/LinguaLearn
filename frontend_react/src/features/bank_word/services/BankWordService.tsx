import API from "../../api/api";
import type { BankWordSearchResponse } from "../bank.word.feature.type";

export const BankWordService = {
  async getAllWords(): Promise<BankWordSearchResponse> {
    const response = await API.get("/dictionary");

    return response.data;
  },

  
}