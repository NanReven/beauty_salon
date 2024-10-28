package service

import (
	"beauty_salon/internal/adapter/repository"
	"beauty_salon/internal/domain/entity"
)

type Appointment interface {
	CreateAppointment(userId int, appointment *entity.AppointmentInput) (int, error)
	CancelAppointment(userId, appointmentId int) (string, error)
	GetAllAppointments(userId int) ([]entity.AppointmentResponse, error)
	GetAppointmentById(userId, appointmentId int) (entity.AppointmentResponse, error)
}

type Master interface {
	GetAllMasters() ([]entity.MasterResponse, error)
	GetMasterById(id int) (entity.MasterResponse, error)
	ReplyToAppointment(input *entity.AppointmentReply, masterId int) error
}

type Favour interface {
	GetAllFavours() ([]entity.FavourResponse, error)
	GetFavourById(id int) (entity.FavourResponse, error)
}

type User interface {
	Register(input *entity.User) (int, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(token string) (int, string, error)
}

type Admin interface {
	CreateMaster(input *entity.Master) (int, error)
	UpdateMasterInfo(input *entity.MasterUpdate, masterId int) error
	CreateFavour(input *entity.Favour) (int, error)
	UpdateFavourInfo(input *entity.FavourUpdate, favourId int) error
}

type Service struct {
	Appointment
	Master
	Favour
	User
	Admin
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		User:        NewUserService(repo.User),
		Master:      NewMasterService(repo.Master),
		Favour:      NewFavourService(repo.Favour),
		Appointment: NewAppointmentService(repo.Appointment, repo.Favour, repo.Master),
		Admin:       NewAdminService(repo.Admin, repo.Master, repo.Favour),
	}
}
