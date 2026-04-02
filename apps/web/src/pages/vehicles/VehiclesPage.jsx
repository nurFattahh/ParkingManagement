import { useEffect, useState } from "react";
import {
  getAllVehicles,
  getVehicleByLicensePlate,
  deleteVehicle
} from "../../services/vehicleService";
import "./VehiclePage.css";

export default function VehiclesPage() {
  const [vehicles, setVehicles] = useState([]);
  const [search, setSearch] = useState(""); 
  const [loading, setLoading] = useState(true);
  const [now, setNow] = useState(Date.now());

  useEffect(() => {
    fetchVehicles();
  }, []);

  // 🔥 interval realtime
  useEffect(() => {
    const interval = setInterval(() => {
      setNow(Date.now());
    }, 1000);

    return () => clearInterval(interval);
  }, []);

  const fetchVehicles = async () => {
    try {
      setLoading(true);
      const data = await getAllVehicles();
      setVehicles(data);
    } catch (error) {
      console.error(error);
      alert("Gagal mengambil data kendaraan");
    } finally {
      setLoading(false);
    }
  };

  const handleSearch = async () => {
    if (!search.trim()) {
      fetchVehicles();
      return;
    }

    try {
      setLoading(true);
      const data = await getVehicleByLicensePlate(search);

      // ⚠️ kalau nanti backend sudah LIKE (array)
      if (Array.isArray(data)) {
        setVehicles(data);
      } else {
        setVehicles([data]);
      }
    } catch (err) {
      console.error(err);
      setVehicles([]);
    } finally {
      setLoading(false);
    }
  };

  // 🔥 format waktu
  const formatDateTime = (isoString) => {
    const date = new Date(isoString);

    const day = String(date.getDate()).padStart(2, "0");
    const month = String(date.getMonth() + 1).padStart(2, "0");
    const year = date.getFullYear();

    const hours = String(date.getHours()).padStart(2, "0");
    const minutes = String(date.getMinutes()).padStart(2, "0");

    return `${day}/${month}/${year} ${hours}:${minutes}`;
  };

  // 🔥 duration realtime
  const formatDuration = (entryTime) => {
    const start = new Date(entryTime).getTime();
    const diff = Math.max(0, now - start);

    const seconds = Math.floor(diff / 1000);
    const minutes = Math.floor(seconds / 60);
    const hours = Math.floor(minutes / 60);

    const remainingMinutes = minutes % 60;
    const remainingSeconds = seconds % 60;

    return `${hours}j ${remainingMinutes}m ${remainingSeconds}d`;
  };

  // 🔥 tarif realtime (estimasi)
  const calculateTariff = (entryTime, vehicleType) => {
    const start = new Date(entryTime).getTime();
    const diffHours = Math.ceil((now - start) / (1000 * 60 * 60));

    if (vehicleType === "motor") return diffHours * 2000;
    if (vehicleType === "mobil") return diffHours * 5000;

    return 0;
  };

  return (
    <div className="vehicles-container">
      <div className="header">
        <h2>Data Kendaraan</h2>

        <div className="search-bar">
          <input
            type="text"
            placeholder="Cari berdasarkan plat nomor..."
            value={search}
            onChange={(e) => setSearch(e.target.value)}
            onKeyDown={(e) => {
              if (e.key === "Enter") handleSearch();
            }}
          />

          <button className="btn-search" onClick={handleSearch}>
            Search
          </button>
        </div>
      </div>

      {loading ? (
        <p>Loading...</p>
      ) : vehicles.length === 0 ? (
        <p>Tidak ada data kendaraan</p>
      ) : (
        <div className="table-wrapper">
          <table className="vehicles-table">
            <thead>
              <tr>
                <th>No</th>
                <th>Plat Nomor</th>
                <th>Tipe Kendaraan</th>
                <th>Waktu Masuk</th>
                <th>Durasi</th>
                <th>Tarif</th>
                <th>Action</th>
              </tr>
            </thead>

            <tbody>
              {vehicles.map((v, index) => (
                <tr key={v.ID}>
                  <td>{index + 1}</td>
                  <td>{v.LicensePlate}</td>
                  <td>{v.VehicleType}</td>
                  <td>{formatDateTime(v.EntryTime)}</td>

                  {/* 🔥 realtime */}
                  <td>{formatDuration(v.EntryTime)}</td>

                  {/* 🔥 estimasi */}
                  <td>
                    Rp {calculateTariff(v.EntryTime, v.VehicleType)}
                  </td>

                  <td className="action-cell">
                    <button
                      className="btn-delete"
                      onClick={async () => {
                        const confirm = window.confirm(
                          "Yakin ingin menghapus kendaraan ini?"
                        );

                        if (confirm) {
                          try {
                            await deleteVehicle(v.ID);
                            alert("Kendaraan berhasil dihapus");
                            fetchVehicles();
                          } catch (error) {
                            console.error(error);
                            alert("Gagal menghapus kendaraan");
                          }
                        }
                      }}
                    >
                      Delete
                    </button>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      )}
    </div>
  );
}