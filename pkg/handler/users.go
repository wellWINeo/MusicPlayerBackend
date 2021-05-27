package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getUser(c *gin.Context) {
	id, ok := c.Get(userIdCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return
	}

	user, err := h.services.Authorization.GetUser(id.(int))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// hide password
	user.Password = ""
	c.JSON(http.StatusOK, user)
}

func (h *Handler) updateUser(c *gin.Context) {

}

func (h *Handler) createUser(c *gin.Context) {

}

func (h *Handler) deleteUser(c *gin.Context) {
	id, ok := c.Get(userIdCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
	}
	h.services.DeleteUser(id.(int))
}
