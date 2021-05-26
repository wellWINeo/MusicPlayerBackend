package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wellWINeo/MusicPlayerBackend"
)

func (h *Handler) signUp(ctx *gin.Context) {
	var input MusicPlayerBackend.User

	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}
}

func (h *Handler) signIn(ctx *gin.Context) {

}
