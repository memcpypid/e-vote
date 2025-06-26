<template>
    <NavbarKpu />
    <div class="min-h-screen bg-gray-50 pb-16">
        <!-- Judul -->
        <h1 class="text-3xl font-bold pt-10 text-center text-red-700 mb-8">Hasil Rekapitulasi</h1>

        <!-- Grafik Persentase -->
        <div class="container mx-auto px-4 grid grid-cols-1 md:grid-cols-2 gap-6">
            <div class="bg-white p-6 rounded-xl shadow">
                <h2 class="text-lg font-semibold mb-4">Persentase Pemilih</h2>
                <PieChart :data="pemilihChart" />
            </div>

            <div class="bg-white p-6 rounded-xl shadow">
                <h2 class="text-lg font-semibold mb-4">Persentase Kandidat</h2>
                <PieChart :data="kandidatChart" />
            </div>

        </div>
        <div class="container mx-auto px-4 mt-12">
            <div class="bg-white p-6 rounded-xl shadow">
                <h2 class="text-lg font-semibold mb-4">Grafik Jumlah Suara per Kandidat</h2>
                <div class="h-96">
                    <BarChart v-if="barChartData" :chartData="barChartData" />
                    <p v-else class="text-gray-500 text-sm">Memuat grafik...</p>
                </div>
            </div>
        </div>

        <!-- Rekap Per TPS -->
        <div class="container mx-auto mt-12 px-4">
            <h2 class="text-2xl font-semibold text-center text-red-600 mb-6">Rekapitulasi Per TPS</h2>
            <div class="overflow-x-auto rounded-xl shadow bg-white">
                <table class="min-w-full text-sm text-gray-700">
                    <thead class="bg-gray-100 text-gray-800 font-semibold">
                        <tr>
                            <th class="px-4 py-3 border">No</th>
                            <th class="px-4 py-3 border">Nama TPS</th>
                            <th class="px-4 py-3 border">Wilayah</th>
                            <th class="px-4 py-3 border">Daerah</th>
                            <th class="px-4 py-3 border">Total Pemilih</th>
                            <th class="px-4 py-3 border">Sudah Memilih</th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr v-for="(rekap, index) in rekapPerTps" :key="rekap.id_tps" class="hover:bg-gray-50">
                            <td class="px-4 py-2 border text-center">{{ index + 1 }}</td>
                            <td class="px-4 py-2 border">{{ rekap.nama_tps }}</td>
                            <td class="px-4 py-2 border">{{ rekap.wilayah }}</td>
                            <td class="px-4 py-2 border">{{ rekap.daerah }}</td>
                            <td class="px-4 py-2 border text-center">{{ rekap.total }}</td>
                            <td class="px-4 py-2 border text-center">{{ rekap.sudah }}</td>
                        </tr>
                        <tr v-if="rekapPerTps.length === 0">
                            <td colspan="6" class="text-center py-4 text-gray-500">
                                Belum ada data TPS.
                            </td>
                        </tr>
                    </tbody>
                </table>
            </div>
        </div>
    </div>
    <FooterKpu />
</template>


<script>
import PieChart from '@/components/PieChart.vue';
import BarChart from '@/components/BarChart.vue';
import api from '@/service/api';
import NavbarKpu from '@/components/NavbarKpu.vue';
import FooterKpu from '@/components/FooterKpu.vue';
export default {
    name: 'HasilRekapitulasiPage',
    components: { PieChart, BarChart, NavbarKpu, FooterKpu },
    data() {
        return {
            pemilihChart: [],
            kandidatChart: [],
            rekapPerTps: [],
            barChartData: null,
        };
    },
    mounted() {
        this.fetchData();
    },
    methods: {
        async fetchData() {
            try {
                const pemilihRes = await api.GetAllmahasiswa();
                const kandidatRes = await api.GetKandidatVoteCount();
                const tpsRes = await api.GetRekapTps();

                const sudah = pemilihRes.data.filter(d => d.sudah_memilih).length;
                const belum = pemilihRes.data.filter(d => !d.sudah_memilih).length;
                this.pemilihChart = [
                    { name: 'Sudah Memilih', value: sudah },
                    { name: 'Belum Memilih', value: belum }
                ];

                this.kandidatChart = kandidatRes.data.map(k => ({
                    name: k.nama_kandidat,
                    value: k.jumlah_suara
                }));
                this.barChartData = {
                    labels: kandidatRes.data.map(k => k.nama_kandidat),
                    datasets: [{
                        label: 'Jumlah Suara',
                        backgroundColor: '#B91C1C',
                        data: kandidatRes.data.map(k => k.jumlah_suara)
                    }]
                }
                this.rekapPerTps = tpsRes.data.map(tps => ({
                    id_tps: tps.id_tps,
                    nama_tps: tps.nama_tps,
                    wilayah: tps.wilayah?.nama_wilayah_tps || '-',
                    daerah: tps.daerah?.nama_daerah_tps || '-',
                    total: tps.total_pemilih,
                    sudah: tps.sudah_memilih
                }));
            } catch (err) {
                console.error('Gagal memuat data rekapitulasi:', err);
            }
        }
    }
};
</script>

<style scoped></style>
