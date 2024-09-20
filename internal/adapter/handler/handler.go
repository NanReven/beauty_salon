package handler

import (
	"beauty_salon/internal/usecase"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	usecase *usecase.Usecase
}

func NewHandler(usecase *usecase.Usecase) *Handler {
	return &Handler{usecase: usecase}
}

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/register", h.Register)
		auth.POST("/login", h.Login)
	}

	api := router.Group("/api")
	{
		masters := api.Group("/masters")
		{
			masters.GET("/get_masters", h.GetAllMasters)
			masters.GET("/:id", h.GetMasterById)
		}

		services := api.Group("/services")
		{
			services.POST("/add_service", h.CreateService)
			services.GET("/get_services", h.GetAllServices)
			services.GET("/:id", h.GetServiceById)
		}

		appointments := api.Group("/appointments")
		{
			appointments.POST("/set_appointment", h.SetAppointment)
			appointments.DELETE("/cancel_appointment", h.CancelAppointment)
			appointments.GET("/get_appointments", h.GetAllAppointments)
			appointments.GET("/:id", h.GetAppointmentById)
		}

		admin := api.Group("/admin")
		{
			admin.POST("/masters", h.CreateMaster)
			admin.PUT("/masters/:id", h.UpdateMasterInfo)
			admin.DELETE("masters/:id", h.DeleteMasterAccount)
			admin.POST("/services", h.CreateService)
			admin.PUT("/services/:id", h.UpdateService)
			admin.DELETE("/services/:id", h.RemoveService)
		}
	}

	return router
}
