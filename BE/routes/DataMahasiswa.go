package routes

import (
	"E-vote/E-voteService/controllers"

	"github.com/gin-gonic/gin"
)

func DatauserRoutes(api *gin.RouterGroup) {
	api.POST("/data-mahasiswa", controllers.CreateDataUserWithImage)     // Create User
	api.POST("/data-mahasiswa/import", controllers.ImportMahasiswaExcel) // Create User
	api.GET("/data-mahasiswa/export", controllers.ExportMahasiswaExcel)  // Create User
	api.PUT("/data-mahasiswa", controllers.UpdateUserDataAuth)           // Update User
	api.PUT("/data-mahasiswa/update-status/:id", controllers.UpdateStatusAnggota)
	// api.POST("/data-user", controllers.CreateDataUser)       // Create User
	api.GET("/data-mahasiswa", controllers.GetUserDataAuth)             // Get DataAnggota by ID
	api.GET("/data-mahasiswa/all", controllers.GetAllUserDataMahasiswa) // Get DataAnggota semua
	api.DELETE("/data-mahasiswa/:id", controllers.DeleteUserData)       // Delete User
}
