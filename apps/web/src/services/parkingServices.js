import { api } from "./api";

export const getActiveParking = () => {
  return api.get("/parking");
};

export const createParking = (data) => {
  return api.post("/parking", data);
};

export const exitParking = (id) => {
  return api.post(`/parking/${id}/exit`);
};