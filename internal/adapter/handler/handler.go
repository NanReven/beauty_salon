package handler

import (
	"beauty_salon/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/register", h.Register)
		auth.POST("/login", h.Login)
	}

	api := router.Group("/api", h.AuthMiddleware)
	{
		masters := api.Group("/masters")
		{
			masters.GET("", h.GetAllMasters)
			masters.GET("/:id", h.GetMasterById)
		}

		Favours := api.Group("/favours")
		{
			Favours.GET("", h.GetAllFavours)
			Favours.GET("/:id", h.GetFavourById)
		}

		appointments := api.Group("/appointments")
		{
			appointments.POST("", h.SetAppointment)
			appointments.DELETE("/:id", h.CancelAppointment)
			appointments.GET("", h.GetAllAppointments)
			appointments.GET("/:id", h.GetAppointmentById)
		}

		admin := api.Group("/admin", h.CheckAdminRole)
		{
			admin.POST("/masters", h.CreateMaster)
			admin.PUT("/masters/:id", h.UpdateMasterInfo)
			admin.POST("/favours", h.CreateFavour)
			admin.PUT("/favours/:id", h.UpdateFavour)
		}
	}

	return router
}
