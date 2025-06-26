package controllers

import (
	"E-vote/E-voteService/config"
	"E-vote/E-voteService/models"

	"crypto/rand"
	"encoding/hex"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// CreateUser untuk membuat akun pengguna baru
func CreateUser(c *gin.Context) {
	var user models.AkunMahasiswa
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash password sebelum menyimpan ke database
	if err := user.HashPassword(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Terjadi kesalahan saat meng-hash password"})
		return
	}

	// Simpan user ke database
	if err := config.DB.Create(&user).Error; err != nil {
		// Menangani error jika terjadi pelanggaran constraint, seperti duplikat email
		if strings.Contains(err.Error(), "Duplicate entry") {
			c.JSON(http.StatusConflict, gin.H{"error": "NIM sudah terdaftar"})
			return
		}
		// Menangani error lain yang tidak spesifik
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Simpan user ke database
	c.JSON(http.StatusOK, gin.H{"message": "Berhasil Membuat Akun"})
}

// GetUser untuk mendapatkan detail pengguna berdasarkan ID
func GetUserById(c *gin.Context) {
	userID := c.Param("id")
	var user models.AkunMahasiswa

	if err := config.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// UpdateUser untuk memperbarui data pengguna
func UpdateUser(c *gin.Context) {
	userID := c.Param("id")
	var user models.AkunMahasiswa

	// Bind data dari request body
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Mencari pengguna berdasarkan ID
	if err := config.DB.Model(&user).Where("id = ?", userID).Updates(user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

// DeleteUser untuk menghapus pengguna berdasarkan ID
func DeleteUser(c *gin.Context) {
	userID := c.Param("id")
	var user models.AkunMahasiswa

	if err := config.DB.Delete(&user, userID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

type ResetPasswordRequest struct {
	Email string `json:"email"`
}

func GenerateResetToken() (string, error) {
	token := make([]byte, 32)
	_, err := rand.Read(token)
	if err != nil {
		return "", errors.New("failed to generate token")
	}
	return hex.EncodeToString(token), nil
}
