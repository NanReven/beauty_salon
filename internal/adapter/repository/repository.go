package repository

import (
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
	CreateAppointment(userId int, appointment *entity.AppointmentInput) (int, error)
	GetAllAppointments(userId int) ([]entity.AppointmentResponse, error)
	GetAppointmentById(userId, appointmentId int) (entity.AppointmentResponse, error)
	CancelAppointment(userId, appointmentId int) (string, error)
}

type Master interface {
	GetAllMasters() ([]entity.MasterResponse, error)
	GetMasterById(id int) (entity.MasterResponse, error)
}

type Favour interface {
	GetAllFavours() ([]entity.FavourResponse, error)
	GetFavourById(id int) (entity.FavourResponse, error)
}

type User interface {
	CreateUser(input *entity.User) (int, error)
	GetUser(email, password string) (entity.User, error)
}

type Repository struct {
	Appointment
	Master
	Favour
	User
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User:        NewUserRepository(db),
		Master:      NewMasterRepository(db),
		Favour:      NewFavourRepository(db),
		Appointment: NewAppointmentRepository(db),
	}
}
