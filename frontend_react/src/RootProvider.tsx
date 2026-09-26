import { AuthContextProvider } from "./features/authentication/context/AuthContext"
import { BankWordContextProvider } from "./features/bank_word/context/BankWordContext"
import { PreferencesContextProvider } from "./features/preferences/context/PreferencesContext"

export const RootProvider = ({ children }: {children: React.ReactNode}) => {
  return (
    <AuthContextProvider>
      <BankWordContextProvider>
        <PreferencesContextProvider>
          {children}
        </PreferencesContextProvider>
      </BankWordContextProvider>
    </AuthContextProvider>
  );
};