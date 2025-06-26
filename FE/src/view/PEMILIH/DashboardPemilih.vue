<template>
    <div class="min-h-screen bg-gray-50 flex flex-col">
        <NavbarPemilih />

        <main class="flex-1 container mx-auto px-4 py-10">
            <h1 class="text-3xl font-bold text-center text-[#800000] mb-10">Profil Mahasiswa</h1>

            <div class="bg-white shadow rounded-xl p-6 md:p-8 max-w-4xl mx-auto w-full space-y-6">
                <!-- Info Akun -->
                <section>
                    <h2 class="text-xl font-semibold text-[#800000] mb-4 border-b pb-2">Akun Mahasiswa</h2>
                    <div class="grid grid-cols-1 md:grid-cols-2 gap-4 text-sm md:text-base">
                        <div><strong>NIM : </strong> {{ user?.nim || '-' }}</div>
                        <div><strong>Role : </strong> {{ user?.role || '-' }}</div>
                    </div>
                </section>

                <!-- Info Mahasiswa -->
                <section v-if="data_mahasiswa">
                    <h2 class="text-xl font-semibold text-[#800000] mb-4 border-b pb-2">Data Mahasiswa</h2>
                    <div class="grid grid-cols-1 md:grid-cols-2 gap-4 text-sm md:text-base">
                        <div><strong>Nama : </strong> {{ data_mahasiswa.nama_depan }} {{ data_mahasiswa.nama_belakang }}
                        </div>
                        <div><strong>Nomor HP : </strong> {{ data_mahasiswa.no_hp }}</div>
                        <div><strong>Alamat : </strong> {{ data_mahasiswa.alamat }}</div>
                        <div><strong>Tempat, Tanggal Lahir: </strong> {{ data_mahasiswa.tempatLahir }},
                            {{ formatTanggal(data_mahasiswa.tanggalLahir) }}</div>
                        <div><strong>NIK : </strong> {{ data_mahasiswa.nik }}</div>
                        <div><strong>Pekerjaan : </strong> {{ data_mahasiswa.pekerjaan }}</div>
                        <div><strong>Status Perkawinan : </strong> {{ data_mahasiswa.statusPerkawinan }}</div>
                        <div><strong>Agama : </strong> {{ data_mahasiswa.agama }}</div>
                        <div>
                            <strong>Status Memilih : </strong>
                            <span :class="data_mahasiswa.sudah_memilih ? 'text-green-600' : 'text-red-600'">
                                {{ data_mahasiswa.sudah_memilih ? 'Sudah Memilih' : 'Belum Memilih' }}
                            </span>
                        </div>
                        <div><strong>Waktu Dibuat : </strong> {{ formatTanggal(data_mahasiswa.createdAt) }}</div>
                        <div><strong>Terakhir Diperbarui : </strong> {{ formatTanggal(data_mahasiswa.updatedAt) }}</div>
                    </div>
                </section>

                <!-- Info TPS -->
                <section v-if="data_mahasiswa?.tps">
                    <h2 class="text-xl font-semibold text-[#800000] mb-4 border-b pb-2">Data TPS</h2>
                    <div class="grid grid-cols-1 md:grid-cols-2 gap-4 text-sm md:text-base">
                        <div><strong>Nama TPS:</strong> {{ data_mahasiswa.tps.nama_tps }}</div>
                        <div><strong>No TPS:</strong> {{ data_mahasiswa.tps.no_tps }}</div>
                        <div><strong>Daerah:</strong> {{ data_mahasiswa.tps.daerah?.nama_daerah_tps || '-' }}</div>
                        <div><strong>Wilayah:</strong> {{ data_mahasiswa.tps.wilayah?.nama_wilayah_tps || '-' }}</div>
                        <div><strong>Status TPS:</strong>
                            <span :class="data_mahasiswa.tps.is_open ? 'text-green-600' : 'text-gray-600'">
                                {{ data_mahasiswa.tps.is_open ? 'Dibuka' : 'Ditutup' }}
                            </span>
                        </div>
                    </div>
                </section>
            </div>
        </main>

        <FooterPemilih />
    </div>
</template>

<script>
import NavbarPemilih from '@/components/NavbarPemilih.vue'
import FooterPemilih from '@/components/FooterPemilih.vue'
import Swal from 'sweetalert2'

export default {
    name: 'DashboardLengkap',
    components: {
        NavbarPemilih,
        FooterPemilih,
    },
    computed: {
        user() {
            return this.$store.state.storeMahasiswa.userMahasiswa?.user || {}
        },
        data_mahasiswa() {
            return this.$store.state.storeMahasiswa.userMahasiswa?.data_mahasiswa || {}
        }
    },
    mounted() {
        this.cekStatusPemilih()
    },
    methods: {
        formatTanggal(dateStr) {
            if (!dateStr) return '-'
            const d = new Date(dateStr)
            return d.toLocaleDateString('id-ID', {
                day: '2-digit',
                month: 'long',
                year: 'numeric',
            })
        },
        cekStatusPemilih() {
            const { sudah_memilih, tps } = this.data_mahasiswa || {}

            if (!tps) return

            if (!tps.is_open) {
                Swal.fire({
                    icon: 'warning',
                    title: 'TPS Ditutup',
                    text: 'TPS Anda saat ini masih ditutup. Silakan tunggu dibuka.',
                })
            } else if (sudah_memilih) {
                Swal.fire({
                    icon: 'success',
                    title: 'Terima Kasih!',
                    text: 'Anda sudah memberikan suara.',
                })
            } else {
                Swal.fire({
                    icon: 'info',
                    title: 'Belum Memilih',
                    text: 'Silakan pilih kandidat Anda saat TPS dibuka.',
                })
            }
        }
    }
}
</script>

<style scoped>
strong {
    font-weight: 600;
}
</style>