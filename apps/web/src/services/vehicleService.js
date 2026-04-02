export async function getAllVehicles() {
    const token = localStorage.getItem("token");
  
    const response = await fetch("http://localhost:8080/api/vehicles", {
      headers: {
        "Authorization": `Bearer ${token}`
      }
    });
  
    if (!response.ok) {
      throw new Error("Gagal fetch kendaraan");
    }
  
    return response.json();
  }

  export async function getVehicleByLicensePlate(licensePlate) {
    const token = localStorage.getItem("token");
  
    const response = await fetch(
      `http://localhost:8080/api/vehicles/user/${licensePlate}`,
      {
        headers: {
          "Authorization": `Bearer ${token}`
        }
      }
    );
  
    if (!response.ok) {
      throw new Error("Kendaraan tidak ditemukan");
    }
  
    return response.json();
  }

  export async  function deleteVehicle(id) {
    const token = localStorage.getItem("token");
  
    const response = await fetch(`http://localhost:8080/api/vehicle/${id}`, {
      method: "DELETE",
      headers: {
        "Authorization": `Bearer ${token}`
      }
    });
  
    if (!response.ok) {
      throw new Error("Gagal menghapus kendaraan");
    }
  }