import React, { useState, type ReactNode } from 'react'
import { UserSidebar } from './UserSidebar';
import { UserNavbar } from './UserNavbar';
import { UserBottomNav } from './UserBottomNav';
import "./UserLayout.css";

export const UserLayout = ({children}: {children: ReactNode}) => {
    const [sidebarOpen, setSidebarOpen] = useState(false);
    const [sidebarCollapsed, setSidebarCollapsed] = useState(false);

    const handleToggleSidebar = () => {
        setSidebarOpen((prev) => !prev);
    }

    const handleClosesidebar = () => {
        setSidebarOpen(false);
    }

    const handleSidebarCollapsed = () => {
        setSidebarCollapsed((prev) => !prev);
    }

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
            />

            <UserNavbar onMenuClick={handleToggleSidebar}/>

            <main className='user-content'>
                {children}
            </main>

            <UserBottomNav />
        </div>
    );
};