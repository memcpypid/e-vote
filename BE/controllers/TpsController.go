package controllers

import (
	"E-vote/E-voteService/config"
	"E-vote/E-voteService/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GET /tps
func GetAllTPS(c *gin.Context) {
	var tps []models.TPS
	if err := config.DB.Preload("Daerah").Preload("Wilayah").Find(&tps).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data TPS"})
		return
	}
	c.JSON(http.StatusOK, tps)
}

// GET /tps/:id
func GetTPSByID(c *gin.Context) {
	id := c.Param("id")
	var tps models.TPS

	if err := config.DB.Preload("Daerah").Preload("Wilayah").First(&tps, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "TPS tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, tps)
}

// POST /tps
func CreateTPS(c *gin.Context) {
	var input struct {
		NamaTPS   string `json:"nama_tps" binding:"required"`
		DaerahID  uint   `json:"daerah_id" binding:"required"`
		WilayahID uint   `json:"wilayah_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Buat TPS tanpa NoTPS dulu
	tps := models.TPS{
		NamaTPS:   input.NamaTPS,
		DaerahID:  &input.DaerahID,
		WilayahID: &input.WilayahID,
	}

	// Simpan ke database agar ID terisi
	if err := config.DB.Create(&tps).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat TPS"})
		return
	}

	// Format NoTPS sebagai 3 digit dari ID (misal 001)
	tps.NoTPS = fmt.Sprintf("%03d", tps.ID)

	// Update lagi NoTPS ke database
	if err := config.DB.Model(&tps).Update("no_tps", tps.NoTPS).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengatur No TPS"})
		return
	}

	c.JSON(http.StatusCreated, tps)
}

// PUT /tps/:id/status
func UpdateTPSStatus(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	var input struct {
		IsOpen bool `json:"is_open"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
		return
	}

	var tps models.TPS
	if err := config.DB.First(&tps, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "TPS tidak ditemukan"})
		return
	}

	tps.IsOpen = input.IsOpen
	if err := config.DB.Save(&tps).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui status TPS"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Status TPS berhasil diperbarui",
		"id_tps":   tps.ID,
		"is_open":  tps.IsOpen,
		"nama_tps": tps.NamaTPS,
	})
}

// PUT /tps/:id
func UpdateTPS(c *gin.Context) {
	id := c.Param("id")
	var tps models.TPS

	if err := config.DB.First(&tps, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "TPS tidak ditemukan"})
		return
	}

	var input struct {
		NamaTPS   string `json:"nama_tps"`
		NoTPS     string `json:"no_tps"`
		DaerahID  uint   `json:"id_daerah"`
		WilayahID uint   `json:"id_wilayah"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update field jika tidak kosong
	if input.NamaTPS != "" {
		tps.NamaTPS = input.NamaTPS
	}
	if input.NoTPS != "" {
		tps.NoTPS = input.NoTPS
	}
	if input.DaerahID != 0 {
		tps.DaerahID = &input.DaerahID
	}
	if input.WilayahID != 0 {
		tps.WilayahID = &input.WilayahID
	}

	if err := config.DB.Save(&tps).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui TPS"})
		return
	}
	c.JSON(http.StatusOK, tps)
}

// DELETE /tps/:id
func DeleteTPS(c *gin.Context) {
	id := c.Param("id")
	var tps models.TPS

	if err := config.DB.First(&tps, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "TPS tidak ditemukan"})
		return
	}

	if err := config.DB.Delete(&tps).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus TPS"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "TPS berhasil dihapus"})
}

// POST /tps/distribute-mahasiswa
func DistribusiTPSUntukMahasiswa(c *gin.Context) {
	var tpsList []models.TPS
	var mahasiswaList []models.DataMahasiswa

	// Ambil semua TPS yang tersedia
	if err := config.DB.Find(&tpsList).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data TPS"})
		return
	}
	if len(tpsList) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Belum ada data TPS untuk distribusi"})
		return
	}

	// Ambil semua mahasiswa yang belum memiliki TPS
	if err := config.DB.Where("tps_id IS NULL").Find(&mahasiswaList).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data mahasiswa"})
		return
	}

	if len(mahasiswaList) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "Semua mahasiswa sudah memiliki TPS"})
		return
	}

	// Bagi mahasiswa ke TPS secara merata (round-robin)
	for i, m := range mahasiswaList {
		tps := tpsList[i%len(tpsList)]
		m.TPSID = &tps.ID
		if err := config.DB.Save(&m).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Gagal menyimpan TPS untuk mahasiswa ID %d", m.ID)})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message":          "Mahasiswa berhasil didistribusikan ke TPS secara merata",
		"jumlah_mahasiswa": len(mahasiswaList),
		"jumlah_tps":       len(tpsList),
	})
}

// GET /tps/rekap
func GetTPSWithPemilihSummary(c *gin.Context) {
	type TPSRekap struct {
		ID         uint   `json:"id"`
		NamaTPS    string `json:"nama_tps"`
		NoTPS      string `json:"no_tps"`
		WilayahID  uint   `json:"wilayah_id"`
		DaerahID   uint   `json:"daerah_id"`
		Total      int64  `json:"total_pemilih"`
		SudahPilih int64  `json:"sudah_memilih"`
	}

	var results []TPSRekap
	if err := config.DB.Table("tps").
		Select("tps.id_tps, tps.nama_tps, tps.no_tps, tps.wilayah_id, tps.daerah_id, COUNT(data_mahasiswas.id_data_mahasiswa) as total, SUM(CASE WHEN data_mahasiswas.sudah_memilih THEN 1 ELSE 0 END) as sudah_pilih").
		Joins("LEFT JOIN data_mahasiswas ON data_mahasiswas.tps_id = tps.id_tps").
		Group("tps.id_tps").
		Scan(&results).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data rekap TPS"})
		return
	}

	c.JSON(http.StatusOK, results)
}

func ConvertToPetugasTPS(c *gin.Context) {
	userID := c.Param("id")
	var user models.AkunMahasiswa

	// Cari user berdasarkan ID
	if err := config.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User tidak ditemukan"})
		return
	}

	// Update role jadi 'tps'
	if user.Role == "tps" {
		user.Role = "mahasiswa"
	} else {
		user.Role = "tps"
	}

	if err := config.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengubah role user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Role user berhasil diubah menjadi petugas TPS",
		"user":    user,
	})
}

// GET /api/petugas/tps
func GetTPSInfoByPetugas(c *gin.Context) {
	userID, _ := c.Get("userID")

	var mahasiswa models.DataMahasiswa
	if err := config.DB.Preload("TPS").Where("user_id = ?", userID).First(&mahasiswa).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data TPS tidak ditemukan"})
		return
	}
	var total int64
	var sudah int64
	if err := config.DB.Model(&models.DataMahasiswa{}).Where("tps_id = ?", mahasiswa.TPSID).Count(&total).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghitung total pemilih"})
		return
	}
	if err := config.DB.Model(&models.DataMahasiswa{}).Where("tps_id = ? AND sudah_memilih = true", mahasiswa.TPSID).Count(&sudah).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghitung jumlah yang sudah memilih"})
		return
	}
	var hasilSuara []struct {
		ID    uint   `json:"id"`
		Nama  string `json:"nama"`
		Suara int64  `json:"suara"`
	}
	if err := config.DB.
		Table("kandidats").
		Select(`
		kandidats.id_kandidat AS id,
		kandidats.nama_kandidat AS nama,
		COUNT(data_mahasiswas.id_data_mahasiswa) AS suara`).
		Joins(`LEFT JOIN data_mahasiswas 
			ON kandidats.id_kandidat = data_mahasiswas.kandidat_id 
			AND data_mahasiswas.tps_id = ? 
			AND data_mahasiswas.sudah_memilih = true`, mahasiswa.TPSID).
		Group("kandidats.id_kandidat, kandidats.nama_kandidat").
		Scan(&hasilSuara).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil rekap suara kandidat"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id_tps":        mahasiswa.TPS.ID,
		"nama_tps":      mahasiswa.TPS.NamaTPS,
		"total_pemilih": total,
		"sudah_memilih": sudah,
		"is_open":       mahasiswa.TPS.IsOpen,
		"kandidat":      hasilSuara,
	})
}
