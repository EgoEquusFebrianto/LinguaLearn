// src/components/Navbar.jsx
import React, { useState } from 'react';

export const NavbarPublic = () => {
  const [isHovered, setIsHovered] = useState(false);

  return (
    <nav className="navbar">
      <div className="navbar-container">
        <div className="navbar-logo">
          <span className="logo-icon">📚</span>
          <span className="logo-text">LinguaLearn</span>
        </div>

        <div className="navbar-actions">
          <button 
            className="btn-login"
            onMouseEnter={() => setIsHovered(true)}
            onMouseLeave={() => setIsHovered(false)}
          >
            Log In
          </button>
          <button 
            className="btn-register"
            onMouseEnter={() => setIsHovered(true)}
            onMouseLeave={() => setIsHovered(false)}
          >
            Register
          </button>
        </div>
      </div>
    </nav>
  );
};