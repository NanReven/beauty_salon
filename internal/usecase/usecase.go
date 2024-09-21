package usecase

import (
	"beauty_salon/internal/adapter/dto"
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
	//CreateMaster()                                // POST
	GetAllMasters() ([]dto.MasterResponse, error)     // GET
	GetMasterById(id int) (dto.MasterResponse, error) // GET
	//DeleteMasterAccount(id int)                   // DELETE
	//UpdateMasterInfo(id int)                      // PUT
}

type Service interface {
	//CreateService()        // POST
	//RemoveService(id int)  // DELETE
	GetAllServices() ([]dto.ServiceResponse, error)     // GET
	GetServiceById(id int) (dto.ServiceResponse, error) // GET
	//UpdateService(id int)
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
		User:    NewUserUsecase(repo.User),
		Master:  NewMasterUsecase(repo.Master),
		Service: newServiceUsecase(repo.Service),
	}
}
