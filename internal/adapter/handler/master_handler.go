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
	}

	master, err := h.service.Master.GetMasterById(id)
	if err != nil {
		if errors.Is(err, entity.ErrMasterNotFound) {
			newErrorResponse(c, http.StatusNotFound, err.Error())
		} else if errors.Is(err, entity.ErrInvalidMasterInput) {
			newErrorResponse(c, http.StatusBadRequest, "invalid master id")
		} else {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
		return
	}
	c.JSON(http.StatusOK, master)
}

func (h *Handler) ReplyToAppointment(c *gin.Context) {
	var input entity.AppointmentReply
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, flag := h.GetUserId(c)
	if flag != true {
		return
	}

	if err := h.service.Master.ReplyToAppointment(&input, id); err != nil {
		if errors.Is(err, entity.ErrInvalidAppointmentInput) {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
		} else if errors.Is(err, entity.ErrAppointmentNotFound) {
			newErrorResponse(c, http.StatusNotFound, err.Error())
		}
		return
	}
	c.JSON(http.StatusOK, "updated")
}
