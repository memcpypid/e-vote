package controllers

import (
	"E-vote/E-voteService/config"
	"E-vote/E-voteService/models"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// GET /kandidat
func GetAllKandidat(c *gin.Context) {
	var kandidat []models.Kandidat
	if err := config.DB.Find(&kandidat).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data kandidat"})
		return
	}
	c.JSON(http.StatusOK, kandidat)
}

// GET /kandidat/:id
func GetKandidatByID(c *gin.Context) {
	id := c.Param("id")
	var kandidat models.Kandidat

	if err := config.DB.First(&kandidat, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Kandidat tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, kandidat)
}

// POST /kandidat
func CreateKandidat(c *gin.Context) {
	var input models.Kandidat

	// Parse multipart form
	if err := c.Request.ParseMultipartForm(10 << 20); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Gagal membaca form data"})
		return
	}

	input.NamaKandidat = c.PostForm("nama_kandidat")
	input.NIMKandidat = c.PostForm("nim_kandidat")
	input.NamaPasangan = stringPointer(c.PostForm("nama_pasangan"))
	input.NIMPasangan = stringPointer(c.PostForm("nim_pasangan"))
	input.Visi = c.PostForm("visi")
	input.Misi = c.PostForm("misi")

	// Handle Foto
	file, err := c.FormFile("foto")
	if err == nil {
		ext := strings.ToLower(filepath.Ext(file.Filename))
		filename := fmt.Sprintf("foto_%d%s", time.Now().UnixNano(), ext)
		path := filepath.Join("uploads", filename)
		if err := c.SaveUploadedFile(file, path); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan foto"})
			return
		}
		input.Foto = &filename
	}

	if err := config.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan data kandidat"})
		return
	}

	c.JSON(http.StatusCreated, input)
}

func UpdateKandidat(c *gin.Context) {
	id := c.Param("id")
	var kandidat models.Kandidat

	if err := config.DB.First(&kandidat, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Kandidat tidak ditemukan"})
		return
	}

	if err := c.Request.ParseMultipartForm(10 << 20); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Gagal membaca form data"})
		return
	}

	// Update fields if available
	if v := c.PostForm("nama_kandidat"); v != "" {
		kandidat.NamaKandidat = v
	}
	if v := c.PostForm("nim_kandidat"); v != "" {
		kandidat.NIMKandidat = v
	}
	if v := c.PostForm("nama_pasangan"); v != "" {
		kandidat.NamaPasangan = stringPointer(v)
	}
	if v := c.PostForm("nim_pasangan"); v != "" {
		kandidat.NIMPasangan = stringPointer(v)
	}
	if v := c.PostForm("visi"); v != "" {
		kandidat.Visi = v
	}
	if v := c.PostForm("misi"); v != "" {
		kandidat.Misi = v
	}

	// Handle Foto baru jika ada
	file, err := c.FormFile("foto")
	if err == nil {
		ext := strings.ToLower(filepath.Ext(file.Filename))
		filename := fmt.Sprintf("foto_%d%s", time.Now().UnixNano(), ext)
		path := filepath.Join("uploads", filename)
		if err := c.SaveUploadedFile(file, path); err == nil {
			// Hapus foto lama jika ada
			if kandidat.Foto != nil {
				os.Remove(filepath.Join("uploads", *kandidat.Foto))
			}
			kandidat.Foto = &filename
		}
	}

	if err := config.DB.Save(&kandidat).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui data kandidat"})
		return
	}

	c.JSON(http.StatusOK, kandidat)
}

// DELETE /kandidat/:id
func DeleteKandidat(c *gin.Context) {
	id := c.Param("id")
	var kandidat models.Kandidat

	if err := config.DB.First(&kandidat, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Kandidat tidak ditemukan"})
		return
	}

	if err := config.DB.Delete(&kandidat).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus data kandidat"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Kandidat berhasil dihapus"})
}

func GetVoteCountPerKandidat(c *gin.Context) {
	type HasilSuara struct {
		IDKandidat   uint   `json:"id_kandidat"`
		NamaKandidat string `json:"nama_kandidat"`
		NIMKandidat  string `json:"nim_kandidat"`
		JumlahSuara  int64  `json:"jumlah_suara"`
	}

	var results []HasilSuara

	if err := config.DB.
		Table("kandidats").
		Select(`kandidats.id_kandidat, kandidats.nama_kandidat, kandidats.nim_kandidat, 
		        COUNT(data_mahasiswas.kandidat_id) as jumlah_suara`).
		Joins(`LEFT JOIN data_mahasiswas ON data_mahasiswas.kandidat_id = kandidats.id_kandidat`).
		Group("kandidats.id_kandidat, kandidats.nama_kandidat, kandidats.nim_kandidat").
		Scan(&results).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghitung suara per kandidat"})
		return
	}

	c.JSON(http.StatusOK, results)
}

// PUT /kandidat/pilih
func PilihKandidat(c *gin.Context) {
	userID, exists := c.Get("userID")
	// var userValid bool = true
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authorized"})
		return
	}
	id := userID.(uint)
	var input struct {
		KandidatID uint `json:"kandidat_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid", "detail": err.Error()})
		return
	}

	var mahasiswa models.DataMahasiswa

	// Cari mahasiswa berdasarkan user_id
	if err := config.DB.Where("user_id = ?", id).First(&mahasiswa).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Mahasiswa tidak ditemukan"})
		return
	}

	// Cek apakah sudah memilih
	if mahasiswa.SudahMemilih {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Mahasiswa sudah memilih sebelumnya"})
		return
	}

	// Update kandidat_id dan status sudah_memilih
	mahasiswa.KandidatID = &input.KandidatID
	mahasiswa.SudahMemilih = true

	if err := config.DB.Save(&mahasiswa).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui data pemilih"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pemilihan berhasil disimpan"})
}
func stringPointer(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}
