<template>
    <header class="bg-[#800000] text-white px-6 py-4 shadow-md">
        <div class="flex justify-between items-center">
            <!-- Logo & Title -->
            <div class="flex items-center space-x-3">
                <img src="/logo-umm.png" alt="UMM Logo" class="w-10 h-10" />
                <h1 class="text-xl font-bold">Panel Petugas TPS</h1>
            </div>

            <!-- Hamburger (Mobile) -->
            <button @click="toggleMenu" class="md:hidden text-white focus:outline-none">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24"
                    stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16" />
                </svg>
            </button>

            <!-- Desktop Menu -->
            <nav class="hidden md:flex space-x-4 items-center">
                <router-link to="/petugas-tps" class="hover:underline">Dashboard</router-link>
                <router-link to="/petugas-tps/kontrol" class="hover:underline">Kontrol TPS</router-link>
                <router-link to="/petugas-tps/daftar" class="hover:underline">Daftar Pemilih</router-link>
                <router-link to="/petugas-tps/rekapitulasi" class="hover:underline">Rekapitulasi</router-link>
                <button @click="logout"
                    class="ml-4 bg-white text-[#800000] px-4 py-2 rounded hover:bg-gray-100 transition">
                    Logout
                </button>
            </nav>
        </div>

        <!-- Mobile Menu -->
        <div v-if="isMenuOpen" class="md:hidden mt-4 space-y-2">
            <router-link @click="toggleMenu" to="/petugas-tps" class="block hover:underline">Dashboard</router-link>
            <router-link @click="toggleMenu" to="/petugas-tps/kontrol" class="block hover:underline">Kontrol
                TPS</router-link>
            <router-link @click="toggleMenu" to="/petugas-tps/daftar" class="block hover:underline">Daftar
                Pemilih</router-link>
            <router-link @click="toggleMenu" to="/petugas-tps/rekapitulasi"
                class="block hover:underline">Rekapitulasi</router-link>
            <button @click="logout"
                class="block w-full text-left bg-white text-[#800000] px-4 py-2 rounded hover:bg-gray-100 transition">
                Logout
            </button>
        </div>
    </header>
</template>

<script>
import Swal from 'sweetalert2';
export default {
    name: 'NavbarPetugas',
    data() {
        return {
            isMenuOpen: false,
        }
    },
    methods: {
        async logout() {
            this.isMenuOpen = false;
            const result = await Swal.fire({
                title: 'Apakah Anda yakin?',
                text: 'Anda akan keluar dari akun!',
                icon: 'warning',
                showCancelButton: true,
                confirmButtonColor: '#d33',
                cancelButtonColor: '#3085d6',
                confirmButtonText: 'Ya, Keluar!',
                cancelButtonText: 'Batal'
            });
            if (result.isConfirmed) {
                await this.$store.dispatch("logoutMahasiswa");
                Swal.fire(
                    'Berhasil!',
                    'Anda telah keluar dari akun.',
                    'success'
                ).then(() => {
                    this.$router.push({ path: "/Login" });
                });
            } else {

                Swal.fire(
                    'Dibatalkan',
                    'Anda tetap berada di akun.',
                    'info'
                );
            }
        },
        toggleMenu() {
            this.isMenuOpen = !this.isMenuOpen
        }
    },
}
</script>