package service

import (
	"beauty_salon/internal/adapter/repository"
	"beauty_salon/internal/domain/entity"
)

type Appointment interface {
	CreateAppointment(userId int, appointment *entity.AppointmentInput) (int, error) // POST
	CancelAppointment(userId, appointmentId int) (string, error)                     // DELETE
	GetAllAppointments(userId int) ([]entity.AppointmentResponse, error)
	GetAppointmentById(userId, appointmentId int) (entity.AppointmentResponse, error)
}

type Master interface {
	//CreateMaster()                                // POST
	GetAllMasters() ([]entity.MasterResponse, error)     // GET
	GetMasterById(id int) (entity.MasterResponse, error) // GET
	//DeleteMasterAccount(id int)                   // DELETE
	//UpdateMasterInfo(id int)                      // PUT
}

type Favour interface {
	//CreateService()        // POST
	//RemoveService(id int)  // DELETE
	GetAllFavours() ([]entity.FavourResponse, error)     // GET
	GetFavourById(id int) (entity.FavourResponse, error) // GET
	//UpdateService(id int)
}

type User interface {
	Register(input *entity.User) (int, error)             // POST
	GenerateToken(email, password string) (string, error) // POST
	ParseToken(token string) (int, string, error)
}

type Service struct {
	Appointment
	Master
	Favour
	User
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		User:        NewUserService(repo.User),
		Master:      NewMasterService(repo.Master),
		Favour:      NewFavourService(repo.Favour),
		Appointment: NewAppointmentService(repo.Appointment, repo.Favour),
	}
}
