<template>
    <header class="bg-[#800000] text-white py-4 px-6 shadow-md">
        <div class="flex justify-between items-center">
            <!-- Logo dan Judul -->
            <div class="flex items-center space-x-3">
                <img src="/logo-umm.png" alt="UMM Logo" class="w-10 h-10" />
                <h1 class="text-lg md:text-xl font-bold whitespace-nowrap">E-Vote - Pemilih</h1>
            </div>

            <!-- Hamburger Mobile -->
            <button class="md:hidden focus:outline-none" @click="isOpen = !isOpen">
                <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24"
                    xmlns="http://www.w3.org/2000/svg">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16" />
                </svg>
            </button>

            <!-- Menu Desktop -->
            <nav class="hidden md:flex space-x-4 items-center">
                <router-link to="/dashboard" class="hover:underline">Dashboard</router-link>
                <router-link to="/pemilihan" class="hover:underline">Pemilihan</router-link>
                <button @click="logout"
                    class="bg-white text-[#800000] px-4 py-2 rounded hover:bg-gray-100 transition font-semibold">
                    Keluar
                </button>
            </nav>
        </div>

        <!-- Dropdown Mobile -->
        <div v-if="isOpen" class="mt-3 md:hidden space-y-2">
            <router-link to="/dashboard"
                class="block bg-white text-[#800000] px-4 py-2 rounded hover:bg-gray-100 transition font-semibold">
                Dashboard
            </router-link>
            <router-link to="/pemilihan"
                class="block bg-white text-[#800000] px-4 py-2 rounded hover:bg-gray-100 transition font-semibold">
                Pemilihan
            </router-link>
            <button @click="logout"
                class="w-full text-left bg-white text-[#800000] px-4 py-2 rounded hover:bg-gray-100 transition font-semibold">
                Keluar
            </button>
        </div>
    </header>
</template>

<script>
import Swal from 'sweetalert2';
export default {
    name: 'NavbarPemilih',
    data() {
        return {
            isOpen: false,
        }
    },
    methods: {
        async logout() {
            this.isOpen = false;
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
    },
}
</script>