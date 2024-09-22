package usecase

import (
	"beauty_salon/internal/adapter/dto"
	"beauty_salon/internal/adapter/repository"
)

type ServiceUsecase struct {
	repo repository.Service
}

func NewServiceUsecase(repo repository.Service) *ServiceUsecase {
	return &ServiceUsecase{repo: repo}
}

func (uc *ServiceUsecase) GetAllServices() ([]dto.ServiceResponse, error) {
	return uc.repo.GetAllServices()
}

func (uc *ServiceUsecase) GetServiceById(id int) (dto.ServiceResponse, error) {
	return uc.repo.GetServiceById(id)
}
