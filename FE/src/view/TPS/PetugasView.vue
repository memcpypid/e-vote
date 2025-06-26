<template>
    <div class="min-h-screen bg-gray-50 flex flex-col">
        <NavbarPetugas />

        <main class="flex-1 container mx-auto px-4 py-8">
            <h1 class="text-2xl font-bold text-[#800000] mb-4 text-center">Panel Petugas TPS</h1>

            <!-- Informasi TPS -->
            <div class="bg-white rounded shadow p-6 mb-6 grid grid-cols-1 md:grid-cols-2 gap-6">
                <div>
                    <h2 class="text-xl font-semibold text-[#800000] mb-2">Informasi TPS</h2>
                    <ul class="text-sm space-y-1">
                        <li><strong>Nomor TPS:</strong> TPS {{ id_tps }}</li>
                        <li><strong>Total Pemilih:</strong> {{ totalPemilih }}</li>
                        <li><strong>Sudah Memilih:</strong> {{ sudahMemilih }}</li>
                        <li><strong>Belum Memilih:</strong> {{ totalPemilih - sudahMemilih }}</li>
                        <li><strong>Status TPS:</strong>
                            <span :class="statusTPS ? 'text-green-600' : 'text-red-600'">
                                {{ statusTPS ? 'Dibuka' : 'Ditutup' }}
                            </span>
                        </li>
                    </ul>
                    <button @click="toggleTPS"
                        class="mt-4 bg-[#800000] text-white px-4 py-2 rounded hover:bg-red-900 transition">
                        {{ statusTPS ? 'Tutup' : 'Buka' }} Pemilihan
                    </button>
                </div>

                <!-- Pie Chart -->
                <div>
                    <h2 class="text-xl font-semibold text-[#800000] mb-2">Partisipasi Pemilih</h2>
                    <canvas id="pieChart" height="200"></canvas>
                </div>
            </div>

            <!-- Rekap Suara -->
            <div class="bg-white rounded shadow p-6">
                <h2 class="text-xl font-semibold text-[#800000] mb-4">Rekapitulasi Suara</h2>
                <canvas id="rekapChart" height="100"></canvas>

                <!-- Persentase per Kandidat -->
                <ul class="mt-6 space-y-2 text-sm">
                    <li v-for="kandidat in kandidatList" :key="kandidat.id" class="flex justify-between border-b pb-2">
                        <span>{{ kandidat.nama }}</span>
                        <span>{{ hitungPersentase(kandidat.suara) }}%</span>
                    </li>
                </ul>
            </div>
        </main>

        <FooterPetugas />
    </div>
</template>

<script>
import NavbarPetugas from '@/components/NavbarPetugas.vue';
import FooterPetugas from '@/components/FooterPetugas.vue';
import Chart from 'chart.js/auto';
import api from '@/service/api';
import Swal from 'sweetalert2';

export default {
    name: 'PetugasPage',
    components: { NavbarPetugas, FooterPetugas },
    data() {
        return {
            statusTPS: false,
            totalPemilih: 0,
            sudahMemilih: 0,
            kandidatList: [],
            id_tps: 0,
            barChart: null,
            pieChart: null,
        };
    },
    computed: {
        totalSuara() {
            return this.kandidatList.reduce((sum, k) => sum + k.suara, 0);
        },
    },
    methods: {
        async toggleTPS() {
            this.statusTPS = !this.statusTPS;
            const payload = { is_open: this.statusTPS };
            try {
                await api.UpdateStatusTps(this.id_tps, payload);
                Swal.fire({
                    icon: 'success',
                    title: 'Status TPS diperbarui',
                    text: this.statusTPS ? 'TPS dibuka' : 'TPS ditutup',
                });
            } catch (error) {
                console.error(error);
                Swal.fire({
                    icon: 'error',
                    title: 'Gagal memperbarui status TPS',
                    text: error.response?.data?.error || 'Terjadi kesalahan saat memperbarui status.',
                });
            }
        },
        hitungPersentase(suara) {
            if (this.totalSuara === 0) return 0;
            return ((suara / this.totalSuara) * 100).toFixed(1);
        },
        async fetchTPSData() {
            try {
                const res = await api.GetRekapPetugas();
                this.totalPemilih = res.data.total_pemilih;
                this.sudahMemilih = res.data.sudah_memilih;
                this.statusTPS = res.data.is_open;
                this.id_tps = res.data.id_tps;
                this.kandidatList = res.data.kandidat.map(k => ({
                    id: k.id,
                    nama: k.nama,
                    suara: k.suara,
                }));
                this.renderPieChart();
                this.renderBarChart();
            } catch (err) {
                console.error("Gagal memuat data TPS:", err);
                Swal.fire({
                    icon: 'error',
                    title: 'Gagal memuat data TPS',
                    text: err.response?.data?.error || 'Terjadi kesalahan saat memuat data.',
                });
            }
        },
        renderBarChart() {
            if (this.barChart) this.barChart.destroy();
            const ctx = document.getElementById('rekapChart').getContext('2d');
            this.barChart = new Chart(ctx, {
                type: 'bar',
                data: {
                    labels: this.kandidatList.map(k => k.nama),
                    datasets: [
                        {
                            label: 'Jumlah Suara',
                            data: this.kandidatList.map(k => k.suara),
                            backgroundColor: ['#800000', '#a94442', '#c9302c'],
                        },
                    ],
                },
                options: {
                    responsive: true,
                    scales: {
                        y: {
                            beginAtZero: true,
                            ticks: { precision: 0 },
                        },
                    },
                },
            });
        },
        renderPieChart() {
            if (this.pieChart) this.pieChart.destroy();
            const ctx = document.getElementById('pieChart').getContext('2d');
            this.pieChart = new Chart(ctx, {
                type: 'pie',
                data: {
                    labels: ['Sudah Memilih', 'Belum Memilih'],
                    datasets: [
                        {
                            data: [this.sudahMemilih, this.totalPemilih - this.sudahMemilih],
                            backgroundColor: ['#800000', '#ccc'],
                        },
                    ],
                },
                options: {
                    responsive: true,
                },
            });
        },
    },
    mounted() {
        this.fetchTPSData();
    },
};
</script>
