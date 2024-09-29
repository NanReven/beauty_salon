package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *Handler) AuthMiddleware(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		newErrorResponse(c, http.StatusUnauthorized, "user is not authorized")
		return
	}
	tokenData := strings.Split(token, " ")
	if len(tokenData) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}
	userId, isMaster, err := h.service.ParseToken(tokenData[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	c.Set("userId", userId)
	c.Set("isMaster", isMaster)
}

func (h *Handler) GetUserId(c *gin.Context) (int, error) {
	id, ok := c.Get("userId")
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return 0, errors.New("user id not found")
	}
	userId, ok := id.(int)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id is of invalid type")
		return 0, errors.New("user id is of invalid type")
	}
	return userId, nil
}
