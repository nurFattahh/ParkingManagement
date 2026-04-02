import { createBrowserRouter, Navigate } from "react-router-dom";

import Layout from "./components/layout/Layout.jsx";
import ProtectedRoute from "./components/auth/ProtectedRoute.jsx";

import LoginPage from "./pages/auth/LoginPage.jsx";
import RegisterPage from "./pages/auth/RegisterPage.jsx";
import DashboardPage from "./pages/dashboard/DashboardPage.jsx";
import ParkingPage from "./pages/parking/ParkingPage.jsx";
import VehiclesPage from "./pages/vehicles/VehiclesPage.jsx";
import HistoryPage from "./pages/history/HistoryPage.jsx";

export const router = createBrowserRouter([
  {
    path: "/login",
    element: <LoginPage />,
  },
  {
    path: "/register",
    element: <RegisterPage />,
  },

  {
    path: "/",
    element: (
      <ProtectedRoute>
        <Layout />
      </ProtectedRoute>
    ),
    children: [
      { index: true, element: <DashboardPage /> },
      { path: "parking", element: <ParkingPage /> },
      { path: "vehicles", element: <VehiclesPage /> },
      { path: "history", element: <HistoryPage /> },
    ],
  },

  {
    path: "*",
    element: <Navigate to="/" />,
  },
]);