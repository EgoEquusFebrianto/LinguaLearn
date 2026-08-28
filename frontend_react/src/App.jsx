import React from 'react'
import { AuthenticationPage } from './component/authentication/authentication'
import "./App.css"

const App = () => {
  return (
    <div className='app'>
      <div className='page-context'>
        <AuthenticationPage />
      </div>
    </div>
  )
}

export default App