import { createRouter, createWebHistory } from "vue-router";
import store from "@/store/store";
const HomeView = () => import("@/view/HomeView.vue");
const LoginView = () => import("@/view/LoginView.vue");
/*======================= Pemilih ====================== */
const PemilihView = () => import("@/view/PEMILIH/PemilihView.vue");
const PemilihDasboard = () => import("@/view/PEMILIH/DashboardPemilih.vue");
/*======================= TPS ====================== */
const PetugasView = () => import("@/view/TPS/PetugasView.vue");

/*======================= KPU ====================== */
const KandidatView = () => import("@/view/KPU/KandidatView.vue");
const DaerahView = () => import("@/view/KPU/DaerahView.vue");
const WilayahView = () => import("@/view/KPU/WilayahView.vue");
const MahasiswaView = () => import("@/view/KPU/MahasiswaView.vue");
const TpsView = () => import("@/view/KPU/TpsView.vue");
const RekapitulasiView = () => import("@/view/KPU/HasilRekapitulasi.vue");
const routes = [
  {
    path: "/",
    component: HomeView,
    meta: { title: "E-Vote" },
  },
  {
    path: "/login",
    component: LoginView,
    meta: { title: "Login" },
  },
  /*Pemilih Mahasiswa*/
  {
    path: "/pemilihan",
    component: PemilihView,
    meta: { title: "Pemilih", requiresAuth: true },
  },
  {
    path: "/dashboard",
    component: PemilihDasboard,
    meta: { title: "Pemilih", requiresAuth: true },
  },
  /*Petugas TPS*/
  {
    path: "/petugas-tps",
    component: PetugasView,
    meta: { title: "TPS", requiresAuth: true },
  },
  /*Petugas KPU*/
  {
    path: "/petugas-kpu/kandidat",
    component: KandidatView,
    meta: { title: "KPU-Kandidat", requiresAuth: true },
  },
  {
    path: "/petugas-kpu/wilayah",
    component: WilayahView,
    meta: { title: "KPU-Kandidat", requiresAuth: true },
  },
  {
    path: "/petugas-kpu/daerah",
    component: DaerahView,
    meta: { title: "KPU-Kandidat", requiresAuth: true },
  },
  {
    path: "/petugas-kpu/mahasiswa",
    component: MahasiswaView,
    meta: { title: "KPU-Kandidat", requiresAuth: true },
  },
  {
    path: "/petugas-kpu/tps",
    component: TpsView,
    meta: { title: "KPU-Kandidat", requiresAuth: true },
  },
  {
    path: "/petugas-kpu/rekapitulasi",
    component: RekapitulasiView,
    meta: { title: "KPU-Kandidat", requiresAuth: true },
  },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
  scrollBehavior() {
    window.scrollTo({
      top: 0,
      behavior: "smooth",
    });
    return { top: 0 };
  },
});
router.beforeEach(async (to, from, next) => {
  var isAuthenticated = null;
  var userRole = null;
  if (to.meta.requiresAuth) {
    await store.dispatch("updateStoreMahasiswa");
  }
  isAuthenticated = store.state.storeMahasiswa.UserMahasiswaIsLoggedIn;
  userRole = store.state.storeMahasiswa.userMahasiswaRole;
  console.log(isAuthenticated);
  console.log(userRole);
  if (to.meta.requiresAuth && !isAuthenticated) {
    alert("Sesi Anda Habis!");
    next({ path: "/Login" });
  } else if (
    to.meta.requiresAuth &&
    to.meta.requiredRole &&
    to.meta.requiredRole !== userRole
  ) {
    alert("Tidak Memiliki Hak Akses!");
    next({ path: "/Login" });
  } else {
    next();
  }
});
router.afterEach((to) => {
  document.title = to.meta.title || "E-Vote";
});

export default router;
