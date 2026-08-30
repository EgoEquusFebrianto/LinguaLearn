import React from 'react';
import "./home-user.css";
import { useAuth } from '../../context/authentication/hook/auth-hook';

export const HomeUser = () => {
    const { user, logout } = useAuth();
    
    console.log(user);

    const handleLogout = () => {
        logout();
    };

    return (
        <div className="home-user-container">
            <h1>Welcome, {user?.full_name || 'User'}! 👋</h1>
            
            <button onClick={handleLogout}>
                Logout
            </button>
        </div>
    );
};