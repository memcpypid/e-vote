package routes

import (
	"E-vote/E-voteService/handlers"
	"E-vote/E-voteService/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	// Tambahkan middleware CSRF token secara global
	// router.Use(middleware.CSRFTokenMiddleware())
	// router.Use(middleware.MonitoringMiddleware())
	protected := router.Group("/api")
	statis := router.Group("/")
	ajax := router.Group("/api")
	userRoutes(protected)
	protected.Use(middleware.VerifyJWT())
	protected.Use(middleware.RequireAJAX())
	DatauserRoutes(protected)

	Daerah(protected)
	Wilayah(protected)
	TpsRoutes(protected)
	KandidatRoutes(protected)
	router.Static("/js", "./dist/js")
	router.Static("/css", "./dist/css")
	router.Static("/img", "./dist/img")
	router.Static("/fonts", "./dist/fonts")
	router.Static("/font", "./dist/font")
	router.NoRoute(func(c *gin.Context) {
		c.File("./dist/index.html")
	})

	// statis.Use(middleware.StaticFileMiddleware())
	// {

	statis.Static("/uploads", "./uploads")
	statis.Static("/assets", "./assets")
	// }
	ajax.POST("/auth/login", handlers.Login)
	ajax.POST("/auth/logout", handlers.Logout)
	// ajax.POST("/auth/reset-password-token", controllers.RequestPasswordReset)
	// ajax.GET("/auth/check-password-token/:token", controllers.CheckTokenValidity)
	// ajax.POST("/auth/reset-password/:token", controllers.ResetPassword)
	ajax.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Health Oke!",
		})
	})
}
