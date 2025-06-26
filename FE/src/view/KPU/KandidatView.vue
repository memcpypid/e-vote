<template>
    <NavbarKpu />
    <div class="min-h-screen bg-gray-50 py-10">
        <div class="container mx-auto px-4">
            <h1 class="text-3xl font-bold text-center text-red-700 mb-8">Manajemen Kandidat</h1>

            <!-- Form Tambah/Edit Kandidat -->
            <form @submit.prevent="submitForm" class="bg-white p-6 rounded-xl shadow-md mb-8">
                <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
                    <div>
                        <label class="block text-sm font-medium text-gray-700 mb-1">Nama Kandidat</label>
                        <input v-model="form.nama_kandidat" type="text" class="w-full border rounded-lg p-2" required />
                    </div>
                    <div>
                        <label class="block text-sm font-medium text-gray-700 mb-1">NIM Kandidat</label>
                        <input v-model="form.nim_kandidat" type="text" class="w-full border rounded-lg p-2" required />
                    </div>
                    <div>
                        <label class="block text-sm font-medium text-gray-700 mb-1">Nama Pasangan (Opsional)</label>
                        <input v-model="form.nama_pasangan" type="text" class="w-full border rounded-lg p-2" />
                    </div>
                    <div>
                        <label class="block text-sm font-medium text-gray-700 mb-1">NIM Pasangan (Opsional)</label>
                        <input v-model="form.nim_pasangan" type="text" class="w-full border rounded-lg p-2" />
                    </div>
                </div>

                <div class="mt-4">
                    <label class="block text-sm font-medium text-gray-700 mb-1">Visi</label>
                    <textarea v-model="form.visi" class="w-full border rounded-lg p-2" rows="3" required></textarea>
                </div>
                <div class="mt-4">
                    <label class="block text-sm font-medium text-gray-700 mb-1">Misi</label>
                    <textarea v-model="form.misi" class="w-full border rounded-lg p-2" rows="4" required></textarea>
                </div>
                <div class="mt-4">
                    <label class="block text-sm font-medium text-gray-700 mb-1">Foto Kandidat</label>
                    <input type="file" @change="handleFileChange" accept="image/*"
                        class="w-full border rounded-lg p-2" />
                    <div v-if="previewFoto" class="mt-2">
                        <img :src="previewFoto" alt="Preview" class="w-24 h-24 object-cover rounded-full" />
                    </div>
                </div>
                <div class="mt-6 flex justify-end space-x-2">
                    <button type="submit" class="bg-red-700 hover:bg-red-800 text-white px-4 py-2 rounded shadow">
                        {{ editMode ? 'Update' : 'Tambah' }}
                    </button>
                    <button v-if="editMode" type="button" @click="cancelEdit"
                        class="px-4 py-2 border border-gray-400 text-gray-600 rounded hover:bg-gray-100">
                        Batal
                    </button>
                </div>
            </form>

            <!-- Tabel Kandidat -->
            <div class="bg-white p-6 rounded-xl shadow-md overflow-x-auto">
                <h2 class="text-xl font-semibold text-gray-800 mb-4">Daftar Kandidat</h2>
                <table class="min-w-full table-auto text-sm border">
                    <thead class="bg-gray-100 text-gray-700">
                        <tr>
                            <th class="p-3 border">#</th>
                            <th class="p-3 border">Nama</th>
                            <th class="p-3 border">NIM</th>
                            <th class="p-3 border">Pasangan</th>
                            <th class="p-3 border">Visi</th>
                            <th class="p-3 border">Aksi</th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr v-for="(kandidat, index) in daftarKandidat" :key="kandidat.id_kandidat"
                            class="hover:bg-gray-50">
                            <td class="p-3 border">{{ index + 1 }}</td>
                            <td class="p-3 border">{{ kandidat.nama_kandidat }}</td>
                            <td class="p-3 border">{{ kandidat.nim_kandidat }}</td>
                            <td class="p-3 border whitespace-nowrap">
                                <div>{{ kandidat.nama_pasangan || '-' }}</div>
                                <div class="text-xs text-gray-500">{{ kandidat.nim_pasangan || '' }}</div>
                            </td>
                            <td class="p-3 border">
                                <span class="block text-ellipsis overflow-hidden whitespace-nowrap max-w-[200px]">
                                    {{ kandidat.visi }}
                                </span>
                            </td>
                            <td class="p-3 border space-x-2 whitespace-nowrap">
                                <button @click="editKandidat(kandidat)" class="text-blue-600 hover:underline">
                                    Edit
                                </button>
                                <button @click="deleteKandidat(kandidat.id_kandidat)"
                                    class="text-red-600 hover:underline">
                                    Hapus
                                </button>
                            </td>
                        </tr>
                        <tr v-if="daftarKandidat.length === 0">
                            <td colspan="6" class="text-center py-4 text-gray-500">Belum ada data kandidat.</td>
                        </tr>
                    </tbody>
                </table>
            </div>
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
    name: 'KandidatPage',
    components: {
        NavbarKpu,
        FooterKpu
    },
    data() {
        return {
            daftarKandidat: [],
            form: {
                id_kandidat: null,
                nama_kandidat: '',
                nim_kandidat: '',
                nama_pasangan: '',
                nim_pasangan: '',
                visi: '',
                misi: ''
            },
            fotoFile: null,
            previewFoto: null,
            editMode: false
        };
    },
    mounted() {
        this.fetchKandidat();
    },
    methods: {
        async fetchKandidat() {
            Swal.fire({
                title: 'Memuat data...',
                allowOutsideClick: false,
                didOpen: () => Swal.showLoading()
            });
            try {
                const res = await api.GetAllKandidat();
                this.daftarKandidat = res.data;
                Swal.close();
            } catch (err) {
                console.error('Gagal mengambil data kandidat:', err);
                Swal.fire('Gagal', 'Gagal mengambil data kandidat', 'error');
            }
        },
        handleFileChange(event) {
            const file = event.target.files[0];
            if (file) {
                this.fotoFile = file;
                this.previewFoto = URL.createObjectURL(file);
            }
        }
        ,
        async submitForm() {
            Swal.fire({
                title: this.editMode ? 'Memperbarui kandidat...' : 'Menambahkan kandidat...',
                allowOutsideClick: false,
                didOpen: () => Swal.showLoading()
            });

            try {
                const formData = new FormData();
                formData.append('nama_kandidat', this.form.nama_kandidat);
                formData.append('nim_kandidat', this.form.nim_kandidat);
                formData.append('nama_pasangan', this.form.nama_pasangan);
                formData.append('nim_pasangan', this.form.nim_pasangan);
                formData.append('visi', this.form.visi);
                formData.append('misi', this.form.misi);
                if (this.fotoFile) {
                    formData.append('foto', this.fotoFile);
                }

                if (this.editMode) {
                    await api.UpdateKandidat(this.form.id_kandidat, formData);
                    await Swal.fire('Berhasil', 'Data kandidat berhasil diperbarui', 'success');
                } else {
                    await api.CreateKandidat(formData);
                    await Swal.fire('Berhasil', 'Data kandidat berhasil ditambahkan', 'success');
                }

                this.resetForm();
                this.fetchKandidat();
                window.location.reload()
            } catch (err) {
                console.error('Gagal menyimpan kandidat:', err);
                Swal.fire('Gagal', 'Terjadi kesalahan saat menyimpan data kandidat', 'error');
            }
        },
        editKandidat(kandidat) {
            this.form = { ...kandidat };
            this.editMode = true;
        },

        cancelEdit() {
            this.resetForm();
        },

        resetForm() {
            this.form = {
                id_kandidat: null,
                nama_kandidat: '',
                nim_kandidat: '',
                nama_pasangan: '',
                nim_pasangan: '',
                visi: '',
                misi: ''
            };
            this.editMode = false;
            this.fotoFile = null;
            this.previewFoto = null;
        }
        ,

        async deleteKandidat(id) {
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
                    title: 'Menghapus data...',
                    allowOutsideClick: false,
                    didOpen: () => Swal.showLoading()
                });
                try {
                    await api.DeleteKandidat(id);
                    Swal.fire('Berhasil', 'Data kandidat berhasil dihapus', 'success');
                    this.fetchKandidat();
                } catch (err) {
                    console.error('Gagal menghapus kandidat:', err);
                    Swal.fire('Gagal', 'Tidak dapat menghapus kandidat', 'error');
                }
            }
        }
    }
};
</script>



<style scoped>
input,
textarea {
    @apply border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-red-500;
}
</style>