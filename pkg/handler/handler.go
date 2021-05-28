package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/wellWINeo/MusicPlayerBackend/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", h.signIn)
		auth.POST("/sign-up", h.signUp)
		auth.POST("/verify", h.verify)
	}

	api := router.Group("/api", h.userIdentity)
	{
		users := api.Group("/users")
		{
			users.GET("/", h.getUser)
			users.PUT("/", h.updateUser)
			users.DELETE("/", h.deleteUser)
		}

		tracks := api.Group("/tracks")
		{
			tracks.GET("/all", h.getAllTracks)
			tracks.GET("/:id/download", h.downloadTrack)
			tracks.GET("/:id", h.getTrack)
			tracks.PUT("/:id", h.updateTrack)
			tracks.POST("/", h.createTrack)
			tracks.DELETE("/:id", h.deleteTrack)
		}

		playlists := api.Group("/playlists")
		{
			playlists.GET("/all", h.getAllPlaylists)

			playlists.GET("/:id", h.getPlaylist)
			playlists.PUT("/:id", h.updatePlaylist)
			playlists.POST("/", h.createPlaylist)
			playlists.DELETE("/:id", h.deletePlaylist)

		}

		likes := api.Group("/like")
		{
			likes.POST("/:id", h.setLike)
			likes.DELETE("/:id", h.unsetLike)
		}

		history := api.Group("/history")
		{
			history.POST("/:id", h.addToHistory)
			history.GET("/", h.getHistory)
		}
	}

	return router
}
