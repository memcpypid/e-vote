package routes

import (
	"E-vote/E-voteService/controllers"

	"github.com/gin-gonic/gin"
)

func Wilayah(api *gin.RouterGroup) {
	api.POST("/wilayah", controllers.CreateWilayah)       // Create Wilayah Admin
	api.PUT("/wilayah/:id", controllers.UpdateWilayah)    // Update Wilayah Admin
	api.GET("/wilayah", controllers.GetAllWilayah)        // Get User by ID
	api.GET("/wilayah/:id", controllers.GetWilayahByID)   // Get User by ID
	api.DELETE("/wilayah/:id", controllers.DeleteWilayah) // Delete Wilayah 	Admin
}
