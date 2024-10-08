package handler

import (
	"beauty_salon/internal/domain/entity"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Register(c *gin.Context) {
	var input entity.User
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.User.Register(&input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "failed to create new user")
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": id,
	})
}

func (h *Handler) Login(c *gin.Context) {
	var input entity.LoginInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.service.User.GenerateToken(input.Email, input.Password)
	if err != nil {
		if errors.Is(err, entity.ErrUserNotFound) {
			newErrorResponse(c, http.StatusNotFound, err.Error())
		} else {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
