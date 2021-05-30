package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wellWINeo/MusicPlayerBackend"
)

func (h *Handler) getAllTracks(c *gin.Context) {

}

func (h *Handler) getTrack(c *gin.Context) {

}

func (h *Handler) updateTrack(c *gin.Context) {

}

func (h *Handler) createTrack(c *gin.Context) {
	var input MusicPlayerBackend.Track
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	trackId, err := h.services.Tracks.CreateTrack(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"trackId": trackId,
	})
}

func (h *Handler) deleteTrack(c *gin.Context) {
	trackId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "can't parse param")
		return
	}

	if err := h.services.Tracks.DeleteTrack(trackId); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) uploadTrack(c *gin.Context) {

}

func (h *Handler) downloadTrack(c *gin.Context) {

}
