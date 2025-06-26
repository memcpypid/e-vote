package routes

import (
	"E-vote/E-voteService/controllers"

	"github.com/gin-gonic/gin"
)

func TpsRoutes(api *gin.RouterGroup) {
	api.POST("/tps", controllers.CreateTPS)
	api.POST("/tps/generate", controllers.DistribusiTPSUntukMahasiswa)
	api.PUT("/tps/:id", controllers.UpdateTPS)
	api.PUT("/tps/status/:id", controllers.UpdateTPSStatus)
	api.PUT("/tps/petugas/:id", controllers.ConvertToPetugasTPS) // Update User tps
	api.GET("/tps/:id", controllers.GetTPSByID)
	api.GET("/tps", controllers.GetAllTPS)
	api.GET("/tps/rekap", controllers.GetTPSWithPemilihSummary)
	api.GET("/tps/petugas", controllers.GetTPSInfoByPetugas)
	api.DELETE("/tps/:id", controllers.DeleteTPS)
}
