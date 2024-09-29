package handler

import (
	"beauty_salon/internal/adapter/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) SetAppointment(c *gin.Context) {
	userId, err := h.GetUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	var appointment dto.AppointmentInput
	if err := c.BindJSON(&appointment); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	appointmentId, err := h.service.Appointment.CreateAppointment(userId, &appointment)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, appointmentId)
}

func (h *Handler) GetAllAppointments(c *gin.Context) {
	id, err := h.GetUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	appointments, err := h.service.Appointment.GetAllAppointments(id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, appointments)
}

func (h *Handler) GetAppointmentById(c *gin.Context) {
	userId, err := h.GetUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	appointmentId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	appointment, err := h.service.Appointment.GetAppointmentById(userId, appointmentId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, appointment)
}

func (h *Handler) CancelAppointment(c *gin.Context) {
	userId, err := h.GetUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	appointmentId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	status, err := h.service.Appointment.CancelAppointment(userId, appointmentId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": status,
	})
}
