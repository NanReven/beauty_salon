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

	id, err := h.service.Admin.CreateMaster(&input)
	if err != nil {
		if errors.Is(err, entity.ErrInvalidMasterInput) {
			newErrorResponse(c, http.StatusBadRequest, "invalid user id or position id")
		} else {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
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

	if err := h.service.UpdateMasterInfo(&input, masterId); err != nil {
		if errors.Is(err, entity.ErrMasterNotFound) {
			newErrorResponse(c, http.StatusNotFound, err.Error())
		} else if errors.Is(err, entity.ErrInvalidMasterInput) {
			newErrorResponse(c, http.StatusBadRequest, "invalid user id or position id")
		} else {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
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
		if errors.Is(err, entity.ErrInvalidFavourInput) {
			newErrorResponse(c, http.StatusBadRequest, "invalid price or duration")
		} else {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
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

	if err := h.service.UpdateFavourInfo(&input, favourId); err != nil {
		if errors.Is(err, entity.ErrFavourNotFound) {
			newErrorResponse(c, http.StatusNotFound, err.Error())
		} else if errors.Is(err, entity.ErrInvalidFavourInput) {
			newErrorResponse(c, http.StatusBadRequest, "invalid price or duration")
		} else {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
		}
		return
	}
	c.JSON(http.StatusOK, "updated")
}
