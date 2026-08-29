import "./App.css"
import { BrowserRouter, Route, Routes } from 'react-router-dom'
import { AuthenticationPage } from "./page/authentication/authentication"
import { LoginPage } from './component/authentication/login-page' 
import { RegisterPage } from './component/authentication/register-page' 
import { HomePublic } from "./page/home-public/home-public"

function AppContent() {  
  return(
    <div className='app'>
      <div className='page-context'>
        <Routes>
          <Route index element={<HomePublic />}/>

          <Route element={<AuthenticationPage />}>
            <Route path='login' element={<LoginPage />} />
            <Route path='register' element={<RegisterPage />} />
          </Route>
        </Routes>
      </div>
    </div>
  );
}

const App = () => {
  return (
    <BrowserRouter>
      <AppContent />
    </BrowserRouter>
  )
}

export default App