package service

import (
	"beauty_salon/internal/adapter/dto"
	"beauty_salon/internal/adapter/repository"
)

type MasterService struct {
	repo repository.Master
}

func NewMasterService(repo repository.Master) *MasterService {
	return &MasterService{repo: repo}
}

func (uc *MasterService) GetAllMasters() ([]dto.MasterResponse, error) {
	return uc.repo.GetAllMasters()
}

func (uc *MasterService) GetMasterById(id int) (dto.MasterResponse, error) {
	return uc.repo.GetMasterById(id)
}
