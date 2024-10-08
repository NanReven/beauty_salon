package handler

import (
	"beauty_salon/internal/domain/entity"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateMaster(c *gin.Context) {
	var input entity.Master
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if input.UserId < 0 {
		newErrorResponse(c, http.StatusBadRequest, "user id must be greater than zero")
		return
	} else if input.PositionId < 0 {
		newErrorResponse(c, http.StatusBadRequest, "position id must be greater than zero")
		return
	}

	id, err := h.service.Admin.CreateMaster(&input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, id)
}

func (h *Handler) UpdateMasterInfo(c *gin.Context) {
	masterId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid master id")
		return
	}

	var input entity.MasterUpdate
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if input.UserId < 0 {
		newErrorResponse(c, http.StatusBadRequest, "user id must be greater than zero")
		return
	} else if input.PositionId < 0 {
		newErrorResponse(c, http.StatusBadRequest, "position id must be greater than zero")
		return
	}

	if err := h.service.UpdateMasterInfo(&input, masterId); err != nil {
		if errors.Is(err, entity.ErrMasterNotFound) {
			newErrorResponse(c, http.StatusNotFound, err.Error())
		}
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "updated")
}

func (h *Handler) CreateFavour(c *gin.Context) {
	var input entity.Favour
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.service.Admin.CreateFavour(&input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, id)
}

func (h *Handler) UpdateFavour(c *gin.Context) {
	favourId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid favour id")
		return
	}

	var input entity.FavourUpdate
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if input.CategoryId < 0 {
		newErrorResponse(c, http.StatusBadRequest, "category id must be greater than zero")
		return
	} else if input.Price < 0 {
		newErrorResponse(c, http.StatusBadRequest, "favour price must be greater than zero")
		return
	}

	if err := h.service.UpdateFavourInfo(&input, favourId); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, "updated")
}
