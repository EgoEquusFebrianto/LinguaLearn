import { useContext } from "react"
import { AuthContext } from "../auth-context"

export const useAuth = () => {
    const context = useContext(AuthContext);

    if (!context) throw new Error("useAuth should used inside AuthContextProvider.");
  
    return context;
}
