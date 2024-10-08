package handler

import (
	"beauty_salon/internal/domain/entity"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllMasters(c *gin.Context) {
	masters, err := h.service.Master.GetAllMasters()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "failed to get masters list")
		return
	}
	c.JSON(http.StatusOK, masters)
}

func (h *Handler) GetMasterById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid master id")
		return
	} else if id < 0 {
		newErrorResponse(c, http.StatusBadRequest, "master id must be greater than zero")
		return
	}

	master, err := h.service.Master.GetMasterById(id)
	if err != nil {
		if errors.Is(err, entity.ErrMasterNotFound) {
			newErrorResponse(c, http.StatusNotFound, err.Error())
		} else {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
		return
	}
	c.JSON(http.StatusOK, master)
}
