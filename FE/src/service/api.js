import axios from "axios";

const currentDomain = "http://192.168.1.53:3000";
// : "http://localhost:3000";

const baseURL = `${currentDomain}/api`;

const API = axios.create({
  baseURL: baseURL,
  withCredentials: true,
  headers: {
    "X-Requested-With": "XMLHttpRequest",
  },
});
export default {
  getFullpathImage(img) {
    return `${currentDomain}/uploads/${img}`;
  },
  /*Auth*/
  LoginPost(data) {
    return API.post("/auth/login", data);
  },
  LogoutPost() {
    return API.post("/auth/logout");
  },
  getUserData() {
    return API.get("/data-mahasiswa");
  },
  /*Kandidat*/
  CreateKandidat(data) {
    return API.post("/kandidat", data, {
      headers: { "Content-Type": "multipart/form-data" },
    });
  },
  GetAllKandidat() {
    return API.get("/kandidat");
  },
  GetKandidatById(id) {
    return API.get(`/kandidat/${id}`);
  },
  GetKandidatVoteCount() {
    return API.get(`/kandidat/vote-count`);
  },
  UpdateKandidat(id, data) {
    return API.put(`/kandidat/${id}`, data, {
      headers: { "Content-Type": "multipart/form-data" },
    });
  },
  PilihKandidat(data) {
    return API.put(`/kandidat/pilih`, data);
  },
  DeleteKandidat(id) {
    return API.delete(`/kandidat/${id}`);
  },
  /*Wilayah*/
  CreateWilayah(data) {
    return API.post("/wilayah", data);
  },
  GetAllWilayah() {
    return API.get("/wilayah");
  },
  GetWilayahById(id) {
    return API.get(`/wilayah/${id}`);
  },
  UpdateWilayah(id, data) {
    return API.put(`/wilayah/${id}`, data);
  },
  DeleteWilayah(id) {
    return API.delete(`/wilayah/${id}`);
  },
  /*Daerah*/
  CreateDaerah(data) {
    return API.post("/daerah", data);
  },
  GetAllDaerah() {
    return API.get("/daerah");
  },
  GetDaerahById(id) {
    return API.get(`/daerah/${id}`);
  },
  GetDaerahByWilayahId(id) {
    return API.get(`/daerah/wilayah/${id}`);
  },
  UpdateDaerah(id, data) {
    return API.put(`/daerah/${id}`, data);
  },
  DeleteDaerah(id) {
    return API.delete(`/daerah/${id}`);
  },
  /*Tps*/
  CreateTps(data) {
    return API.post("/tps", data);
  },
  generateTPSMahasiswa() {
    return API.post("/tps/generate");
  },
  GetAllTps() {
    return API.get("/tps");
  },
  GetRekapTps() {
    return API.get("/tps/rekap");
  },
  GetRekapPetugas() {
    return API.get("/tps/petugas");
  },
  GetTpsById(id) {
    return API.get(`/tps/${id}`);
  },
  UpdateTps(id, data) {
    return API.put(`/tps/${id}`, data);
  },
  UpdateToPetugasTps(id) {
    return API.put(`/tps/petugas/${id}`);
  },
  UpdateStatusTps(id, data) {
    return API.put(`/tps/status/${id}`, data);
  },
  DeleteTps(id) {
    return API.delete(`/tps/${id}`);
  },
  /* Mahasiswa */
  Createmahasiswa(data) {
    return API.post("/data-mahasiswa", data);
  },
  Importmahasiswa(data) {
    return API.post("/data-mahasiswa/import", data);
  },
  Exportmahasiswa() {
    return API.get("/data-mahasiswa/export");
  },
  GetAllmahasiswa() {
    return API.get("/data-mahasiswa/all");
  },
  GetmahasiswaById(id) {
    return API.get(`/data-mahasiswa/${id}`);
  },
  Updatemahasiswa(id, data) {
    return API.put(`/data-mahasiswa/${id}`, data);
  },
  Deletemahasiswa(id) {
    return API.delete(`/data-mahasiswa/${id}`);
  },
};
