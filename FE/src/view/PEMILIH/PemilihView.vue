<template>
    <div class="min-h-screen flex flex-col bg-gray-50">
        <NavbarPemilih />

        <main class="flex-1 container mx-auto px-4 py-8">
            <h1 class="text-2xl font-bold text-[#800000] mb-6 text-center">Pilih Kandidat Ketua</h1>

            <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
                <div v-for="kandidat in kandidatList" :key="kandidat.id_kandidat" @click="pilihKandidat(kandidat)"
                    class="cursor-pointer bg-white border rounded shadow hover:shadow-lg transition p-4 text-center"
                    :class="{
                        'border-[#800000] ring-2 ring-[#800000]': kandidatTerpilih === kandidat.id_kandidat,
                    }">
                    <img :src="getfullpath(kandidat.foto)" alt="Foto Kandidat"
                        class="w-24 h-24 object-cover mx-auto rounded-full mb-4" />
                    <h2 class="text-xl font-semibold text-[#800000] mb-1">{{ kandidat.nama_kandidat }}</h2>
                    <p class="text-sm text-gray-600">Visi: {{ kandidat.visi }}</p>
                    <p class="text-sm text-gray-600">Misi: {{ kandidat.misi }}</p>
                </div>
            </div>

            <div class="text-center mt-8">
                <button @click="kirimSuara" :disabled="!kandidatTerpilih"
                    class="bg-[#800000] text-white px-6 py-3 rounded disabled:opacity-50 hover:bg-red-900 transition">
                    Konfirmasi Pilihan
                </button>
            </div>
        </main>

        <FooterPemilih />
    </div>
</template>

<script>
import NavbarPemilih from '@/components/NavbarPemilih.vue'
import FooterPemilih from '@/components/FooterPemilih.vue'
import api from '@/service/api'
import Swal from 'sweetalert2'

export default {
    name: 'PemilihPage',
    components: {
        NavbarPemilih,
        FooterPemilih,
    },
    data() {
        return {
            kandidatList: [],
            kandidatTerpilih: null,
            kandidatTerpilihnama: null,
        }
    },
    computed: {
        mahasiswa() {
            return this.$store.state.storeMahasiswa.userMahasiswa?.data_mahasiswa || {}
        },
        tps() {
            return this.mahasiswa?.tps || {}
        }
    },
    mounted() {
        this.fetchAllKandidat()
        this.cekStatusTPS()
    },
    methods: {
        getfullpath(img) {
            return api.getFullpathImage(img)
        },
        async fetchAllKandidat() {
            try {
                const res = await api.GetAllKandidat()
                this.kandidatList = res.data
            } catch (err) {
                Swal.fire('Gagal', 'Gagal memuat data kandidat', 'error')
            }
        },
        pilihKandidat(kandidat) {
            this.kandidatTerpilih = kandidat.id_kandidat
            this.kandidatTerpilihnama = kandidat.nama_kandidat
        },
        async kirimSuara() {
            if (this.mahasiswa.sudah_memilih) {
                Swal.fire({
                    icon: 'info',
                    title: 'Sudah Memilih',
                    text: 'Anda sudah memberikan suara. Terima kasih!',
                })
                return
            }

            if (!this.tps.is_open) {
                Swal.fire({
                    icon: 'warning',
                    title: 'TPS Ditutup',
                    text: 'TPS Anda saat ini masih ditutup. Silakan tunggu dibuka.',
                })
                return
            }

            if (!this.kandidatTerpilih) return

            try {
                const payload = { kandidat_id: this.kandidatTerpilih }
                await api.PilihKandidat(payload)

                Swal.fire({
                    icon: 'success',
                    title: 'Berhasil Memilih!',
                    text: `Kamu telah memilih kandidat dengan Nama: ${this.kandidatTerpilihnama}`,
                }).then(() => {
                    this.$router.push('/dashboard')
                })
            } catch (err) {
                Swal.fire({
                    icon: 'error',
                    title: 'Gagal',
                    text: 'Terjadi kesalahan saat mengirim suara.',
                })
            }
        },
        cekStatusTPS() {
            if (!this.tps || Object.keys(this.tps).length === 0) return

            if (!this.tps.is_open) {
                Swal.fire({
                    icon: 'warning',
                    title: 'TPS Ditutup',
                    text: 'TPS Anda belum dibuka, harap menunggu.',
                })
            } else if (this.mahasiswa.sudah_memilih) {
                Swal.fire({
                    icon: 'info',
                    title: 'Sudah Memilih',
                    text: 'Anda sudah memilih. Terima kasih.',
                })
            }
        }
    }
}
</script>
