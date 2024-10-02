package handler

import (
	"beauty_salon/internal/domain/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateMaster(c *gin.Context) {
	var input entity.Master
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.service.CreateMaster(&input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, id)
}

func (h *Handler) DeleteMasterAccount(c *gin.Context) {

}
func (h *Handler) UpdateMasterInfo(c *gin.Context) {

}

func (h *Handler) CreateFavour(c *gin.Context) {

}

func (h *Handler) RemoveFavour(c *gin.Context) {

}

func (h *Handler) UpdateFavour(c *gin.Context) {

}
