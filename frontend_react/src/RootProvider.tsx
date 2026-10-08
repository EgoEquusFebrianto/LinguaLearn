import { QueryClient, QueryClientProvider } from "@tanstack/react-query"
import { AuthContextProvider } from "./features/authentication/context/AuthContext"
import { BankWordContextProvider } from "./features/bank_word/context/BankWordContext"
import { PreferencesContextProvider } from "./features/preferences/context/PreferencesContext"

export const RootProvider = ({ children }: {children: React.ReactNode}) => {
  const queryClient =new QueryClient();
  
  return (
    <QueryClientProvider client={queryClient}>
      <AuthContextProvider>
        <BankWordContextProvider>
          <PreferencesContextProvider>
            {children}
          </PreferencesContextProvider>
        </BankWordContextProvider>
      </AuthContextProvider>
    </QueryClientProvider>

  );
};