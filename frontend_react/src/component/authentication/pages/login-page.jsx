import React from 'react'
import { FaLock, FaHome } from 'react-icons/fa'
import { IoIosMail } from "react-icons/io";
import './page-layout.css'
import { Link } from 'react-router-dom';

export const LoginPage = () => {
  return (
    <div className='container'>
        <form>
            <Link to="/" className="home-button">
                <FaHome />
            </Link>
            <h1>Login</h1>
            <div className='input-box'>
                <input 
                    type='email'
                    autoComplete="off"
                    placeholder=""
                    required
                />
                <label>Email</label>
                <IoIosMail size={20} className='icon'/>
            </div>

            <div className='input-box'>
                <input 
                    type='password'
                    placeholder=""
                    required
                />
                <label>Password</label>
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
                    Don't have an account? <Link to='/register'>Register</Link>
                </p>
            </div>
        </form>
    </div>
  )
}
