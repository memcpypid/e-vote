<template>
  <div class="min-h-screen bg-gray-100 flex flex-col justify-center items-center px-4">
    <!-- Logo & Heading -->
    <div class="flex items-center space-x-3 mb-6">
      <img src="/logo-umm.png" alt="UMM Logo" class="w-12 h-12" />
      <h1 class="text-2xl md:text-3xl font-bold text-[#800000]">Login</h1>
    </div>

    <!-- Card -->
    <div class="bg-white shadow-lg rounded-lg p-8 w-full max-w-md">
      <form @submit.prevent="login">
        <div class="mb-4">
          <label for="nim" class="block text-sm font-medium text-gray-700 mb-1">NIM</label>
          <input v-model="nim" type="text" required
            class="w-full border border-gray-300 rounded px-3 py-2 focus:outline-none focus:ring-2 focus:ring-[#800000]" />
        </div>

        <div class="mb-6">
          <label for="pic" class="block text-sm font-medium text-gray-700 mb-1">PIC (Kode Rahasia)</label>
          <input v-model="pic" type="password" required
            class="w-full border border-gray-300 rounded px-3 py-2 focus:outline-none focus:ring-2 focus:ring-[#800000]" />
        </div>

        <button type="submit"
          class="w-full bg-[#800000] text-white py-2 rounded hover:bg-red-900 transition font-semibold">
          Login
        </button>
      </form>
    </div>

    <p class="text-sm text-gray-600 mt-6">
      Belum memiliki kode PIC? Silakan hubungi panitia pemilihan.
    </p>
  </div>
</template>

<script>
import Swal from 'sweetalert2';
export default {
  data() {
    return {
      nim: null,
      pic: null,
      isLoading: false
    }
  },
  computed: {
    isUserLoggedIn() {
      return this.$store.state.storeMahasiswa.UserMahasiswaIsLoggedIn;
    },
    userRole() {
      return this.$store.state.storeMahasiswa.userMahasiswaRole;
    }
  },
  mounted() {
    this.checkSession()
  },
  methods: {
    async checkSession() {
      try {
        this.isLoading = true
        await this.$store.dispatch("updateStoreMahasiswa");
        console.log(this.isUserLoggedIn)
        if (this.isUserLoggedIn) {
          Swal.fire({
            icon: 'success',
            title: 'Berhasil Login',
            text: 'Selamat datang kembali!',
            timer: 1500,
            showConfirmButton: false,
          });
          if (this.userRole === "mahasiswa") {
            await this.$router.push("/dashboard");
          } else if (this.userRole === "tps") {
            await this.$router.push("/petugas-tps");
          }
          else if (this.userRole === "kpu") {
            await this.$router.push("/petugas-kpu/rekapitulasi");
          } else {
            await this.$router.push("/auth/login");
          }
        }
      } catch (error) {
        Swal.fire({
          icon: 'error',
          title: 'Gagal Login',
          text: `${error}`,
          timer: 1500,
          showConfirmButton: false,
        }).finally(() => {
          this.isLoading = false;
        });
      }
    },
    async login() {
      try {
        this.isLoading = true
        await this.$store.dispatch("loginMahasiswa", { nim: this.nim, pic: this.pic });
        console.log(this.isUserLoggedIn)
        if (this.isUserLoggedIn) {
          Swal.fire({
            icon: 'success',
            title: 'Berhasil Login',
            text: 'Selamat datang kembali!',
            timer: 1500,
            showConfirmButton: false,
          });
          if (this.userRole === "mahasiswa") {
            await this.$router.push("/dashboard");
          } else if (this.userRole === "tps") {
            await this.$router.push("/petugas-tps");
          }
          else if (this.userRole === "kpu") {
            await this.$router.push("/petugas-kpu/rekapitulasi");
          } else {
            await this.$router.push("/auth/login");
          }
        }
      } catch (error) {
        Swal.fire({
          icon: 'error',
          title: 'Gagal Login',
          text: `${error}`,
          timer: 1500,
          showConfirmButton: false,
        }).finally(() => {
          this.isLoading = false;
        });
      }
    }
  }
}
</script>
