package usecase

import (
	"beauty_salon/internal/adapter/dto"
	"beauty_salon/internal/adapter/repository"
)

type MasterUsecase struct {
	repo repository.Master
}

func NewMasterUsecase(repo repository.Master) *MasterUsecase {
	return &MasterUsecase{repo: repo}
}

func (uc *MasterUsecase) GetAllMasters() ([]dto.MasterResponse, error) {
	return uc.repo.GetAllMasters()
}

func (uc *MasterUsecase) GetMasterById(id int) (dto.MasterResponse, error) {
	return uc.repo.GetMasterById(id)
}
