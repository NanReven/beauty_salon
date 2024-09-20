package handler

import (
	"log"

	"github.com/gin-gonic/gin"
)

func newErrorResponse(c *gin.Context, statucCode int, message string) {
	log.Println(message)
	c.AbortWithStatusJSON(statucCode, gin.H{"error_message": message})
}
