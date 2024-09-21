package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateService(c *gin.Context) {

}

func (h *Handler) RemoveService(c *gin.Context) {

}

func (h *Handler) GetAllServices(c *gin.Context) {
	services, err := h.usecase.Service.GetAllServices()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, services)
}

func (h *Handler) GetServiceById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	service, err := h.usecase.Service.GetServiceById(id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, service)
}

func (h *Handler) UpdateService(c *gin.Context) {

}
