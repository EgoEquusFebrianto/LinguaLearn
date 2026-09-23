import { useNavigate } from "react-router-dom";

export const NavbarPublic = () => {
  const navigate = useNavigate();

  const onHandleClick = (value: string) => {
    navigate(`/${value}`)
  };

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
            onClick={() => onHandleClick("login")}
          >
            Log In
          </button>
          <button 
            className="btn-register"
              onClick={() => onHandleClick("register")}
          >
            Register
          </button>
        </div>
      </div>
    </nav>
  );
};