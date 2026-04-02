import { Link, useLocation, useNavigate } from "react-router-dom";
import "./sidebar.css";

export default function Sidebar() {
  const location = useLocation();
  const navigate = useNavigate();

  const username = localStorage.getItem("username");

  const logout = () => {
    localStorage.removeItem("token");
    navigate("/login");
  };

  const menu = [
    { name: "Parkir Kendaraan", path: "/parking", icon: "🚗" },
    { name: "Data Kendaraan", path: "/vehicles", icon: "📋" },
    { name: "Riwayat", path: "/history", icon: "📊" },
    { name: "Settings", path: "/settings", icon: "⚙️" },
  ];

  return (
    <aside className="sidebar">

      <div className="logo">
        Parking System
      </div>

      <div className="account">
        <img
          src="https://i.pravatar.cc/100"
          alt="avatar"
          className="avatar"
        />

        <div className="account-info">
          <div className="username">{username || "User"}</div>
          <button className="logout-btn" onClick={logout}>
            Logout
          </button>
        </div>
      </div>

      <nav className="menu">
        {menu.map((item) => (
          <Link
            key={item.path}
            to={item.path}
            className={`menu-item ${
              location.pathname === item.path ? "active" : ""
            }`}
          >
            <span className="icon">{item.icon}</span>
            <span>{item.name}</span>
          </Link>
        ))}
      </nav>

    </aside>
  );
}