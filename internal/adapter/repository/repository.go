package repository

import (
	"beauty_salon/internal/adapter/dto"
	"beauty_salon/internal/domain/entity"

	"github.com/jmoiron/sqlx"
)

const (
	usersTable               = "users"
	mastersTable             = "masters"
	servicesTable            = "services"
	appointmentsTable        = "appointments"
	categoriesTable          = "categories"
	positionsTable           = "positions"
	appointmentServicesTable = "appointmentServices"
)

type Appointment interface {
	CreateAppointment(userId int, appointment *dto.AppointmentInput) (int, error)
	GetAllAppointments(userId int) ([]dto.AppointmentResponse, error)
	GetAppointmentById(userId, appointmentId int) (dto.AppointmentResponse, error)
	CancelAppointment(userId, appointmentId int) (string, error)
}

type Master interface {
	GetAllMasters() ([]dto.MasterResponse, error)
	GetMasterById(id int) (dto.MasterResponse, error)
}

type Service interface {
	GetAllServices() ([]dto.ServiceResponse, error)
	GetServiceById(id int) (dto.ServiceResponse, error)
}

type User interface {
	CreateUser(input *entity.User) (int, error)
	GetUser(email, password string) (entity.User, error)
}

type Repository struct {
	Appointment
	Master
	Service
	User
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User:        NewUserRepository(db),
		Master:      NewMasterRepository(db),
		Service:     NewServiceRepository(db),
		Appointment: NewAppointmentService(db),
	}
}
