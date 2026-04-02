export default function ParkingPage() {
    return (
      <div>
  
        <h1 className="text-2xl font-bold mb-6">
          Parkir Kendaraan
        </h1>
  
        <div className="grid grid-cols-3 gap-4 mb-6">
  
          <div className="bg-white p-4 rounded shadow">
            Total Masuk
          </div>
  
          <div className="bg-white p-4 rounded shadow">
            Total Keluar
          </div>
  
          <div className="bg-white p-4 rounded shadow">
            Slot Parkir
          </div>
  
        </div>
  
        <div className="bg-white p-4 rounded shadow">
  
          <h2 className="font-semibold mb-4">
            Tambah Kendaraan
          </h2>
  
          <form className="flex gap-3">
  
            <input
              placeholder="Plat Nomor"
              className="border p-2 rounded w-48"
            />
  
            <select className="border p-2 rounded">
              <option>Motor</option>
              <option>Mobil</option>
            </select>
  
            <button className="bg-blue-500 text-white px-4 rounded">
              Masuk
            </button>
  
          </form>
  
        </div>
  
      </div>
    );
  }