package usecase

import (
	"beauty_salon/internal/adapter/dto"
	"beauty_salon/internal/adapter/repository"
)

type AppointmentUsecase struct {
	repo repository.Appointment
}

func NewAppointmentUsecase(repo repository.Appointment) *AppointmentUsecase {
	return &AppointmentUsecase{repo: repo}
}

func (uc *AppointmentUsecase) CreateAppointment(userId int, appointment *dto.AppointmentInput) (int, error) {
	return uc.repo.CreateAppointment(userId, appointment)
}

func (uc *AppointmentUsecase) GetAllAppointments(userId int) ([]dto.AppointmentResponse, error) {
	return uc.repo.GetAllAppointments(userId)
}

func (uc *AppointmentUsecase) GetAppointmentById(userId, appointmentId int) (dto.AppointmentResponse, error) {
	return uc.repo.GetAppointmentById(userId, appointmentId)
}

func (uc *AppointmentUsecase) CancelAppointment(userId, appointmentId int) (string, error) {
	return uc.repo.CancelAppointment(userId, appointmentId)
}
