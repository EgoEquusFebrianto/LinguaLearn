import React, { useState, type ReactNode } from 'react'
import { UserSidebar } from './UserSidebar';
import { UserNavbar } from './UserNavbar';
import { UserBottomNav } from './UserBottomNav';
import "./UserLayout.css";
import { useAuth } from '../../features/authentication/hooks/useAuth';

export const UserLayout = ({children}: {children: ReactNode}) => {
    const {logout} = useAuth()

    const [sidebarOpen, setSidebarOpen] = useState(false);
    const [sidebarCollapsed, setSidebarCollapsed] = useState(false);

    const handleToggleSidebar = () => {
        setSidebarOpen((prev) => !prev);
    };

    const handleClosesidebar = () => {
        setSidebarOpen(false);
    };

    const handleSidebarCollapsed = () => {
        setSidebarCollapsed((prev) => !prev);
    };

    const handleUserLogout = async () => {
        await logout();
    };

    return (
        <div 
            className={`user-layout ${
                sidebarCollapsed ? "sidebar-collapsed" : ""
            }`}
        >
            <UserSidebar
                isOpen={sidebarOpen}
                isCollapsed={sidebarCollapsed}
                onClose={handleClosesidebar}
                onCollapse={handleSidebarCollapsed}
                onLogout={handleUserLogout}
            />

            <UserNavbar onMenuClick={handleToggleSidebar}/>

            <main className='user-content'>
                {children}
            </main>

            <UserBottomNav />
        </div>
    );
};