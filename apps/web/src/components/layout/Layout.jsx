import { Outlet } from "react-router-dom";
import Sidebar from "./Sidebar";
import Navbar from "./Navbar";

export default function Layout() {
  return (
    <div style={{ display: "flex", width: "100vw", height: "100vh" }}>
      <Sidebar />
        <div style={{ padding: "20px", width: "100%" }}>
          <Outlet />
        </div>
      </div>

  );
}