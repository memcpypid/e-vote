package controllers

import (
	"E-vote/E-voteService/config"
	"E-vote/E-voteService/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DaerahResponse struct {
	ID         uint   `json:"id_daerah"`
	NamaDaerah string `json:"nama_daerah"`
	WilayahID  uint   `json:"wilayahID"`
}

type WilayahResponse struct {
	IDWilayah   uint             `json:"id_wilayah"`
	NamaWilayah string           `json:"nama_wilayah_tps"`
	Daerah      []DaerahResponse `json:"daerah"`
}

// Create Wilayah
func CreateWilayah(c *gin.Context) {
	var wilayah models.WilayahTPS
	if err := c.ShouldBindJSON(&wilayah); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the Wilayah name already exists
	var existingWilayah models.WilayahTPS
	if err := config.DB.Where("nama_wilayah = ?", wilayah.NamaWilayah).First(&existingWilayah).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Wilayah name must be unique"})
		return
	}

	// Create the new Wilayah
	if result := config.DB.Create(&wilayah); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": wilayah})
}

// Get All Wilayah
func GetAllWilayah(c *gin.Context) {
	var wilayah []models.WilayahTPS
	// Memuat semua wilayah dan daerah yang terkait
	if result := config.DB.Preload("Daerah").Find(&wilayah); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// Menyusun ulang respons dengan menghilangkan relasi 'wilayah' pada setiap 'daerah'
	wilayahResponses := make([]WilayahResponse, 0)
	for _, w := range wilayah {
		var wilayahResponse WilayahResponse
		wilayahResponse.IDWilayah = w.ID
		wilayahResponse.NamaWilayah = w.NamaWilayah

		// Memetakan data daerah tanpa relasi 'wilayah'
		for _, daerah := range w.Daerah {
			wilayahResponse.Daerah = append(wilayahResponse.Daerah, DaerahResponse{
				ID:         daerah.ID,
				NamaDaerah: daerah.NamaDaerah,
				WilayahID:  daerah.WilayahID,
			})
		}

		// Menambahkan wilayahResponse ke dalam array
		wilayahResponses = append(wilayahResponses, wilayahResponse)
	}

	// Mengembalikan data wilayah dan daerah yang telah dimodifikasi
	c.JSON(http.StatusOK, wilayahResponses)
}

func GetWilayahByID(c *gin.Context) {
	id := c.Param("id")
	var wilayah models.WilayahTPS

	// Memuat wilayah dan daerah yang terkait
	if result := config.DB.Preload("Daerah").First(&wilayah, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Wilayah not found"})
		return
	}
	var wilayahResponse WilayahResponse
	wilayahResponse.IDWilayah = wilayah.ID
	wilayahResponse.NamaWilayah = wilayah.NamaWilayah
	for _, daerah := range wilayah.Daerah {
		wilayahResponse.Daerah = append(wilayahResponse.Daerah, DaerahResponse{
			ID:         daerah.ID,
			NamaDaerah: daerah.NamaDaerah,
			WilayahID:  daerah.WilayahID,
		})
	}
	c.JSON(http.StatusOK, wilayahResponse)
}

// Update Wilayah
func UpdateWilayah(c *gin.Context) {
	id := c.Param("id")
	var wilayah models.WilayahTPS
	if result := config.DB.First(&wilayah, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Wilayah not found"})
		return
	}

	// Bind the incoming JSON to the wilayah struct
	if err := c.ShouldBindJSON(&wilayah); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the wilayah in the database
	if result := config.DB.Save(&wilayah); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, wilayah)
}

// Delete Wilayah
func DeleteWilayah(c *gin.Context) {
	id := c.Param("id")
	if result := config.DB.Delete(&models.WilayahTPS{}, id); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Wilayah deleted"})
}
