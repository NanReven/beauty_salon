package service

import (
	"beauty_salon/internal/adapter/repository"
	"beauty_salon/internal/domain/entity"
)

type AppointmentService struct {
	repo repository.Appointment
}

func NewAppointmentService(repo repository.Appointment) *AppointmentService {
	return &AppointmentService{repo: repo}
}

func (uc *AppointmentService) CreateAppointment(userId int, appointment *entity.AppointmentInput) (int, error) {
	return uc.repo.CreateAppointment(userId, appointment)
}

func (uc *AppointmentService) GetAllAppointments(userId int) ([]entity.AppointmentResponse, error) {
	return uc.repo.GetAllAppointments(userId)
}

func (uc *AppointmentService) GetAppointmentById(userId, appointmentId int) (entity.AppointmentResponse, error) {
	return uc.repo.GetAppointmentById(userId, appointmentId)
}

func (uc *AppointmentService) CancelAppointment(userId, appointmentId int) (string, error) {
	return uc.repo.CancelAppointment(userId, appointmentId)
}
