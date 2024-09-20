package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateService(c *gin.Context) {

}

func (h *Handler) RemoveService(c *gin.Context) {

}

func (h *Handler) GetAllServices(c *gin.Context) {
	c.JSON(http.StatusOK, "ok")

}

func (h *Handler) GetServiceById(c *gin.Context) {

}

func (h *Handler) UpdateService(c *gin.Context) {

}
