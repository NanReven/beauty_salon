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

func (serv *AdminService) CreateFavour(input *entity.Favour) (int, error) {
	if input.Price < 0 {
		return 0, errors.New("invalid favour price")
	}
	return serv.adminRepo.CreateFavour(input)
}

func (serv *AdminService) UpdateMasterInfo(input *entity.MasterUpdate) error {
	if input.MasterId < 0 || input.UserId < 0 {
		return errors.New("invalid ids")
	}
	if input.UserId != 0 {
		masterName, err := serv.masterRepo.GetMasterName(input.UserId)
		if err != nil {
			return err
		}
		slugified := slug.Make(masterName)
		if err := serv.masterRepo.UpdateUserId(input.MasterId, input.UserId, slugified); err != nil {
			return err
		}
	}
	if input.PositionId != 0 {
		if err := serv.masterRepo.UpdatePositionId(input.MasterId, input.PositionId); err != nil {
			return err
		}
	}
	if input.Bio != "" {
		if err := serv.masterRepo.UpdateBio(input.MasterId, input.Bio); err != nil {
			return err
		}
	}
	return nil
}
