package controllers

import (
	"E-vote/E-voteService/config"
	"E-vote/E-voteService/models"

	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

// CreateDataUserWithImage - Untuk Anggota
func CreateDataUserWithImage(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authorized"})
		return
	}
	id := userID.(uint)
	NamaLengkap := c.PostForm("nama_lengkap")
	Alamat := c.PostForm("alamat")
	TanggalLahirStr := c.PostForm("tanggalLahir")
	NIK := c.PostForm("nik")
	TempatLahir := c.PostForm("tempatLahir")
	Pekerjaan := c.PostForm("pekerjaan")
	StatusPerkawinan := c.PostForm("statusPerkawinan")
	Agama := c.PostForm("agama")

	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "JabatanStrukturalID Tidak Valid"})
	// 	return
	// }

	// Parsing tanggal lahir
	var TanggalLahir time.Time
	var err error
	if TanggalLahirStr != "" {
		TanggalLahir, err = time.Parse("2006-01-02", TanggalLahirStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Format tanggal lahir tidak valid"})
			return
		}
	}

	// Buat data anggota
	user := models.DataMahasiswa{
		UserID:           uint(id),
		NamaDepan:        NamaLengkap,
		Alamat:           Alamat,
		TanggalLahir:     TanggalLahir,
		NIK:              NIK,
		TempatLahir:      TempatLahir,
		Pekerjaan:        Pekerjaan,
		StatusPerkawinan: StatusPerkawinan,
		Agama:            Agama,
	}

	// Mulai transaksi database
	tx := config.DB.Begin()
	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Proses file upload
	files, err := c.MultipartForm()
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"error": "Gagal membaca file dari form-data"})
		return
	}

	fileFoto3x4, exists := files.File["file[]"]
	keterangan := c.PostFormArray("keterangan[]")

	// Validasi jumlah file dengan jumlah keterangan
	if exists && len(fileFoto3x4) > 0 {
		if len(keterangan) != len(fileFoto3x4) {
			tx.Rollback()
			c.JSON(http.StatusBadRequest, gin.H{"error": "Jumlah keterangan tidak sesuai dengan jumlah file"})
			return
		}

		allowedExtensions := []string{".jpg", ".jpeg", ".png", ".JPG", ".JPEG", ".PNG"}
		maxFileSize := int64(5 * 1024 * 1024) // 5MB
		var imageRecords []models.ImageDataMahasiswa

		for i, file := range fileFoto3x4 {
			// Validasi ukuran file
			if file.Size > maxFileSize {
				tx.Rollback()
				c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Ukuran file %s terlalu besar (max 5MB)", file.Filename)})
				return
			}

			// Validasi format file
			ext := filepath.Ext(file.Filename)
			isAllowed := false
			for _, allowedExt := range allowedExtensions {
				if ext == allowedExt {
					isAllowed = true
					break
				}
			}
			if !isAllowed {
				tx.Rollback()
				c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Format file %s tidak didukung", file.Filename)})
				return
			}

			// Buat folder jika belum ada
			folderPath := fmt.Sprintf("./uploads/data-anggota/%s", keterangan[i])
			if err := os.MkdirAll(folderPath, os.ModePerm); err != nil {
				tx.Rollback()
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat direktori"})
				return
			}

			// Simpan file dengan nama unik
			newFileName := fmt.Sprintf("%d_%s%s", user.ID, time.Now().Format("20060102150405"), ext)
			destination := fmt.Sprintf("%s/%s", folderPath, newFileName)

			if err := c.SaveUploadedFile(file, destination); err != nil {
				tx.Rollback()
				c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Gagal menyimpan file %s", file.Filename)})
				return
			}

			// Tambahkan ke daftar image yang akan disimpan
			imageRecords = append(imageRecords, models.ImageDataMahasiswa{
				DataUserID: user.ID,
				ImageURL:   destination,
				Keterangan: keterangan[i],
			})
		}

		// Simpan semua data gambar ke database dalam satu batch insert
		if err := tx.Create(&imageRecords).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan data gambar ke database"})
			return
		}
	}

	// Commit transaksi jika semua berhasil
	tx.Commit()

	// Sukses menyimpan data
	c.JSON(http.StatusOK, gin.H{"message": "Data dan gambar berhasil disimpan"})
}

// UpdateUserDataAuth - Untuk Anggota
func UpdateUserDataAuth(c *gin.Context) {
	AnggotaID, exists := c.Get("dataAnggotaID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authorized"})
		return
	}
	anggotaId := AnggotaID.(uint)
	// Cek apakah data anggota ada
	var anggota models.DataMahasiswa
	if err := config.DB.First(&anggota, anggotaId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data anggota tidak ditemukan"})
		return
	}

	// Map untuk menyimpan field yang akan diperbarui
	updateFields := make(map[string]interface{})

	// Cek setiap field secara individu, hanya update jika ada data baru
	if daerahID := c.PostForm("daerahId"); daerahID != "" {
		if val, err := strconv.Atoi(daerahID); err == nil {
			updateFields["daerah_id"] = uint(val)
		}
	}
	if wilayahID := c.PostForm("wilayahId"); wilayahID != "" {
		if val, err := strconv.Atoi(wilayahID); err == nil {
			updateFields["wilayah_id"] = uint(val)
		}
	}
	if jabatanStrukturalID := c.PostForm("jabatanStrukturalId"); jabatanStrukturalID != "" {
		if val, err := strconv.Atoi(jabatanStrukturalID); err == nil {
			updateFields["jabatan_struktural_id"] = uint(val)
		}
	}
	if namaLengkap := c.PostForm("nama_lengkap"); namaLengkap != "" {
		updateFields["nama_lengkap"] = namaLengkap
	}
	if alamat := c.PostForm("alamat"); alamat != "" {
		updateFields["alamat"] = alamat
	}
	if tanggalLahir := c.PostForm("tanggalLahir"); tanggalLahir != "" {
		if parsedDate, err := time.Parse("2006-01-02", tanggalLahir); err == nil {
			updateFields["tanggal_lahir"] = parsedDate
		}
	}
	if nik := c.PostForm("nik"); nik != "" {
		updateFields["nik"] = nik
	}
	if tempatLahir := c.PostForm("tempatLahir"); tempatLahir != "" {
		updateFields["tempat_lahir"] = tempatLahir
	}
	if pekerjaan := c.PostForm("pekerjaan"); pekerjaan != "" {
		updateFields["pekerjaan"] = pekerjaan
	}
	if statusPerkawinan := c.PostForm("statusPerkawinan"); statusPerkawinan != "" {
		updateFields["status_perkawinan"] = statusPerkawinan
	}
	if agama := c.PostForm("agama"); agama != "" {
		updateFields["agama"] = agama
	}
	if AlamatKantor := c.PostForm("alamatkantor"); AlamatKantor != "" {
		updateFields["alamat_kantor"] = AlamatKantor
	}

	// Jika ada data yang perlu diupdate, lakukan update
	if len(updateFields) > 0 {
		if err := config.DB.Model(&anggota).Updates(updateFields).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui data anggota"})
			return
		}
	}
	// ======================= UPDATE GAMBAR ======================= //
	files, err := c.MultipartForm()
	if err == nil {
		fileFoto3x4, exists := files.File["file[]"]
		keterangan := c.PostFormArray("keterangan[]")

		if exists && len(fileFoto3x4) > 0 {
			allowedExtensions := []string{".jpg", ".jpeg", ".png", ".JPG", ".JPEG", ".PNG"}
			maxFileSize := int64(10 * 1024 * 1024) // 20MB

			// Mulai transaksi database
			tx := config.DB.Begin()
			defer func() {
				if r := recover(); r != nil {
					tx.Rollback()
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Terjadi kesalahan saat memperbarui data"})
				}
			}()

			for i, file := range fileFoto3x4 {
				// Validasi ukuran file
				if file.Size > maxFileSize {
					c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Ukuran file %s terlalu besar (max 5MB)", file.Filename)})
					return
				}
				// Validasi format file
				ext := filepath.Ext(file.Filename)
				isAllowed := false
				for _, allowedExt := range allowedExtensions {
					if ext == allowedExt {
						isAllowed = true
						break
					}
				}
				if !isAllowed {
					c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Format file %s tidak didukung", file.Filename)})
					return
				}
				// Buat folder jika belum ada
				folderPath := fmt.Sprintf("./uploads/data-anggota/%s", keterangan[i])
				if err := os.MkdirAll(folderPath, os.ModePerm); err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat direktori"})
					return
				}

				// Simpan file dengan nama unik
				newFileName := fmt.Sprintf("%d_%s%s", anggotaId, time.Now().Format("20060102150405"), ext)
				destination := fmt.Sprintf("%s/%s", folderPath, newFileName)

				if err := c.SaveUploadedFile(file, destination); err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Gagal menyimpan file %s", file.Filename)})
					return
				}

				// Cek apakah gambar sudah ada berdasarkan `anggotaId` dan `keterangan`
				var existingImage models.ImageDataMahasiswa
				result := tx.Where("data_user_id = ? AND keterangan = ?", anggotaId, keterangan[i]).First(&existingImage)

				if result.RowsAffected > 0 {
					// Jika gambar sudah ada, hapus file lama
					if err := os.Remove(existingImage.ImageURL); err != nil {
						log.Println("Gagal menghapus file lama:", err)
					}

					// Update data gambar di database
					if err := tx.Model(&existingImage).Updates(models.ImageDataMahasiswa{
						ImageURL:   destination,
						Keterangan: keterangan[i],
					}).Error; err != nil {
						tx.Rollback()
						c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui data gambar di database"})
						return
					}
				} else {
					// Jika belum ada, tambahkan gambar baru
					newImage := models.ImageDataMahasiswa{
						DataUserID: anggotaId,
						ImageURL:   destination,
						Keterangan: keterangan[i],
					}

					if err := tx.Create(&newImage).Error; err != nil {
						tx.Rollback()
						c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan data gambar ke database"})
						return
					}
				}
			}

			// Commit transaksi jika semua operasi berhasil
			tx.Commit()
		}
	}
	// Berhasil memperbarui data
	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diperbarui"})
}

// GetUSerdata - Untuk Anggota/admin
func GetUserDataAuth(c *gin.Context) {
	userID, exists := c.Get("userID")
	var userValid bool = true
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authorized"})
		return
	}
	id := userID.(uint)
	var data models.DataMahasiswa
	var user models.AkunMahasiswa
	// Fetch user by ID and include related ImageUser records
	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data user Tidak Ditemukan"})
		userValid = false
		return
	}
	if err := config.DB.Preload("ImageUsers").Preload("TPS").Preload("TPS.Wilayah").Preload("TPS.Daerah").Where(models.DataMahasiswa{UserID: uint(id)}).First(&data).Error; err != nil {
		if !userValid {
			c.JSON(http.StatusNotFound, gin.H{"error": "Data Anggota Tidak Ditemukan"})
			return
		}
	}
	users := gin.H{"role": user.Role, "nim": user.NIM}
	c.JSON(http.StatusOK, gin.H{"user": users, "data_mahasiswa": data})
}

// DeleteUserData - Untuk-admin
func UpdateStatusAnggota(c *gin.Context) {
	// Ambil ID dari parameter
	id := c.Param("id")
	var user models.DataMahasiswa

	// Cek apakah ID valid (harus berupa angka)
	anggotaID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	// Fetch data anggota dari database
	if err := config.DB.First(&user, anggotaID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Anggota tidak ditemukan"})
		return
	}

	// Ambil status baru dari request body
	var requestBody struct {
		Status     string `json:"status"`
		Keterangan string `json:"keterangan"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format JSON tidak valid"})
		return
	}

	// Validasi status yang diperbolehkan
	validStatuses := map[string]bool{
		"PENDING": true,
		"SUCCESS": true,
		"CANCEL":  true,
	}

	if !validStatuses[requestBody.Status] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Status tidak valid. Gunakan: PENDING, SUCCESS, atau CANCEL"})
		return
	}

	// Update status anggota di database
	if err := config.DB.Model(&user).Update("status", requestBody.Status).Update("keterangan", requestBody.Keterangan).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui status anggota"})
		return
	}

	// Response berhasil
	c.JSON(http.StatusOK, gin.H{
		"message": "Status anggota berhasil diperbarui",
		"user": map[string]interface{}{
			"id":            user.ID,
			"nama_depan":    user.NamaDepan,
			"nama_belakang": user.NamaBelakang,
			"status":        requestBody.Status,
			"updated":       time.Now(),
		},
	})
}

// DeleteUserData - Untuk Admin
func DeleteUserData(c *gin.Context) {
	id := c.Param("id")
	var user models.DataMahasiswa
	// Fetch user by ID
	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	// Delete all associated images
	if err := config.DB.Where("data_user_id = ?", user.ID).Delete(&models.ImageDataMahasiswa{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting images"})
		return
	}
	// Delete the user from database
	if err := config.DB.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User and associated images deleted successfully"})
}

// GetAllUserDataMahasiswa - Untuk KPU melihat semua data mahasiswa dan akun terkait
func GetAllUserDataMahasiswa(c *gin.Context) {
	var userList []models.DataMahasiswa

	err := config.DB.
		Preload("User").       // Relasi ke AkunMahasiswa
		Preload("TPS").        // Relasi ke TPS
		Preload("Kandidat").   // Relasi ke kandidat yang dipilih
		Preload("ImageUsers"). // Relasi ke gambar pengguna
		Joins("JOIN akun_mahasiswas ON akun_mahasiswas.id_user = data_mahasiswas.user_id").
		Where("akun_mahasiswas.role IN ?", []string{"mahasiswa", "tps"}).
		Find(&userList).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data mahasiswa"})
		return
	}

	c.JSON(http.StatusOK, userList)
}

func ExportMahasiswaExcel(c *gin.Context) {
	var mahasiswa []models.DataMahasiswa
	if err := config.DB.Preload("User").Preload("TPS").Find(&mahasiswa).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data mahasiswa"})
		return
	}

	f := excelize.NewFile()
	sheet := "Mahasiswa"
	f.NewSheet(sheet)

	// Header
	headers := []string{"No", "NIM", "Nama", "No TPS", "Sudah Memilih", "No HP", "Alamat", "Tempat Lahir", "Tanggal Lahir", "Pekerjaan", "Status Perkawinan", "Agama"}
	for i, h := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheet, cell, h)
	}

	// Data
	for i, m := range mahasiswa {
		row := i + 2
		data := []interface{}{
			i + 1,
			m.User.NIM,
			fmt.Sprintf("%s %s", m.NamaDepan, m.NamaBelakang),
			func() interface{} {
				if m.TPS != nil {
					return m.TPS.NoTPS
				}
				return "-"
			}(),
			func() string {
				if m.SudahMemilih {
					return "Ya"
				}
				return "Belum"
			}(),
			m.NoHp,
			m.Alamat,
			m.TempatLahir,
			m.TanggalLahir.Format("2006-01-02"),
			m.Pekerjaan,
			m.StatusPerkawinan,
			m.Agama,
		}

		for j, v := range data {
			cell, _ := excelize.CoordinatesToCellName(j+1, row)
			f.SetCellValue(sheet, cell, v)
		}
	}

	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header("Content-Disposition", "attachment; filename=mahasiswa.xlsx")
	c.Header("File-Name", "mahasiswa.xlsx")
	f.Write(c.Writer)
}
func ImportMahasiswaExcel(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File tidak ditemukan"})
		return
	}

	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuka file"})
		return
	}
	defer src.Close()

	f, err := excelize.OpenReader(src)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format file tidak valid"})
		return
	}

	rows, err := f.GetRows("Mahasiswa")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Sheet 'Mahasiswa' tidak ditemukan"})
		return
	}

	for i, row := range rows {
		// Skip header & validasi minimal jumlah kolom
		if i == 0 || len(row) < 13 {
			continue
		}

		// Parse NIM
		nimStr := strings.ReplaceAll(row[1], ",", "") // buang koma jika ada
		nim, err := strconv.Atoi(nimStr)
		if err != nil {
			continue // skip jika gagal
		}

		// Parse No TPS (boleh kosong)
		// var noTPS *int
		// if strings.TrimSpace(row[4]) != "" {
		// 	no, err := strconv.Atoi(row[4])
		// 	if err == nil {
		// 		noTPS = &no
		// 	}
		// }

		// Parse Sudah Memilih
		sudahMemilih := strings.ToLower(row[5]) == "true" || strings.ToLower(row[5]) == "ya"

		// Parse Tanggal Lahir
		tanggalLahir, err := time.Parse("2006-01-02", row[9])
		if err != nil {
			continue // skip jika tanggal invalid
		}

		// Buat akun mahasiswa
		user := models.AkunMahasiswa{
			NIM:  nim,
			PIC:  "default123", // password default
			Role: "mahasiswa",
		}
		user.HashPassword()
		config.DB.Create(&user)

		// Simpan data mahasiswa
		data := models.DataMahasiswa{
			UserID:           user.ID,
			TPSID:            nil, // bisa di-set dari NoTPS jika kamu punya relasi ID TPS
			NoHp:             row[6],
			Alamat:           row[7],
			TempatLahir:      row[8],
			TanggalLahir:     tanggalLahir,
			Pekerjaan:        row[10],
			StatusPerkawinan: row[11],
			Agama:            row[12],
			NamaDepan:        row[2],
			NamaBelakang:     row[3],
			SudahMemilih:     sudahMemilih,
			NIK:              row[13],
		}

		config.DB.Create(&data)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Berhasil import data mahasiswa"})
}
