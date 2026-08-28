import React from 'react'
import { FaLock, FaUser } from 'react-icons/fa'
import './login-page.css'

export const LoginPage = () => {
  return (
    <div className='container'>
        <form>
            <h1>Login</h1>
            <div className='input-box'>
                <input 
                    type='text' 
                    placeholder='Username'
                    required
                />
                <FaUser className='icon'/>
            </div>

            <div className='input-box'>
                <input 
                    type='password' 
                    placeholder='Password'
                    required
                />
                <FaLock className='icon'/>
            </div>

            <div className='remember-forgot'>
                <label>
                    <input type='checkbox'/>
                    Rembember me
                </label>
                <a href='#'>Forgot passoword</a>
            </div>

            <button type='submit'>Login</button>

            <div className='register-link'>
                <p>
                    Don't have an account? <a href='#'>Register</a>
                </p>
            </div>
        </form>
    </div>
  )
}
