import { useInfiniteQuery } from "@tanstack/react-query"
import { BankWordService } from "../services/BankWordService";

export const useInfiniteDictionary = () => {
    const query = useInfiniteQuery({
        queryKey: ["dictionary"],
        queryFn: async ({pageParam}) => BankWordService.getDictionaryContent({
            page: pageParam
        }),
        initialPageParam: 1,
        getNextPageParam: (lastPage) => {
            if (lastPage.page < lastPage.total_pages) {
                return lastPage.page + 1;
            }

            return undefined;
        }
    });
    
    const dictionaryContent = query.data?.pages.flatMap(
        (page) => page.data
    ) ?? [];
    
    return {...query, dictionaryContent};
};