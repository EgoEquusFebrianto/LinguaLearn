import { FaLock, FaHome } from 'react-icons/fa'
import { IoIosMail } from "react-icons/io";
import { Link, useNavigate } from 'react-router-dom';
import { useState } from 'react';
import { showErrorToast, showInfoToast } from '../../../utils/toastHelper';
import './PageLayout.css'

export const LoginPage = ({login}) => {
    const navigate = useNavigate();
    const [loading, setLoading] = useState(false)
    const defaultForm = {
        email: '',
        password: '',
        remember_me: false,
    }

    const [formData, setFormData] = useState(defaultForm);

    const handleChange = (event) => {
        const { name, value, type, checked } = event.target;

        setFormData((previous) => ({
            ...previous,
            [name]: type === "checkbox"
                ? checked
                : value,
        }));
    };

    const handleSubmit = async(event) => {
        event.preventDefault();
        setLoading(true);

        try {
            const response = await login(formData);

            showInfoToast(response.message);
            setFormData(defaultForm);
            navigate("/");
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

    return (
        <div className='container'>
            <form onSubmit={handleSubmit}>
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
                        value={formData.email}
                        onChange={handleChange}
                        name='email'
                    />
                    <label>Email</label>
                    <IoIosMail size={20} className='icon'/>
                </div>

                <div className='input-box'>
                    <input 
                        type='password'
                        placeholder=""
                        required
                        value={formData.password}
                        onChange={handleChange}
                        name='password'
                    />
                    <label>Password</label>
                    <FaLock className='icon'/>
                </div>

                <div className='remember-forgot'>
                    <label>
                        <input 
                            type='checkbox'
                            checked={formData.remember_me}
                            onChange={handleChange}
                            name='remember_me'
                        />
                        Rembember me
                    </label>
                    <a href='#'>Forgot passoword</a>
                </div>

                <button
                    type="submit"
                    disabled={loading}
                >
                    {loading ? "Logging in..." : "Login"}
                </button>

                <div className='register-link'>
                    <p>
                        Don't have an account? <Link to='/register'>Register</Link>
                    </p>
                </div>
            </form>
        </div>
    );
}