import React from 'react'
import { FaBars } from 'react-icons/fa';

type UserNavbarProps = {
    onMenuClick: () => void;
}

export const UserNavbar = ({onMenuClick}: UserNavbarProps) => {
    return (
        <header className='user-navbar'>

            {/* 
                Navbar Mobile / Tablet User
            */}
            <div className='mobile-navbar'>
                <button
                    type='button'
                    className='hamburger-button'
                    onClick={onMenuClick}
                    aria-label='Open navigation menu'
                >
                    <FaBars />
                </button>

                <div className='navbar-brand'>
                    <div className='brand-logo'>
                        L
                    </div>

                    <span>LinguaLearn</span>
                </div>
            </div>

            {/* Navbar Desktop User */}
            <div className='desktop-navbar'>
                <div/>
                <div className='navbar-user-actions'>
                    <button type='button'>
                        C
                    </button>

                    <button type='button'>
                        🇮🇩 Indonesia
                    </button>

                </div>
            </div>
            
        </header>
    )
}