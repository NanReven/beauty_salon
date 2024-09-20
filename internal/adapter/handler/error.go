package handler

import (
	"log"

	"github.com/gin-gonic/gin"
)

type error struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statucCode int, message string) {
	log.Println(message)
	c.AbortWithStatusJSON(statucCode, error{Message: message})
}
