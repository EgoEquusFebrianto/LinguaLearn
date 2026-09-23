import { useContext } from "react"
import { AuthContext } from "../context/AuthContext"

export const useAuth = () => {
    const context = useContext(AuthContext);

    if (!context) throw new Error("useAuth should used inside AuthContextProvider.");
  
    return context;
}
