import React from 'react'
import { AuthenticationPage } from './component/authentication/authentication'
import "./App.css"
import { BrowserRouter, Route, Router, Routes } from 'react-router-dom'
import { LoginPage } from './component/authentication/pages/login-page'
import { RegisterPage } from './component/authentication/pages/register-page'
import { Home } from './page/home'

function AppContent() {
  return(
    <div className='app'>
      <div className='page-context'>
        <Routes>
          <Route index element={<Home />}/>

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