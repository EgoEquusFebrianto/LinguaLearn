import "./App.css"
import { BrowserRouter, Route, Routes } from 'react-router-dom'
import { AuthenticationPage } from "./page/authentication/authentication"
import { LoginPage } from './component/authentication/login-page' 
import { RegisterPage } from './component/authentication/register-page' 
import { HomePublic } from "./page/home-public/home-public"
import { HomeUser } from "./page/home-users/home-user"
import { useAuth } from "./context/authentication/hook/auth-hook"
import { Toaster } from "react-hot-toast"

const ProtectedRoute = ({isAuthenticated}) => {

  if (isAuthenticated) return <HomeUser />;

  return <HomePublic />
};

function AppContent() {
  const { isAuthenticated, login, register} = useAuth();

  return(
    <div className='app'>
      <div className='page-context'>
        <Routes>
          <Route 
            index 
            element={<ProtectedRoute isAuthenticated={isAuthenticated}/>}
          />

          <Route element={<AuthenticationPage />}>
            <Route path='login' element={<LoginPage login={login} />} />
            <Route path='register' element={<RegisterPage register={register} />} />
          </Route>
        </Routes>
      </div>
    </div>
  );
}

const App = () => {
  return (
    <BrowserRouter>
      <Toaster position="bottom-right" reverseOrder={false}/>
      <AppContent />
    </BrowserRouter>
  )
}

export default App