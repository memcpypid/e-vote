<template>
    <NavbarKpu />
    <div class="min-h-screen min-w-full bg-gray-50 px-4 py-10">
        <h1 class="text-2xl md:text-3xl font-bold text-center text-red-700 mb-8">Manajemen Wilayah</h1>

        <form @submit.prevent="submitForm" class="bg-white p-6 rounded-xl shadow-md max-w-xl mx-auto mb-8">
            <label class="block text-sm font-medium text-gray-700 mb-1">Nama Wilayah</label>
            <input v-model="form.nama_wilayah_tps" type="text"
                class="mt-1 block w-full border rounded p-2 focus:outline-none focus:ring focus:border-red-500"
                required />
            <div class="mt-4 text-right space-x-2">
                <button type="submit" class="bg-red-700 text-white px-4 py-2 rounded hover:bg-red-800 transition">
                    {{ editMode ? 'Update' : 'Tambah' }}
                </button>
                <button v-if="editMode" type="button" @click="cancelEdit"
                    class="px-4 py-2 border rounded hover:bg-gray-100">
                    Batal
                </button>
            </div>
        </form>

        <div class="bg-white p-6 rounded-xl shadow-md overflow-x-auto">
            <h2 class="text-lg font-semibold text-gray-800 mb-4">Daftar Wilayah</h2>
            <table class="min-w-full border text-sm">
                <thead>
                    <tr class="bg-gray-100 text-gray-700 text-left">
                        <th class="p-2 border">#</th>
                        <th class="p-2 border">Nama Wilayah</th>
                        <th class="p-2 border">Aksi</th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="(wilayah, index) in wilayahList" :key="wilayah.id_wilayah" class="hover:bg-gray-50">
                        <td class="p-2 border">{{ index + 1 }}</td>
                        <td class="p-2 border">{{ wilayah.nama_wilayah_tps }}</td>
                        <td class="p-2 border space-x-2">
                            <button @click="editWilayah(wilayah)" class="text-blue-600 hover:underline">Edit</button>
                            <button @click="deleteWilayah(wilayah.id_wilayah)"
                                class="text-red-600 hover:underline">Hapus</button>
                        </td>
                    </tr>
                    <tr v-if="wilayahList.length === 0">
                        <td colspan="3" class="text-center text-gray-500 py-4">Belum ada data wilayah.</td>
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
    name: 'WilayahPage',
    components: {
        NavbarKpu,
        FooterKpu
    },
    data() {
        return {
            wilayahList: [],
            form: {
                id_wilayah: null,
                nama_wilayah_tps: ''
            },
            editMode: false
        };
    },
    mounted() {
        this.fetchWilayah();
    },
    methods: {
        async fetchWilayah() {
            Swal.fire({
                title: 'Memuat data...',
                allowOutsideClick: false,
                didOpen: () => Swal.showLoading()
            });
            try {
                const res = await api.GetAllWilayah();
                this.wilayahList = res.data;
                Swal.close();
            } catch (err) {
                console.error('Gagal mengambil data wilayah:', err);
                Swal.fire('Gagal', 'Gagal mengambil data wilayah', 'error');
            }
        },

        async submitForm() {
            Swal.fire({
                title: this.editMode ? 'Memperbarui wilayah...' : 'Menambahkan wilayah...',
                allowOutsideClick: false,
                didOpen: () => Swal.showLoading()
            });
            try {
                if (this.editMode) {
                    await api.UpdateWilayah(this.form.id_wilayah, this.form);
                    Swal.fire('Berhasil', 'Wilayah berhasil diperbarui', 'success');
                } else {
                    await api.CreateWilayah(this.form);
                    Swal.fire('Berhasil', 'Wilayah berhasil ditambahkan', 'success');
                }
                this.resetForm();
                this.fetchWilayah();
            } catch (err) {
                console.error('Gagal menyimpan wilayah:', err);
                Swal.fire('Gagal', 'Gagal menyimpan wilayah', 'error');
            }
        },

        editWilayah(wilayah) {
            this.form = { ...wilayah };
            this.editMode = true;
        },

        cancelEdit() {
            this.resetForm();
        },

        resetForm() {
            this.form = {
                id_wilayah: null,
                nama_wilayah_tps: ''
            };
            this.editMode = false;
        },

        async deleteWilayah(id) {
            const confirm = await Swal.fire({
                title: 'Yakin ingin menghapus?',
                text: 'Data ini akan dihapus secara permanen.',
                icon: 'warning',
                showCancelButton: true,
                confirmButtonText: 'Ya, hapus',
                cancelButtonText: 'Batal'
            });

            if (confirm.isConfirmed) {
                Swal.fire({
                    title: 'Menghapus wilayah...',
                    allowOutsideClick: false,
                    didOpen: () => Swal.showLoading()
                });
                try {
                    await api.delete(`/wilayah/${id}`);
                    Swal.fire('Berhasil', 'Wilayah berhasil dihapus', 'success');
                    this.fetchWilayah();
                } catch (err) {
                    console.error('Gagal menghapus wilayah:', err);
                    Swal.fire('Gagal', 'Gagal menghapus wilayah', 'error');
                }
            }
        }
    }
};
</script>

<style scoped>
input {
    @apply border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-red-500;
}
</style>