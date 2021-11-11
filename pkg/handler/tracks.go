package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wellWINeo/MusicPlayerBackend"
	"net/http"
	"path"
	"regexp"
	"strconv"
	"time"
)

func (h *Handler) getAllTracks(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "can't get user id")
		return
	}
	allTracks, err := h.services.GetAllTracks(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, allTracks)
}

func (h *Handler) getTrack(c *gin.Context) {
	trackId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "can't parse param")
		return
	}

	track, err := h.services.Tracks.GetTrack(trackId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, track)

}

func (h *Handler) updateTrack(c *gin.Context) {
	var input MusicPlayerBackend.Track

	trackId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "can't parse param")
		return
	}
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Tracks.UpdateTrack(trackId, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
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

// api endpoint to upload files & assign them to track metadata in DB
func (h *Handler) uploadTrack(c *gin.Context) {

	// get id from URL
	trackId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// parsing file
	file, err := c.FormFile("file")
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// saving file to tmp directory
	fileName := path.Join(h.tempPath, strconv.FormatInt(time.Now().UnixNano(), 10))
	err = c.SaveUploadedFile(file, fileName)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.Tracks.UploadTrack(trackId, fileName, h.dataPath)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// response
	c.Status(http.StatusOK)
}

func (h *Handler) streamTrack(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "can't get user id")
		return
	}

	trackId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "can't parse param")
		return
	}

	partId := c.Param("part")
	// TODO something wrong with regexp
	re := regexp.MustCompile(`(?m)\d{3}`)

	if !re.MatchString(partId) && partId != "index" {
		newErrorResponse(c, http.StatusBadRequest, "invalid part parameter")
		return
	}

	if err := h.services.History.AddHistory(trackId, userId); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.File(path.Join(h.dataPath, fmt.Sprint(trackId), partId))
}
