package routes

import (
	"E-vote/E-voteService/controllers"

	"github.com/gin-gonic/gin"
)

func KandidatRoutes(api *gin.RouterGroup) {
	api.POST("/kandidat", controllers.CreateKandidat)
	api.PUT("/kandidat/pilih", controllers.PilihKandidat)
	api.PUT("/kandidat/:id", controllers.UpdateKandidat)
	api.GET("/kandidat/:id", controllers.GetKandidatByID)
	api.GET("/kandidat", controllers.GetAllKandidat)
	api.GET("/kandidat/vote-count", controllers.GetVoteCountPerKandidat)
	api.DELETE("/kandidat/:id", controllers.DeleteKandidat)
}
