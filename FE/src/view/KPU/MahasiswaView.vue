<template>
    <NavbarKpu />
    <div class="min-h-screen bg-gray-50 py-10">
        <div class="container mx-auto px-4">
            <h1 class="text-3xl font-bold text-center text-red-700 mb-8">Data Mahasiswa</h1>

            <div class="flex flex-col sm:flex-row sm:justify-between sm:items-center gap-4 mb-6">
                <div class="flex items-center gap-2">
                    <input type="file" @change="handleFile" accept=".xlsx"
                        class="file:px-3 file:py-1 file:rounded file:border-none file:bg-blue-600 file:text-white file:cursor-pointer text-sm" />
                    <button @click="uploadExcel"
                        class="bg-green-600 hover:bg-green-700 text-white px-4 py-2 rounded shadow text-sm">
                        Upload Excel
                    </button>
                </div>
                <div class="flex gap-2 items-center">
                    <button @click="GenerateTPSMahasiswa"
                        class="bg-green-600 hover:bg-green-700 text-white px-4 py-2 rounded shadow text-sm">
                        Bagi TPS
                    </button>
                </div>
                <div class="flex gap-2 items-center">
                    <input v-model="searchQuery" @input="filterMahasiswa" type="text" placeholder="Cari NIM/Nama..."
                        class="border px-3 py-2 rounded text-sm w-64" />
                    <button @click="exportToExcel"
                        class="bg-green-600 hover:bg-green-700 text-white px-4 py-2 rounded shadow text-sm">
                        Export ke Excel
                    </button>
                </div>

            </div>

            <div class="overflow-x-auto bg-white shadow-md rounded-xl">
                <table class="min-w-full table-auto">
                    <thead class="bg-gray-100 text-gray-700 text-left text-sm">
                        <tr>
                            <th class="p-3 border">#</th>
                            <th class="p-3 border">NIM</th>
                            <th class="p-3 border">Nama</th>
                            <th class="p-3 border">No TPS</th>
                            <th class="p-3 border">Sudah Memilih</th>
                            <th class="p-3 border">Role</th>
                            <th class="p-3 border">Aksi</th>
                        </tr>
                    </thead>
                    <tbody class="text-sm">
                        <tr v-for="(mhs, index) in paginatedMahasiswa" :key="mhs.id_data_mahasiswa"
                            class="hover:bg-gray-50">
                            <td class="p-3 border">{{ index + 1 + (currentPage - 1) * itemsPerPage }}</td>
                            <td class="p-3 border">{{ mhs.user?.nim || '-' }}</td>
                            <td class="p-3 border">{{ mhs.nama_depan }} {{ mhs.nama_belakang }}</td>
                            <td class="p-3 border">{{ mhs.tps?.no_tps || '-' }}</td>

                            <td class="p-3 border">
                                <span :class="mhs.sudah_memilih ? 'text-green-600' : 'text-red-600'">
                                    {{ mhs.sudah_memilih ? 'Ya' : 'Belum' }}
                                </span>
                            </td>
                            <td class="p-3 border">{{ mhs.user?.role || '-' }}</td>
                            <td class="p-3 border">
                                <button @click="ubahRole(mhs.user?.id_user)"
                                    class="text-blue-600 hover:underline text-sm">
                                    {{ mhs.user?.role === 'mahasiswa' ? "Jadikan Petugas" : "Jadikan Mahasiswa" }}
                                </button>
                            </td>
                        </tr>
                        <tr v-if="filteredMahasiswa.length === 0">
                            <td colspan="6" class="text-center text-gray-500 py-6">
                                Belum ada data mahasiswa.
                            </td>
                        </tr>
                    </tbody>
                </table>
                <div class="flex justify-between items-center mt-4 px-4 py-2">
                    <button @click="prevPage" :disabled="currentPage === 1"
                        class="px-3 py-1 bg-gray-200 rounded">Sebelumnya</button>
                    <span>Halaman {{ currentPage }}</span>
                    <button @click="nextPage" :disabled="currentPage * itemsPerPage >= filteredMahasiswa.length"
                        class="px-3 py-1 bg-gray-200 rounded">Berikutnya</button>
                </div>
            </div>
        </div>
    </div>
    <FooterKpu />
</template>

<script>
import api from '@/service/api';
import * as XLSX from "xlsx";
import Swal from 'sweetalert2';
import NavbarKpu from '@/components/NavbarKpu.vue';
import FooterKpu from '@/components/FooterKpu.vue';

export default {
    name: 'MahasiswaPage',
    components: { NavbarKpu, FooterKpu },
    data() {
        return {
            mahasiswaList: [],
            filteredMahasiswa: [],
            searchQuery: '',
            file: null,
            currentPage: 1,
            itemsPerPage: 50
        };
    },
    computed: {
        paginatedMahasiswa() {
            const start = (this.currentPage - 1) * this.itemsPerPage;
            return this.filteredMahasiswa.slice(start, start + this.itemsPerPage);
        }
    },
    mounted() {
        this.fetchMahasiswa();
    },
    methods: {
        handleFile(e) {
            this.file = e.target.files[0];
        },
        async uploadExcel() {
            if (!this.file) return Swal.fire("Gagal", "Silakan pilih file terlebih dahulu!", "warning");
            const formData = new FormData();
            formData.append("file", this.file);
            try {
                Swal.fire({ title: "Mengunggah file...", allowOutsideClick: false, didOpen: () => Swal.showLoading() });
                await api.Importmahasiswa(formData);
                await Swal.fire("Berhasil", "Data mahasiswa berhasil diimpor", "success");
                this.fetchMahasiswa();
                this.file = null;
                window.location.reload()
            } catch (err) {
                console.error("Import error:", err);
                Swal.fire("Gagal", "Gagal mengimpor data", "error");
            }
        },
        async fetchMahasiswa() {
            try {
                Swal.fire({ title: "Memuat data mahasiswa...", allowOutsideClick: false, didOpen: () => Swal.showLoading() });
                const res = await api.GetAllmahasiswa();
                this.mahasiswaList = res.data;
                this.filteredMahasiswa = res.data;
                Swal.close();
            } catch (err) {
                console.error("Fetch mahasiswa error:", err);
                Swal.fire("Gagal", "Gagal mengambil data mahasiswa", "error");
            }
        },
        filterMahasiswa() {
            const q = this.searchQuery.toLowerCase();
            this.filteredMahasiswa = this.mahasiswaList.filter(m => {
                const nama = `${m.nama_depan} ${m.nama_belakang}`.toLowerCase();
                const nim = (m.user?.nim || '').toString().toLowerCase();  // ✅ cast ke string
                return nama.includes(q) || nim.includes(q);
            });
            this.currentPage = 1;
        },
        prevPage() {
            if (this.currentPage > 1) this.currentPage--;
        },
        nextPage() {
            if (this.currentPage * this.itemsPerPage < this.filteredMahasiswa.length) this.currentPage++;
        },
        async ubahRole(id_user) {
            if (!id_user) return Swal.fire("Gagal", "ID user tidak ditemukan", "warning");
            const konfirmasi = await Swal.fire({ title: 'Yakin ingin mengubah Role?', icon: 'question', showCancelButton: true, confirmButtonText: 'Ya', cancelButtonText: 'Batal' });
            if (!konfirmasi.isConfirmed) return;
            try {
                Swal.fire({ title: "Mengubah role...", allowOutsideClick: false, didOpen: () => Swal.showLoading() });
                await api.UpdateToPetugasTps(id_user)
                Swal.fire("Berhasil", "Role berhasil diubah!", "success");
                this.fetchMahasiswa();
            } catch (err) {
                console.error("Ubah role error:", err);
                Swal.fire("Gagal", "Gagal mengubah role", "error");
            }
        },
        exportToExcel() {
            try {
                Swal.fire({ title: "Menyiapkan file Excel...", allowOutsideClick: false, didOpen: () => Swal.showLoading() });
                const data = this.filteredMahasiswa.map((mhs, index) => ({
                    No: index + 1,
                    NIM: mhs.user?.nim || '-',
                    Nama: `${mhs.nama_depan} ${mhs.nama_belakang}`,
                    "No TPS": mhs.tps?.no_tps || '-',
                    "Sudah Memilih": mhs.sudah_memilih ? "Ya" : "Belum",
                    "Nomor HP": mhs.no_hp,
                    Alamat: mhs.alamat,
                    "Tempat Lahir": mhs.tempatLahir,
                    "Tanggal Lahir": mhs.tanggalLahir,
                    Pekerjaan: mhs.pekerjaan,
                    "Status Perkawinan": mhs.statusPerkawinan,
                    Agama: mhs.agama,
                }));
                const worksheet = XLSX.utils.json_to_sheet(data);
                const workbook = XLSX.utils.book_new();
                XLSX.utils.book_append_sheet(workbook, worksheet, "Data Mahasiswa");
                XLSX.writeFile(workbook, "data_mahasiswa.xlsx");
                Swal.fire("Berhasil", "Data berhasil diexport ke Excel", "success");
            } catch (err) {
                console.error("Export error:", err);
                Swal.fire("Gagal", "Terjadi kesalahan saat export Excel", "error");
            }
        },
        async GenerateTPSMahasiswa() {
            try {
                await api.generateTPSMahasiswa()
                await Swal.fire("Berhasil", "Berhasil Membagi TPS Ke Mahasiswa", "success");
                this.fetchMahasiswa();
            } catch (error) {
                Swal.fire("Gagal", "Terjadi kesalahan saat Pembagian TPS", "error");
            }
        }
    }
};
</script>

<style scoped>
.table {
    @apply text-sm;
}
</style>
