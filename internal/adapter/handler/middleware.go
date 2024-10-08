package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	ErrUserIDNotFound    = errors.New("user id not found")
	ErrUserIDInvalidType = errors.New("user id is invalid")
)

func (h *Handler) AuthMiddleware(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		newErrorResponse(c, http.StatusUnauthorized, "user is unauthorized")
		return
	}

	tokenData := strings.Split(token, " ")
	if len(tokenData) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	userId, role, err := h.service.ParseToken(tokenData[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set("userId", userId)
	c.Set("role", role)
}

func (h *Handler) GetUserId(c *gin.Context) (int, bool) {
	id, ok := c.Get("userId")
	if !ok {
		newErrorResponse(c, http.StatusUnauthorized, "user is unauthorized")
		return 0, false
	}

	userId, ok := id.(int)
	if !ok {
		newErrorResponse(c, http.StatusBadRequest, "invalid user id")
		return 0, false
	}

	return userId, true
}

func (h *Handler) GetUserRole(c *gin.Context) (string, error) {
	role, ok := c.Get("role")
	if !ok {
		newErrorResponse(c, http.StatusUnauthorized, "user role is not set")
		return "", errors.New("user role is not set")
	}
	userRole, ok := role.(string)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user role is invalid")
		return "", errors.New("user role is invalid")
	}
	return userRole, nil
}

func (h *Handler) CheckAdminRole(c *gin.Context) {
	role, err := h.GetUserRole(c)
	if err != nil || role != "admin" {
		newErrorResponse(c, http.StatusForbidden, "only for admins")
		return
	}
}
