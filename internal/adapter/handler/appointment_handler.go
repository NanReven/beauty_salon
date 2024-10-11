package handler

import (
	"beauty_salon/internal/domain/entity"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) SetAppointment(c *gin.Context) {
	userId, ok := h.GetUserId(c)
	if !ok {
		return
	}

	var appointment entity.AppointmentInput
	if err := c.BindJSON(&appointment); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	appointmentId, err := h.service.Appointment.CreateAppointment(userId, &appointment)
	if err != nil {
		if errors.Is(err, entity.ErrInvalidAppointmentInput) {
			newErrorResponse(c, http.StatusBadRequest, "invalid request data")
		} else if errors.Is(err, entity.ErrFavourNotFound) {
			newErrorResponse(c, http.StatusNotFound, err.Error())
		} else if errors.Is(err, entity.ErrMasterIsUnavailable) {
			newErrorResponse(c, http.StatusConflict, "master has accepted appointment")
		} else {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
		return
	}
	c.JSON(http.StatusCreated, appointmentId)
}

func (h *Handler) GetAllAppointments(c *gin.Context) {
	userId, ok := h.GetUserId(c)
	if !ok {
		return
	}

	appointments, err := h.service.Appointment.GetAllAppointments(userId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, appointments)
}

func (h *Handler) GetAppointmentById(c *gin.Context) {
	userId, ok := h.GetUserId(c)
	if !ok {
		return
	}

	appointmentId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	appointment, err := h.service.Appointment.GetAppointmentById(userId, appointmentId)
	if err != nil {
		if errors.Is(err, entity.ErrAppointmentNotFound) {
			newErrorResponse(c, http.StatusNotFound, err.Error())
		} else if errors.Is(err, entity.ErrInvalidAppointmentInput) {
			newErrorResponse(c, http.StatusBadRequest, "invalid appointment id")
		} else {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
		return
	}
	c.JSON(http.StatusOK, appointment)
}

func (h *Handler) CancelAppointment(c *gin.Context) {
	userId, ok := h.GetUserId(c)
	if !ok {
		return
	}

	appointmentId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	status, err := h.service.Appointment.CancelAppointment(userId, appointmentId)
	if err != nil {
		if errors.Is(err, entity.ErrAppointmentNotFound) {
			newErrorResponse(c, http.StatusNotFound, err.Error())
		} else if errors.Is(err, entity.ErrAppointmentCancelled) {
			newErrorResponse(c, http.StatusConflict, err.Error())
		} else if errors.Is(err, entity.ErrInvalidAppointmentInput) {
			newErrorResponse(c, http.StatusBadRequest, "invalid appointment id")
		} else {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": status,
	})
}
