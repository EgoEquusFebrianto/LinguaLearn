import { AuthContextProvider } from "./features/authentication/context/AuthContext"

export const RootProvider = ({ children }: {children: React.ReactNode}) => {
  return (
    <AuthContextProvider>
        {children}
    </AuthContextProvider>
  )
}