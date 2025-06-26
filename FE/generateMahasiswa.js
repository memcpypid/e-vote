const XLSX = require("xlsx");
const { faker } = require("@faker-js/faker");

// Atur lokal ke Indonesia
faker.locale = "id_ID";

const jumlahData = 10;
const dataMahasiswa = [];

for (let i = 0; i < jumlahData; i++) {
  const nim = 202210370311376 + i;
  const namaDepan = faker.person.firstName();
  const namaBelakang = faker.person.lastName();
  const noTPS = ""; // dikosongkan, bukan null atau false
  const sudahMemilih = false;
  const noHp = faker.phone.number("08##########");
  const alamat = faker.location.streetAddress();
  const tempatLahir = faker.location.city();
  const tanggalLahir = faker.date.birthdate({
    min: 18,
    max: 24,
    mode: "age",
  });
  const pekerjaan = faker.person.jobTitle();
  const status = faker.helpers.arrayElement(["Belum Menikah", "Menikah"]);
  const agama = faker.helpers.arrayElement([
    "Islam",
    "Kristen",
    "Katolik",
    "Hindu",
    "Budha",
  ]);
  const nik = faker.number
    .int({ min: 1000000000000000, max: 9999999999999999 })
    .toString();

  dataMahasiswa.push({
    No: i + 1,
    NIM: nim.toString(), // agar tidak diformat sebagai angka ribuan
    "Nama Depan": namaDepan,
    "Nama Belakang": namaBelakang,
    "No TPS": noTPS,
    "Sudah Memilih": sudahMemilih,
    "No HP": noHp,
    Alamat: alamat,
    "Tempat Lahir": tempatLahir,
    "Tanggal Lahir": tanggalLahir.toISOString().split("T")[0],
    Pekerjaan: pekerjaan,
    "Status Perkawinan": status,
    Agama: agama,
    NIK: nik,
  });
}

// Generate dan simpan Excel
const worksheet = XLSX.utils.json_to_sheet(dataMahasiswa);
const workbook = XLSX.utils.book_new();
XLSX.utils.book_append_sheet(workbook, worksheet, "Mahasiswa");

const outputPath = "mahasiswa_template_dummy.xlsx";
XLSX.writeFile(workbook, outputPath);

console.log("✔️  File", outputPath, "berhasil dibuat");
