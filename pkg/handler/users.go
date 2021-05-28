package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wellWINeo/MusicPlayerBackend"
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
	var user MusicPlayerBackend.User
	value, ok := c.Get(userIdCtx)
	if !ok {
		newErrorResponse(c, http.StatusUnauthorized, "user id in token not found")
		return
	}
	if err := c.BindJSON(&user); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	userId := value.(int)
	user.Id = userId

	if err := h.services.UpdateUser(user); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.Status(http.StatusOK)
}

func (h *Handler) deleteUser(c *gin.Context) {
	id, ok := c.Get(userIdCtx)
	if !ok {
		newErrorResponse(c, http.StatusUnauthorized, "user id in token not found")
		return
	}
	h.services.DeleteUser(id.(int))
}
