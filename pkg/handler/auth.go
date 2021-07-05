package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/wellWINeo/MusicPlayerBackend"
)

func (h *Handler) signUp(ctx *gin.Context) {
	var input MusicPlayerBackend.User

	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.CreateUser(input)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	logrus.Println(id)
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type SignInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(ctx *gin.Context) {
	var input SignInInput

	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	token, err := h.services.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		logrus.Print("here")
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

type VerifyInput struct {
	UserId int `json:"user_id,string" binding:"required"`
	Code   int `json:"code,string" binding:"required"`
}

func (h *Handler) verify(ctx *gin.Context) {
	input := VerifyInput{}

	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.services.GetUser(input.UserId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	value, ok := h.services.Verify(input.Code)
	if !ok {
		newErrorResponse(ctx, http.StatusInternalServerError, "No user with such code")
		return
	}
	value.Id = user.Id

	if value == user {
		user.IsVerified = true
		err := h.services.UpdateUser(user)
		if err != nil {
			newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
			return
		}
	} else {
		newErrorResponse(ctx, http.StatusInternalServerError, "not equal users structs")
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"msg": "account verified",
	})

}
