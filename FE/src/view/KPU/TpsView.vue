<template>
    <NavbarKpu />
    <div class="min-h-screen bg-gray-50 px-4 sm:px-6 lg:px-8">
        <h1 class="text-2xl font-bold text-center pt-10 text-red-700 mb-6">Manajemen TPS</h1>

        <form @submit.prevent="submitForm" class="bg-white p-6 rounded-xl shadow-md mb-6">
            <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
                <div>
                    <label class="block text-sm font-medium text-gray-700">Nama TPS</label>
                    <input v-model="form.nama_tps" type="text" class="mt-1 w-full border rounded p-2" required />
                </div>
                <div>
                    <label class="block text-sm font-medium text-gray-700">No TPS</label>
                    <input v-model.number="form.no_tps" type="number" class="mt-1 w-full border rounded p-2" disabled />
                </div>
                <div>
                    <label class="block text-sm font-medium text-gray-700">Wilayah</label>
                    <select v-model="form.wilayah_id" @change="fetchDaerahByWilayahId(form.wilayah_id)"
                        class="mt-1 w-full border rounded p-2">
                        <option disabled value="">Pilih Wilayah</option>
                        <option v-for="w in wilayahList" :key="w.id_wilayah" :value="w.id_wilayah">
                            {{ w.nama_wilayah_tps }}
                        </option>
                    </select>
                </div>
                <div>
                    <label class="block text-sm font-medium text-gray-700">Daerah</label>
                    <select v-model="form.daerah_id" class="mt-1 w-full border rounded p-2">
                        <option disabled value="">Pilih Daerah</option>
                        <option v-for="d in daerahList" :key="d.id_daerah" :value="d.id_daerah">
                            {{ d.nama_daerah_tps }}
                        </option>
                    </select>
                </div>
            </div>

            <div class="mt-4 text-right">
                <button type="submit" class="bg-red-700 text-white px-4 py-2 rounded">
                    {{ editMode ? 'Update' : 'Tambah' }}
                </button>
                <button v-if="editMode" @click="cancelEdit" class="ml-2 px-4 py-2 border rounded">Batal</button>
            </div>
        </form>

        <div class="bg-white p-6 rounded-xl shadow-md overflow-x-auto">
            <h2 class="text-lg font-semibold mb-4">Daftar TPS</h2>
            <table class="w-full border text-sm">
                <thead class="bg-gray-100 text-gray-700">
                    <tr>
                        <th class="p-2 border">#</th>
                        <th class="p-2 border">Nama TPS</th>
                        <th class="p-2 border">No TPS</th>
                        <th class="p-2 border">Wilayah</th>
                        <th class="p-2 border">Daerah</th>
                        <th class="p-2 border">Status</th>
                        <th class="p-2 border">Aksi</th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="(t, index) in tpsList" :key="t.id_tps">
                        <td class="p-2 border text-center">{{ index + 1 }}</td>
                        <td class="p-2 border">{{ t.nama_tps }}</td>
                        <td class="p-2 border">{{ t.no_tps }}</td>
                        <td class="p-2 border">{{ getWilayahName(t.wilayah_id) }}</td>
                        <td class="p-2 border">{{ getDaerahName(t.daerah_id) }}</td>
                        <td class="p-2 border text-center">{{ t.is_open ? 'Dibuka' : 'Ditutup' }}</td>
                        <td class="p-2 border space-x-2 text-center">
                            <button @click="editTPS(t)" class="text-blue-600 hover:underline">Edit</button>
                            <button @click="deleteTPS(t.id_tps)" class="text-red-600 hover:underline">Hapus</button>
                        </td>
                    </tr>
                    <tr v-if="tpsList.length === 0">
                        <td colspan="7" class="text-center text-gray-500 py-4">Belum ada data TPS.</td>
                    </tr>
                </tbody>
            </table>
        </div>
    </div>
    <FooterKpu />
</template>

<script>
import Swal from 'sweetalert2';
import api from '@/service/api';
import NavbarKpu from '@/components/NavbarKpu.vue';
import FooterKpu from '@/components/FooterKpu.vue';

export default {
    name: 'TPSPage',
    components: { NavbarKpu, FooterKpu },
    data() {
        return {
            form: {
                id_tps: null,
                nama_tps: '',
                no_tps: '',
                wilayah_id: '',
                daerah_id: ''
            },
            tpsList: [],
            wilayahList: [],
            daerahList: [],
            daerahListAll: [],
            editMode: false
        };
    },
    mounted() {
        this.fetchTPS();
        this.fetchWilayah();
        this.fetchDaerah();
    },
    methods: {
        async fetchTPS() {
            try {
                const res = await api.GetAllTps();
                this.tpsList = res.data;
            } catch (err) {
                Swal.fire('Error', 'Gagal mengambil data TPS', 'error');
            }
        },
        async fetchWilayah() {
            try {
                const res = await api.GetAllWilayah();
                this.wilayahList = res.data;
            } catch (err) {
                Swal.fire('Error', 'Gagal mengambil data wilayah', 'error');
            }
        },
        async fetchDaerah() {
            try {
                const res = await api.GetAllDaerah();
                this.daerahListAll = res.data;
            } catch (err) {
                Swal.fire('Error', 'Gagal mengambil data daerah', 'error');
            }
        },
        async fetchDaerahByWilayahId(id) {
            try {
                const res = await api.GetDaerahByWilayahId(id);
                this.daerahList = res.data;
            } catch (err) {
                Swal.fire('Error', 'Gagal memuat daerah berdasarkan wilayah', 'error');
            }
        },
        getWilayahName(id) {
            const w = this.wilayahList.find(w => w.id_wilayah === id);
            return w ? w.nama_wilayah_tps : '-';
        },
        getDaerahName(id) {
            const d = this.daerahListAll.find(d => d.id_daerah === id);
            return d ? d.nama_daerah_tps : '-';
        },
        async submitForm() {
            try {
                Swal.showLoading();
                if (this.editMode) {
                    await api.UpdateTps(this.form.id_tps, this.form);
                    Swal.fire('Berhasil', 'Data TPS berhasil diperbarui', 'success');
                } else {
                    await api.CreateTps(this.form);
                    Swal.fire('Berhasil', 'Data TPS berhasil ditambahkan', 'success');
                }
                this.resetForm();
                this.fetchTPS();
            } catch (err) {
                Swal.fire('Error', 'Gagal menyimpan TPS', 'error');
            }
        },
        async editTPS(data) {
            this.form = { ...data };
            await this.fetchDaerahByWilayahId(this.form.wilayah_id);
            this.editMode = true;
        },
        cancelEdit() {
            this.resetForm();
        },
        resetForm() {
            this.form = {
                id_tps: null,
                nama_tps: '',
                no_tps: '',
                wilayah_id: '',
                daerah_id: ''
            };
            this.editMode = false;
        },
        async deleteTPS(id) {
            const confirm = await Swal.fire({
                title: 'Yakin ingin menghapus TPS ini?',
                icon: 'warning',
                showCancelButton: true,
                confirmButtonText: 'Ya, Hapus!',
                cancelButtonText: 'Batal'
            });
            if (confirm.isConfirmed) {
                try {
                    await api.DeleteTps(id);
                    this.fetchTPS();
                    Swal.fire('Berhasil', 'TPS berhasil dihapus', 'success');
                } catch (err) {
                    Swal.fire('Error', 'Gagal menghapus TPS', 'error');
                }
            }
        }
    }
};
</script>

<style scoped>
input,
select {
    @apply border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-red-500;
}
</style>
