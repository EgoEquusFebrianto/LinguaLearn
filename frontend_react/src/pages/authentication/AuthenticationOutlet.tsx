import React from 'react'
import { Outlet } from 'react-router-dom'
import "./AuthtenticationOutlet.css"

export const AuthenticationPage = () => {
  return (
    <div className='wrapper'>
        <Outlet />
    </div>
  )
}