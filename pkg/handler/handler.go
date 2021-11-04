package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/wellWINeo/MusicPlayerBackend/pkg/service"
)

type Handler struct {
	services *service.Service
	dataPath string
}

func NewHandler(services *service.Service, savePath string) *Handler {
	return &Handler{
		services: services,
		dataPath: savePath,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	// login group
	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", h.signIn)
		auth.POST("/sign-up", h.signUp)
		auth.POST("/verify", h.verify)
	}

	api := router.Group("/api", h.userIdentity)
	{
		// working with user profile
		users := api.Group("/users")
		{
			users.GET("/", h.getUser)
			users.PUT("/", h.updateUser)
			users.DELETE("/", h.deleteUser)
		}

		// api for user's library
		tracks := api.Group("/tracks", h.accessCheckTrack)
		{
			tracks.GET("/all", h.getAllTracks)
			tracks.GET("/download/:id", h.downloadTrack)
			tracks.POST("/upload/:id", h.uploadTrack)
			tracks.GET("/listen/:id/:part", h.loadTrackPart)
			tracks.GET("/:id", h.getTrack)
			tracks.PUT("/:id", h.updateTrack)
			tracks.POST("/", h.createTrack)
			tracks.DELETE("/:id", h.deleteTrack)
		}

		// api for user's playlists
		playlists := api.Group("/playlists", h.accessCheckPlaylist)
		{
			playlists.GET("/all", h.getAllPlaylists)

			playlists.GET("/:id", h.getPlaylist)
			playlists.PUT("/:id", h.updatePlaylist)
			playlists.POST("/", h.createPlaylist)
			playlists.DELETE("/:id", h.deletePlaylist)

			playlists.POST("/mod/:id", h.addToPlaylist)
			playlists.DELETE("/mod/:id", h.removeFromPlaylist)

		}

		likes := api.Group("/like", h.accessCheckTrack)
		{
			likes.POST("/:id", h.setLike)
			likes.DELETE("/:id", h.unsetLike)
			likes.GET("/", h.getAllLikes)
		}

		api.GET("/history", h.getHistory)
		api.POST("/buy", h.buyPremium)
	}

	return router
}
