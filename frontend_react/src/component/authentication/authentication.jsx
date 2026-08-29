import React from 'react'
import "./authentication.css"
import { Outlet } from 'react-router-dom'

export const AuthenticationPage = () => {
  return (
    <div className='wrapper'>
        <Outlet />
    </div>
  )
}