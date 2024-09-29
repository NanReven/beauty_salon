package service

import (
	"beauty_salon/internal/adapter/repository"
	"beauty_salon/internal/domain/entity"
)

type MasterService struct {
	repo repository.Master
}

func NewMasterService(repo repository.Master) *MasterService {
	return &MasterService{repo: repo}
}

func (uc *MasterService) GetAllMasters() ([]entity.MasterResponse, error) {
	return uc.repo.GetAllMasters()
}

func (uc *MasterService) GetMasterById(id int) (entity.MasterResponse, error) {
	return uc.repo.GetMasterById(id)
}
