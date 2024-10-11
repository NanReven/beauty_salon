package service

import (
	"beauty_salon/internal/adapter/repository"
	"beauty_salon/internal/domain/entity"
	"log"
	"time"
)

type AppointmentService struct {
	repo       repository.Appointment
	favourRepo repository.Favour
}

func NewAppointmentService(repo repository.Appointment, favourRepo repository.Favour) *AppointmentService {
	return &AppointmentService{repo: repo, favourRepo: favourRepo}
}

func (uc *AppointmentService) CreateAppointment(userId int, appointment *entity.AppointmentInput) (int, error) {
	if appointment.MasterId <= 0 || len(appointment.Services) == 0 || appointment.AppointmentStart.IsZero() {
		return 0, entity.ErrInvalidAppointmentInput
	}

	appointmentEnd := appointment.AppointmentStart.Time
	var totalSum float64

	for _, id := range appointment.Services {
		favour, err := uc.favourRepo.GetFavourById(id)
		if err != nil {
			return 0, err
		}
		duration := time.Duration(favour.Duration.Hour()*3600+favour.Duration.Minute()*60) * time.Second
		appointmentEnd = appointmentEnd.Add(duration)
		totalSum += favour.Price
	}

	if !uc.CheckMasterAvailability(appointment.AppointmentStart.Time, appointmentEnd, appointment.MasterId) {
		return 0, entity.ErrMasterIsUnavailable
	}

	return uc.repo.CreateAppointment(userId, appointment, appointmentEnd, totalSum)
}

func (uc *AppointmentService) GetAllAppointments(userId int) ([]entity.AppointmentResponse, error) {
	appointments, err := uc.repo.GetAllAppointments(userId)
	if err != nil {
		return appointments, err
	}

	for i := 0; i < len(appointments); i++ {
		services, err := uc.repo.GetFavoursByAppointmentId(appointments[i].Id)
		if err != nil {
			return appointments, err
		}
		appointments[i].Services = services
	}
	return appointments, nil
}

func (uc *AppointmentService) GetAppointmentById(userId, appointmentId int) (entity.AppointmentResponse, error) {
	if appointmentId <= 0 {
		return entity.AppointmentResponse{}, entity.ErrInvalidAppointmentInput
	}

	appointment, err := uc.repo.GetAppointmentById(userId, appointmentId)
	if err != nil {
		return appointment, err
	}

	services, err := uc.repo.GetFavoursByAppointmentId(appointmentId)
	if err != nil {
		return appointment, err
	}

	appointment.Services = services
	return appointment, nil
}

func (uc *AppointmentService) CancelAppointment(userId, appointmentId int) (string, error) {
	if appointmentId <= 0 {
		return "", entity.ErrInvalidAppointmentInput
	}

	appointment, err := uc.repo.GetAppointmentById(userId, appointmentId)
	if err != nil {
		return "", err
	} else if appointment.Status == "cancelled" {
		return "", entity.ErrAppointmentCancelled
	}

	return uc.repo.CancelAppointment(userId, appointmentId)
}

func (uc *AppointmentService) CheckMasterAvailability(appointmentStart time.Time, appointmentEnd time.Time, masterId int) bool {
	acceptedAppointments, err := uc.repo.GetAcceptedAppointments(appointmentStart, masterId)
	if err != nil {
		return false
	}

	log.Println("Input appointment:", appointmentStart, ":", appointmentEnd)
	for _, appointment := range acceptedAppointments {
		log.Println("Accepted appointment:", appointment.AppointmentStart, ":", appointment.AppointmentEnd)
		if appointmentStart.Before(appointment.AppointmentEnd.Time) && appointmentEnd.After(appointment.AppointmentStart.Time) {
			log.Println("Conflict found with appointment:", appointment.AppointmentStart, "-", appointment.AppointmentEnd)
			return false
		}
	}
	return true
}
