package service

import (
	"beauty_salon/internal/adapter/repository"
	"beauty_salon/internal/domain/entity"
	"errors"
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
	if id < 0 {
		return entity.MasterResponse{}, errors.New("invalid master id")
	}
	return uc.repo.GetMasterById(id)
}

func (uc *MasterService) GetMasterName(userId int) (string, error) {
	if userId < 0 {
		return "", errors.New("invalid user id")
	}
	return uc.repo.GetMasterName(userId)
}
