package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authHeader = "Authorization"
	userIdCtx = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	// checking header
	header := c.GetHeader(authHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "Empty authorization header")
		return
	}

	spilittedHeader := strings.Split(header, " ")
	if len(spilittedHeader) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "Invalid auth header form")
		return
	}

	// checking token
	userId, err := h.services.Authorization.ParseToken(spilittedHeader[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userIdCtx, userId)
}
