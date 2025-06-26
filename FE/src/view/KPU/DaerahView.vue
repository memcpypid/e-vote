<template>
    <NavbarKpu />
    <div class="min-h-screen bg-gray-50 px-4 sm:px-6 lg:px-8">
        <h1 class="text-2xl font-bold justify-center text-center pt-10 text-red-700 mb-6">Manajemen Daerah</h1>

        <form @submit.prevent="submitForm" class="bg-white p-6 rounded-xl shadow-md mb-6">
            <div>
                <label class="block text-sm font-medium text-gray-700">Wilayah</label>
                <select v-model="form.id_wilayah" class="mt-1 block w-full border rounded p-2" required>
                    <option disabled value="">Pilih Wilayah</option>
                    <option v-for="wilayah in wilayahList" :key="wilayah.id_wilayah" :value="wilayah.id_wilayah">
                        {{ wilayah.nama_wilayah_tps }}
                    </option>
                </select>
            </div>
            <div class="mt-4">
                <label class="block text-sm font-medium text-gray-700">Nama Daerah</label>
                <input v-model="form.nama_daerah_tps" type="text" class="mt-1 block w-full border rounded p-2"
                    required />
            </div>
            <div class="mt-4 text-right">
                <button type="submit" class="bg-red-700 text-white px-4 py-2 rounded">
                    {{ editMode ? 'Update' : 'Tambah' }}
                </button>
                <button v-if="editMode" type="button" @click="cancelEdit"
                    class="ml-2 px-4 py-2 border rounded">Batal</button>
            </div>
        </form>

        <div class="bg-white p-6 rounded-xl shadow-md">
            <h2 class="text-lg font-semibold text-gray-800 mb-4">Daftar Daerah</h2>
            <table class="min-w-full border">
                <thead>
                    <tr class="bg-gray-100 text-gray-600 text-left">
                        <th class="p-2 border">#</th>
                        <th class="p-2 border">Nama Daerah</th>
                        <th class="p-2 border">Wilayah</th>
                        <th class="p-2 border">Aksi</th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="(daerah, index) in daerahList" :key="daerah.id_daerah">
                        <td class="p-2 border">{{ index + 1 }}</td>
                        <td class="p-2 border">{{ daerah.nama_daerah_tps }}</td>
                        <td class="p-2 border">{{ getWilayahName(daerah.id_wilayah) }}</td>
                        <td class="p-2 border space-x-2">
                            <button @click="editDaerah(daerah)" class="text-blue-600 hover:underline">Edit</button>
                            <button @click="deleteDaerah(daerah.id_daerah)"
                                class="text-red-600 hover:underline">Hapus</button>
                        </td>
                    </tr>
                    <tr v-if="daerahList.length === 0">
                        <td colspan="4" class="text-center text-gray-500 py-4">Belum ada data daerah.</td>
                    </tr>
                </tbody>
            </table>
        </div>
    </div>
    <FooterKpu />
</template>

<script>
import api from '@/service/api';
import NavbarKpu from '@/components/NavbarKpu.vue';
import FooterKpu from '@/components/FooterKpu.vue';
import Swal from 'sweetalert2';

export default {
    name: 'DaerahPage',
    components: {
        NavbarKpu,
        FooterKpu
    },
    data() {
        return {
            daerahList: [],
            wilayahList: [],
            form: {
                id_daerah: null,
                nama_daerah_tps: '',
                id_wilayah: ''
            },
            editMode: false
        };
    },
    mounted() {
        this.fetchDaerah();
        this.fetchWilayah();
    },
    methods: {
        async fetchDaerah() {
            try {
                Swal.fire({ title: 'Memuat data...', allowOutsideClick: false, didOpen: () => Swal.showLoading() });
                const res = await api.GetAllDaerah();
                this.daerahList = res.data;
                Swal.close();
            } catch (err) {
                Swal.fire('Gagal', 'Gagal mengambil data daerah', 'error');
            }
        },
        async fetchWilayah() {
            try {
                const res = await api.GetAllWilayah();
                this.wilayahList = res.data;
            } catch (err) {
                Swal.fire('Gagal', 'Gagal mengambil data wilayah', 'error');
            }
        },
        getWilayahName(id) {
            const wilayah = this.wilayahList.find(w => w.id_wilayah === id);
            return wilayah ? wilayah.nama_wilayah_tps : '-';
        },
        async submitForm() {
            try {
                Swal.fire({ title: 'Menyimpan data...', allowOutsideClick: false, didOpen: () => Swal.showLoading() });
                if (this.editMode) {
                    await api.UpdateDaerah(this.form.id_daerah, this.form);
                    Swal.fire('Sukses', 'Data berhasil diperbarui', 'success');
                } else {
                    await api.CreateDaerah(this.form);
                    Swal.fire('Sukses', 'Data berhasil ditambahkan', 'success');
                }
                this.resetForm();
                this.fetchDaerah();
            } catch (err) {
                Swal.fire('Gagal', 'Gagal menyimpan data', 'error');
            }
        },
        editDaerah(daerah) {
            this.form = { ...daerah };
            this.editMode = true;
        },
        cancelEdit() {
            this.resetForm();
        },
        resetForm() {
            this.form = {
                id_daerah: null,
                nama_daerah_tps: '',
                id_wilayah: ''
            };
            this.editMode = false;
        },
        async deleteDaerah(id) {
            const confirm = await Swal.fire({
                title: 'Yakin ingin menghapus?',
                icon: 'warning',
                showCancelButton: true,
                confirmButtonColor: '#d33',
                cancelButtonColor: '#3085d6',
                confirmButtonText: 'Ya, hapus!'
            });
            if (confirm.isConfirmed) {
                try {
                    await api.DeleteDaerah(id);
                    this.fetchDaerah();
                    Swal.fire('Terhapus', 'Data berhasil dihapus', 'success');
                } catch (err) {
                    Swal.fire('Gagal', 'Gagal menghapus data', 'error');
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
