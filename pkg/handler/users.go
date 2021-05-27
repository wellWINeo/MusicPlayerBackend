package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getUser(c *gin.Context) {
	id, _ := c.Get(userIdCtx)
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) updateUser(c *gin.Context) {

}

func (h *Handler) createUser(c *gin.Context) {

}

func (h *Handler) deleteUser(c *gin.Context) {

}
