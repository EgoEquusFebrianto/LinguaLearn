import { BrowserRouter, Route, Routes } from 'react-router-dom';
import { AuthenticationPage } from "./page/authentication/AuthenticationOutlet";
import { LoginPage } from './page/authentication/components/LoginPage';
import { RegisterPage } from './page/authentication/components/RegisterPage';
import { HomePublic } from "./page/home_public/HomePublic";
import { HomeUser } from "./page/home_authorized_user/HomeUser";
import { useAuth } from "./features/authentication/hooks/useAuth";
import { Toaster } from "react-hot-toast";
import "./App.css";

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