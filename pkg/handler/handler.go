package handler

import "github.com/gin-gonic/gin"

type Handler struct{}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-in")
		auth.POST("/sign-up")
	}

	api := router.Group("/api")
	{
		users := api.Group("/users")
		{
			users.GET("/")
			users.PUT("/")
			users.POST("/")
			users.DELETE("/")
		}

		api.GET("/tracks_list")
		tracks := api.Group("/tracks")
		{
			tracks.GET("/")
			tracks.PUT("/")
			tracks.POST("/")
			tracks.DELETE("/")
		}

		api.POST("/like")
		api.DELETE("/like")

		api.GET("/history")
		api.POST("/history")
	}

	return router
}
