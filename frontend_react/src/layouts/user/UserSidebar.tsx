import { 
    FaBook, 
    FaGamepad, 
    FaHome, 
    FaInfoCircle, 
    FaSignOutAlt, 
    FaUser, 
    FaUserEdit 
} from 'react-icons/fa';
import { NavLink } from 'react-router-dom';
import { Preferences } from '../../components/preferences/Preferences';

type UserSidebarProps = {
    isOpen: boolean;
    isCollapsed: boolean;
    onClose: () => void;
    onCollapse: () => void;
}

export const UserSidebar = ({
    isOpen, 
    isCollapsed, 
    onClose, 
    onCollapse
}: UserSidebarProps
) => {
    return (
        <>
            <aside
                className={`user-sidebar ${
                    isOpen ? "sidebar-open" : ""
                } ${
                    isCollapsed ? "sidebar-is-collapsed" : ""
                }`}
            >
                <div className='sidebar-brand'>
                    <div className='brand-logo'>
                        L
                    </div>

                    {!isCollapsed && (<span>LinguaLearn</span>)}
                </div>

                <nav className='sidebar-navigation'>
                    <NavLink to={"/home"}>
                        <FaHome />
                        {!isCollapsed && (<span>Beranda</span>)}
                    </NavLink>

                    <div className='sidebar-section'>
                        {!isCollapsed && (
                            <span className='sidebar-section-title'>
                                FEATURE
                            </span>
                        )}

                        <NavLink
                            to={"/dictionary"}
                        >
                            <FaBook/>
                            {!isCollapsed && (<span>Kamus</span>)}
                        </NavLink>

                        <NavLink 
                            to={"/games"}
                        >
                            <FaGamepad />
                            {!isCollapsed  && (<span>Games</span>)}
                        </NavLink>
                    </div>

                    <div className='sidebar-section sidebar-account'>
                        {!isCollapsed && (
                            <span className='sidebar-section-title'>
                                ACCOUNT
                            </span>
                        )}

                        <NavLink
                            to={"/profile"}
                        >
                            <FaUser/>
                            {!isCollapsed && (<span>Profile</span>)}
                        </NavLink>

                        <NavLink
                            to={"/profile/edit"}
                        >
                            <FaUserEdit/>
                            {!isCollapsed && (<span>Edit Profile</span>)}
                        </NavLink>
                    </div>

                    <div className="sidebar-section sidebar-preferences">
                        {!isCollapsed && (
                            <span className="sidebar-section-title">
                                PREFERENCES
                            </span>
                        )}

                        <Preferences
                            showThemeLabel={!isCollapsed}
                            showLanguageLabel={!isCollapsed}
                        />
                    </div>
                    <NavLink
                        to={"/about"}
                    >
                        <FaInfoCircle/>
                        {!isCollapsed && (<span>About</span>)}
                    </NavLink>
                </nav>

                <div className='sidebar-bottom'>                    
                    <button
                        type='button'
                        className='logout-button'
                    >
                        <FaSignOutAlt />
                        {!isCollapsed && (<span>Logout</span>)}
                    </button>

                    <button
                        type='button'
                        className='collapse-button'
                        onClick={onCollapse}
                    >
                        {isCollapsed ? ">>" : "Collapse <<"}
                    </button>
                </div>
            </aside>

            {isOpen && (
                <div
                    className='sidebar-overlay'
                    onClick={onClose}
                />
            )}
        </>
    )
}