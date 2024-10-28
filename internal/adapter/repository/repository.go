package repository

import (
	"beauty_salon/internal/domain"
	"beauty_salon/internal/domain/entity"
	"time"

	"github.com/jmoiron/sqlx"
)

type Appointment interface {
	CreateAppointment(userId int, appointment *entity.AppointmentInput, appointmentEnd time.Time, totalSum float64) (int, error)
	GetAllAppointments(userId int) ([]entity.AppointmentResponse, error)
	GetAppointmentById(userId, appointmentId int) (entity.AppointmentResponse, error)
	GetFavoursByAppointmentId(appointmentId int) ([]entity.FavourResponse, error)
	CancelAppointment(userId, appointmentId int) (string, error)
	GetAcceptedAppointments(appointmentDate time.Time, masterId int) ([]entity.AppointmentTime, error)
}

type Master interface {
	GetAllMasters() ([]entity.MasterResponse, error)
	GetMasterById(id int) (entity.MasterResponse, error)
	GetMasterName(userId int) (string, error)
	UpdateUserId(masterId, userId int, slugified string) error
	UpdatePositionId(masterId, positionId int) error
	UpdateBio(masterId int, bio string) error
	GetMasterEmail(masterId int) (string, error)
	ReplyToAppointment(input *entity.AppointmentReply) error
	GetMasterAppointment(masterId int, appointmentId int) error
}

type Favour interface {
	GetAllFavours() ([]entity.FavourResponse, error)
	GetFavourById(id int) (entity.FavourResponse, error)
	UpdateCategoryId(favourId, categoryId int) error
	UpdateFavourTitle(favourId int, title string) error
	UpdateFavourDuration(favourId int, duration domain.CustomDuration) error
	UpdateFavourPrice(favourId int, price float64) error
}

type User interface {
	CreateUser(input *entity.User) (int, error)
	GetUser(email, password string) (entity.User, error)
}

type Admin interface {
	CreateMaster(input *entity.Master, slug string) (int, error)
	CreateFavour(input *entity.Favour) (int, error)
}

type Repository struct {
	Appointment
	Master
	Favour
	User
	Admin
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User:        NewUserRepository(db),
		Master:      NewMasterRepository(db),
		Favour:      NewFavourRepository(db),
		Appointment: NewAppointmentRepository(db),
		Admin:       NewAdminRepository(db),
	}
}
