package service

import (
	"beauty_salon/internal/adapter/repository"
	"beauty_salon/internal/domain/entity"
	"errors"

	"github.com/gosimple/slug"
)

type AdminService struct {
	adminRepo  repository.Admin
	masterRepo repository.Master
}

func NewAdminService(adminRepo repository.Admin, masterRepo repository.Master) *AdminService {
	return &AdminService{adminRepo: adminRepo, masterRepo: masterRepo}
}

func (serv *AdminService) CreateMaster(input *entity.Master) (int, error) {
	if input.UserId < 0 || input.PositionId < 0 {
		return 0, errors.New("wrong ids")
	}
	masterName, err := serv.masterRepo.GetMasterName(input.UserId)
	if err != nil {
		return 0, err
	}
	slugified := slug.Make(masterName)
	return serv.adminRepo.CreateMaster(input, slugified)
}
