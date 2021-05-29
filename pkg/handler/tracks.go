package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wellWINeo/MusicPlayerBackend"
)

func (h *Handler) getAllTracks(c *gin.Context) {

}

func (h *Handler) getTrack(c *gin.Context) {

}

func (h *Handler) updateTrack(c *gin.Context) {
	id, ok := c.Get(userIdCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "Can't get userId")
		return
	}

	var input MusicPlayerBackend.Track
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
}

func (h *Handler) createTrack(c *gin.Context) {

}

func (h *Handler) deleteTrack(c *gin.Context) {

}

func (h *Handler) uploadTrack(c *gin.Context) {

}

func (h *Handler) downloadTrack(c *gin.Context) {

}
