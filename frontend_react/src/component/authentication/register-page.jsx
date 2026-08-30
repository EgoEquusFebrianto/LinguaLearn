import { useState } from 'react'
import { FaLock, FaUser, FaHome } from 'react-icons/fa'
import { IoIosMail } from "react-icons/io";
import "./page-layout.css"
import { Link, useNavigate } from 'react-router-dom';
import { showErrorToast, showSuccessToast } from '../../utils/toast-utils';

export const RegisterPage = ({register}) => {
    const navigate = useNavigate();
    const [loading, setLoading] = useState(false)
    const defaultKey = {
        full_name: '',
        email: '',
        password: '',
        confirm_password: '',
    };
    
    const [formData, setFormData] = useState(defaultKey);

    const handleChange = (event) => {
        const { name, value } = event.target;

        setFormData((previous) => ({
            ...previous,
            [name]: value,
        }));
    };

    const handleSubmit = async(event) => {
        event.preventDefault();
        setLoading(true);

        try {
            await register(formData);

            showSuccessToast("Register Successfully, Please Login.");
            setFormData(defaultKey);
            navigate("/login")
        } catch (error) {
            console.error("Failed Login", error)
            const errorMessage = error.response?.data?.error || 
                                error.response?.data?.message ||
                                'Login failed. Please try again.';
                                
            showErrorToast(errorMessage);
        } finally {
            setLoading(false);
        }
    };

    console.log(formData)

    return (
        <div className='container register-container'>
            <form onSubmit={handleSubmit}>
                <Link to="/" className="home-button">
                    <FaHome />
                </Link>
                <h1>Register</h1>
                <div className='input-box'>
                    <input 
                        type='text' 
                        placeholder=""
                        required
                        name='full_name'
                        value={formData.full_name}
                        onChange={handleChange}
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
                        name='email'
                        value={formData.email}
                        onChange={handleChange}
                    />
                    <label>Email</label>
                    <IoIosMail size={20} className='icon'/>
                </div>

                <div className='input-box'>
                    <input 
                        type='password' 
                        placeholder=''
                        required
                        name='password'
                        value={formData.password}
                        onChange={handleChange}
                    />
                    <label>Password</label>
                    <FaLock className='icon'/>
                </div>

                <div className='input-box last-input'>
                    <input 
                        type='password' 
                        placeholder=''
                        required
                        name='confirm_password'
                        value={formData.confirm_password}
                        onChange={handleChange}
                    />
                    <label>Confirm Password</label>
                    <FaLock className='icon'/>
                </div>

                <button 
                    type='submit'
                    disabled={loading}
                >
                    {loading ? "Registering user..." : "Register"}
                </button>

                <div className='register-link'>
                    <p>
                        Already have an account? <Link to='/login'>Login</Link>
                    </p>
                </div>
            </form>
        </div>
    );
}