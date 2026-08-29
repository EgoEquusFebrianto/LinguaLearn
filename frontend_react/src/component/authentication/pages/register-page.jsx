import React from 'react'
import { FaLock, FaUser, FaHome } from 'react-icons/fa'
import { IoIosMail } from "react-icons/io";
import "./page-layout.css"
import { Link } from 'react-router-dom';

export const RegisterPage = () => {
  return (
    <div className='container register-container'>
        <form>
            <Link to="/" className="home-button">
                <FaHome />
            </Link>
            <h1>Register</h1>
            <div className='input-box'>
                <input 
                    type='text' 
                    placeholder=""
                    required
                />
                <label>Full Name</label>
                <FaUser className='icon'/>
            </div>

            <div className='input-box'>
                <input 
                    type='email' 
                    autoComplete='off'
                    placeholder=''
                    required
                />
                <label>Email</label>
                <IoIosMail size={20} className='icon'/>
            </div>

            <div className='input-box'>
                <input 
                    type='password' 
                    placeholder=''
                    required
                />
                <label>Password</label>
                <FaLock className='icon'/>
            </div>

            <div className='input-box last-input'>
                <input 
                    type='password' 
                    placeholder=''
                    required
                />
                <label>Confirm Password</label>
                <FaLock className='icon'/>
            </div>

            <button type='submit'>Register</button>

            <div className='register-link'>
                <p>
                    Already have an account? <Link to='/login'>Login</Link>
                </p>
            </div>
        </form>
    </div>
  )
}
