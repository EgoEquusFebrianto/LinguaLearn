import React from 'react'
import { FaGamepad, FaHome, FaUser } from 'react-icons/fa'
import { NavLink } from 'react-router-dom'

export const UserBottomNav = () => {
    return (
        <nav className='bottom-navigation'>
            <NavLink
                to={"/Home"}
            >
                <FaHome/>
                <span>About</span>
            </NavLink>

            <NavLink
                to={"/games"}
            >
                <FaGamepad/>
                <span>Challenge</span>
            </NavLink>

            <NavLink
                to={"/profile"}
            >
                <FaUser/>
                <span>Profile</span>
            </NavLink>
        </nav>
    )
}
