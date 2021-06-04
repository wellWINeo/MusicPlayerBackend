package handler

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wellWINeo/MusicPlayerBackend"
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

// func (h *Handler) uploadTrack(c *gin.Context) {
// 	var buf []byte

// 	trackId, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		newErrorResponse(c, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	buf, err = io.ReadAll(c.Request.Body)

// 	if err := h.services.Tracks.UploadTrack(trackId, buf); err != nil {
// 		newErrorResponse(c, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	c.Status(http.StatusOK)
// }

var MEDIA_TYPES = map[string]interface{}{
	// audio
	"audio/basic": nil,
	"audio/L24": nil,
	"audio/mp4": nil,
	"audio/acc": nil,
	"audio/mpeg": nil,
	"audio/ogg": nil,
	"audio/vorbis": nil,
	"audio/x-ms-wma": nil,
	"audio/x-ms-wax": nil,
	"audio/vnd.rn-realaudio": nil,
	"audio/vnd.wave": nil,
	"audio/web,": nil,
	// video
	"video/mpeg": nil,
	"video/mp4": nil,
	"video/ogg": nil,
	"video/quicktime": nil,
	"video/webm,": nil,
	"video/x-ms-wmv": nil,
	"video/x-flv": nil,
	"video/x-msvideo": nil,
	"video/3gpp": nil,
	"video/3gpp2": nil,
	//
	"application/octet-stream": nil,

}

// func (h *Handler) downloadTrack(c *gin.Context) {
// 	userId, err := getUserId(c)
// 	if err != nil {
// 		newErrorResponse(c, http.StatusInternalServerError, "can't get user id")
// 		return
// 	}

// 	trackId, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		newErrorResponse(c, http.StatusBadRequest, "can't parse param")
// 		return
// 	}

// 	if err := h.services.History.AddHistory(trackId, userId); err != nil {
// 		newErrorResponse(c, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	blob, err := h.services.DownloadTrack(trackId)
// 	if err != nil {
// 		newErrorResponse(c, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	logrus.Printf("size: %d", len(blob))

// 	io.Copy(c.Writer, bytes.NewReader(blob))

// 	c.Writer.Header().Add("Content-type", "application/octet-stream")
// 	c.Status(http.StatusOK)
// }

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

	c.SaveUploadedFile(file, fmt.Sprintf("/home/o__ni/server_go/%d", trackId))
	c.Status(http.StatusOK)
}

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

	file, err := os.Open(fmt.Sprintf("/home/o__ni/server_go/%d", trackId))
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	defer file.Close()

	io.Copy(c.Writer, file)

	c.Status(http.StatusOK)
}
