import { useNavigate } from "react-router-dom";

export default function Navbar() {
  const navigate = useNavigate();

  const handleLogout = () => {
    const confirm = window.confirm("Yakin ingin logout?");

    if (confirm) {
      localStorage.removeItem("token");
      navigate("/login");
    }
  };

  return (
    <div className="navbar">
      <h3>Parking System</h3>

      <button onClick={handleLogout}>
        Logout
      </button>
    </div>
  );
}