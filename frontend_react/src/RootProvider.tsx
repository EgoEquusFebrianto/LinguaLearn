import { AuthContextProvider } from "./context/authentication/auth-context"

export const RootProvider = ({ children }) => {
  return (
    <AuthContextProvider>
        {children}
    </AuthContextProvider>
  )
}
