<template>
    <div class="min-h-screen bg-white flex flex-col">
        <!-- Navbar -->
        <Navbar />

        <!-- Hero -->
        <section class="flex-1 flex flex-col items-center justify-center text-center text-[#800000] px-6 py-12">
            <h2 class="text-4xl md:text-5xl font-extrabold mb-4">Selamat Datang di Sistem E-Voting UMM</h2>
            <p class="text-lg md:text-xl mb-8 max-w-xl">
                Wujudkan pemilihan ketua yang modern, aman, dan transparan secara digital.
            </p>
            <router-link to="/login"
                class="bg-[#800000] text-white px-6 py-3 rounded font-semibold hover:bg-red-900 transition">
                Mulai Memilih
            </router-link>
        </section>

        <!-- Fitur -->
        <section id="fitur" class="bg-gray-50 py-12 text-gray-800">
            <div class="max-w-4xl mx-auto px-6">
                <h3 class="text-2xl font-bold mb-6 text-center">Fitur Utama</h3>
                <div class="grid md:grid-cols-3 gap-6">
                    <div class="p-4 bg-white rounded shadow text-center">
                        <h4 class="font-semibold text-lg mb-2 text-[#800000]">Login Aman</h4>
                        <p>Mahasiswa hanya bisa memilih sekali menggunakan NIM dan kode PIC unik.</p>
                    </div>
                    <div class="p-4 bg-white rounded shadow text-center">
                        <h4 class="font-semibold text-lg mb-2 text-[#800000]">Hasil Real-Time</h4>
                        <p>Hasil perolehan suara ditampilkan langsung setelah voting selesai.</p>
                    </div>
                    <div class="p-4 bg-white rounded shadow text-center">
                        <h4 class="font-semibold text-lg mb-2 text-[#800000]">Data Aman</h4>
                        <p>Semua data disimpan terenkripsi dan aman dalam sistem.</p>
                    </div>
                </div>
            </div>
        </section>

        <!-- Footer -->
        <Footer />
    </div>
</template>

<script>
import Navbar from '../components/NavbarLanding.vue'
import Footer from '../components/FooterLanding.vue'
export default {
    components: {
        Navbar,
        Footer
    }
}
</script>