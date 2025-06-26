package controllers

import (
	"E-vote/E-voteService/config"
	"E-vote/E-voteService/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

type DaerahCreate struct {
	NamaDaerah string `json:"nama_daerah_tps"`
	WilayahID  uint   `json:"id_wilayah"`
}
type DaerahSimpleResponse struct {
	IDDaerah      uint   `json:"id_daerah"`
	NamaDaerahTPS string `json:"nama_daerah_tps"`
	IDWilayah     uint   `json:"id_wilayah"`
}

// CreateDaerah handles the creation of a new Daerah
func CreateDaerah(c *gin.Context) {
	var daerah DaerahCreate
	// Bind the incoming JSON to the daerah struct
	if err := c.ShouldBindJSON(&daerah); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the wilayah exists using the provided kode_wilayah
	var wilayah models.WilayahTPS
	if err := config.DB.Where("id_wilayah = ?", daerah.WilayahID).First(&wilayah).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Wilayah not found"})
		return
	}

	// Check if the Daerah name already exists in the same Wilayah
	var existingDaerah models.DaerahTPS
	if err := config.DB.Where("nama_daerah = ? AND wilayah_id = ?", daerah.NamaDaerah, wilayah.ID).First(&existingDaerah).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Daerah name must be unique within this Wilayah"})
		return
	}

	// Create the new Daerah
	newDaerah := models.DaerahTPS{
		NamaDaerah: daerah.NamaDaerah,
		WilayahID:  wilayah.ID,
	}

	// Now create the daerah in the database
	if result := config.DB.Create(&newDaerah); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// Respond with the created daerah
	c.JSON(http.StatusOK, gin.H{"data": newDaerah})
}

// Get All Daerah
func GetAllDaerah(c *gin.Context) {
	var daerah []models.DaerahTPS
	if result := config.DB.Find(&daerah); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	response := make([]DaerahSimpleResponse, 0)
	for _, d := range daerah {
		response = append(response, DaerahSimpleResponse{
			IDDaerah:      d.ID,
			NamaDaerahTPS: d.NamaDaerah,
			IDWilayah:     d.WilayahID,
		})
	}

	c.JSON(http.StatusOK, response)
}

// Get Daerah by ID
func GetDaerahByID(c *gin.Context) {
	id := c.Param("id")
	var daerah models.DaerahTPS
	if result := config.DB.Preload("Wilayah").First(&daerah, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Daerah not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": daerah})
}
func GetDaerahByIDWilayah(c *gin.Context) {
	id := c.Param("id")
	var daerah []models.DaerahTPS
	if result := config.DB.Preload("Wilayah").Where("wilayah_id = ?", id).Find(&daerah); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Daerah not found"})
		return
	}

	c.JSON(http.StatusOK, daerah)
}

// Update Daerah
func UpdateDaerah(c *gin.Context) {
	id := c.Param("id")
	var daerah models.DaerahTPS
	if result := config.DB.First(&daerah, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Daerah not found"})
		return
	}

	// Bind the incoming JSON to the daerah struct
	if err := c.ShouldBindJSON(&daerah); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the Daerah name already exists in the same Wilayah
	// var existingDaerah models.Daerah
	// if err := config.DB.Where("nama_daerah = ? AND wilayah_id = ?", daerah.NamaDaerah, daerah.WilayahID).First(&existingDaerah).Error; err == nil {
	// 	c.JSON(http.StatusConflict, gin.H{"error": "Daerah name must be unique within this Wilayah"})
	// 	return
	// }

	// Update the daerah in the database
	if result := config.DB.Save(&daerah); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": daerah})
}

// Delete Daerah
func DeleteDaerah(c *gin.Context) {
	id := c.Param("id")
	if result := config.DB.Delete(&models.DaerahTPS{}, id); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Daerah deleted"})
}
