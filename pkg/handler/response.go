package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type error struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, StatusCode int, msg string) {
	logrus.Error(msg)
	c.AbortWithStatusJSON(StatusCode, error{msg})
}
