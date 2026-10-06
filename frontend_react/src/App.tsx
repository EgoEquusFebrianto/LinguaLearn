import { BrowserRouter, Route, Routes } from 'react-router-dom';
import { AuthenticationPage } from "./pages/authentication/AuthenticationOutlet";
import { LoginPage } from './pages/authentication/components/LoginPage';
import { RegisterPage } from './pages/authentication/components/RegisterPage';
import { HomePublic } from "./pages/home_public/HomePublic";
import { HomeUser } from "./pages/home_authorized_user/HomeUser";
import { useAuth } from "./features/authentication/hooks/useAuth";
import { Toaster } from "react-hot-toast";
import { UserLayout } from './layouts/user/UserLayout';
import { UserDictionary } from './pages/dictionary/UserDictionary';
import "./App.css";

type ProtectedRouteProps = {
  isAuthenticated: boolean;
};

const ProtectedRoute = ({isAuthenticated}: ProtectedRouteProps) => {
  if (isAuthenticated) {
    return (
      <UserLayout>
        <HomeUser />
      </UserLayout>
    )
  };

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
            // element={<ProtectedRoute isAuthenticated={true}/>}
          />

          <Route element={<AuthenticationPage />}>
            <Route path='login' element={<LoginPage login={login} />} />
            <Route path='register' element={<RegisterPage register={register} />} />
          </Route>

          <Route path='dictionary' element={<UserDictionary />}/>
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