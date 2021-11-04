package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wellWINeo/MusicPlayerBackend"
	"io"
	"net/http"
	"os"
	"path"
	"regexp"
	"strconv"
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

var MEDIA_TYPES = map[string]interface{}{
	// audio
	"audio/basic":            nil,
	"audio/L24":              nil,
	"audio/mp4":              nil,
	"audio/acc":              nil,
	"audio/mpeg":             nil,
	"audio/ogg":              nil,
	"audio/vorbis":           nil,
	"audio/x-ms-wma":         nil,
	"audio/x-ms-wax":         nil,
	"audio/vnd.rn-realaudio": nil,
	"audio/vnd.wave":         nil,
	"audio/web,":             nil,
	// video
	"video/mpeg":      nil,
	"video/mp4":       nil,
	"video/ogg":       nil,
	"video/quicktime": nil,
	"video/webm,":     nil,
	"video/x-ms-wmv":  nil,
	"video/x-flv":     nil,
	"video/x-msvideo": nil,
	"video/3gpp":      nil,
	"video/3gpp2":     nil,
	//
	"application/octet-stream": nil,
}

func (h *Handler) uploadTrack(c *gin.Context) {
	trackId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = c.SaveUploadedFile(file, path.Join(h.dataPath, strconv.Itoa(trackId)))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.Status(http.StatusOK)
}

func (h *Handler) loadTrackPart(c *gin.Context) {
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
	re := regexp.MustCompile(`(?m)\d{3}`)

	if !re.MatchString(partId) && partId != "list" {
		newErrorResponse(c, http.StatusBadRequest, "not a number in part")
		return
	}

	if err := h.services.History.AddHistory(trackId, userId); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.File(path.Join(h.dataPath, fmt.Sprint(trackId), partId))
}

// deprecated
func (h *Handler) downloadTrack(c *gin.Context) {
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

	if err := h.services.History.AddHistory(trackId, userId); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	file, err := os.Open(path.Join(h.dataPath, strconv.Itoa(trackId)))
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	defer file.Close()

	io.Copy(c.Writer, file)

	c.Status(http.StatusOK)
}
