package handler

import (
	"beauty_salon/internal/domain/entity"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllFavours(c *gin.Context) {
	favours, err := h.service.Favour.GetAllFavours()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, favours)
}

func (h *Handler) GetFavourById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id")
		return
	}

	favour, err := h.service.Favour.GetFavourById(id)
	if err != nil {
		if errors.Is(err, entity.ErrFavourNotFound) {
			newErrorResponse(c, http.StatusNotFound, err.Error())
		} else {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
		return
	}
	c.JSON(http.StatusOK, favour)
}
