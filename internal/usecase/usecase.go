package usecase

import (
	"beauty_salon/internal/adapter/repository"
	"beauty_salon/internal/domain/entity"
)

type Appointment interface {
	CreateAppointment() // POST
	CancelAppointment() // DELETE
	GetAllAppointments()
	GetAppointmentById()
}

type Master interface {
	CreateMaster()              // POST
	GetAllMasters()             // GET
	GetMasterById(id int)       // GET
	DeleteMasterAccount(id int) // DELETE
	UpdateMasterInfo(id int)    // PUT
}

type Service interface {
	CreateService()        // POST
	RemoveService(id int)  // DELETE
	GetAllServices()       // GET
	GetServiceById(id int) // GET
	UpdateService(id int)
}

type User interface {
	Register(input *entity.User) (int, error)             // POST
	GenerateToken(email, password string) (string, error) // POST
}

type Usecase struct {
	Appointment
	Master
	Service
	User
}

func NewUsecase(repo *repository.Repository) *Usecase {
	return &Usecase{
		User: NewUserUsecase(repo.User),
	}
}
