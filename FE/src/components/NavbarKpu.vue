<template>
    <header class="bg-[#800000] text-white shadow-md">
        <div class="max-w-7xl mx-auto px-4 py-4 flex justify-between items-center">
            <div class="flex items-center space-x-3">
                <img src="/logo-umm.png" alt="UMM Logo" class="w-10 h-10" />
                <h1 class="text-lg md:text-xl font-bold">Panel Petugas KPU</h1>
            </div>

            <!-- Mobile Toggle -->
            <button class="md:hidden" @click="mobileOpen = !mobileOpen">
                <svg class="w-6 h-6" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"
                    stroke-linecap="round" stroke-linejoin="round">
                    <path v-if="!mobileOpen" d="M4 6h16M4 12h16M4 18h16" />
                    <path v-else d="M6 18L18 6M6 6l12 12" />
                </svg>
            </button>

            <!-- Desktop Nav -->
            <nav class="hidden md:flex space-x-4 items-center">
                <router-link to="/petugas-kpu/kandidat" class="hover:underline">Kandidat</router-link>
                <router-link to="/petugas-kpu/wilayah" class="hover:underline">Wilayah</router-link>
                <router-link to="/petugas-kpu/daerah" class="hover:underline">Daerah</router-link>
                <router-link to="/petugas-kpu/tps" class="hover:underline">TPS</router-link>
                <router-link to="/petugas-kpu/mahasiswa" class="hover:underline">Mahasiswa</router-link>
                <router-link to="/petugas-kpu/rekapitulasi" class="hover:underline">Rekapitulasi</router-link>
                <button @click="logout"
                    class="ml-4 bg-white text-[#800000] px-4 py-2 rounded hover:bg-gray-100 transition">
                    Logout
                </button>
            </nav>
        </div>

        <!-- Mobile Nav -->
        <nav v-if="mobileOpen" class="md:hidden bg-[#800000] px-4 pb-4 space-y-2">
            <router-link @click="mobileOpen = false" to="/petugas-kpu/kandidat"
                class="block text-white">Kandidat</router-link>
            <router-link @click="mobileOpen = false" to="/petugas-kpu/wilayah"
                class="block text-white">Wilayah</router-link>
            <router-link @click="mobileOpen = false" to="/petugas-kpu/daerah"
                class="block text-white">Daerah</router-link>
            <router-link @click="mobileOpen = false" to="/petugas-kpu/tps" class="block text-white">TPS</router-link>
            <router-link @click="mobileOpen = false" to="/petugas-kpu/mahasiswa"
                class="block text-white">Mahasiswa</router-link>
            <router-link @click="mobileOpen = false" to="/petugas-kpu/rekapitulasi"
                class="block text-white">Rekapitulasi</router-link>
            <button @click="logout"
                class="w-full bg-white text-[#800000] mt-2 px-4 py-2 rounded hover:bg-gray-100 transition">
                Logout
            </button>
        </nav>
    </header>
</template>

<script>
import Swal from 'sweetalert2';
export default {
    name: 'NavbarPetugasKPU',
    data() {
        return {
            mobileOpen: false
        };
    },
    methods: {
        async logout() {
            this.mobileOpen = false;
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
        }
    }
};
</script>