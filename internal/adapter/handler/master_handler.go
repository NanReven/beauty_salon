package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateMaster(c *gin.Context) {

}

func (h *Handler) GetAllMasters(c *gin.Context) {
	masters, err := h.usecase.Master.GetAllMasters()
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
	master, err := h.usecase.Master.GetMasterById(id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, master)

}
func (h *Handler) DeleteMasterAccount(c *gin.Context) {

}
func (h *Handler) UpdateMasterInfo(c *gin.Context) {

}
