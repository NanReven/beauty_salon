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
	favourRepo repository.Favour
}

func NewAdminService(adminRepo repository.Admin, masterRepo repository.Master, favourRepo repository.Favour) *AdminService {
	return &AdminService{adminRepo: adminRepo, masterRepo: masterRepo, favourRepo: favourRepo}
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

func (serv *AdminService) UpdateMasterInfo(input *entity.MasterUpdate, masterId int) error {
	if masterId < 0 || input.UserId < 0 {
		return errors.New("invalid ids")
	}
	if input.UserId != 0 {
		masterName, err := serv.masterRepo.GetMasterName(input.UserId)
		if err != nil {
			return err
		}
		slugified := slug.Make(masterName)
		if err := serv.masterRepo.UpdateUserId(masterId, input.UserId, slugified); err != nil {
			return err
		}
	}
	if input.PositionId != 0 {
		if err := serv.masterRepo.UpdatePositionId(masterId, input.PositionId); err != nil {
			return err
		}
	}
	if input.Bio != "" {
		if err := serv.masterRepo.UpdateBio(masterId, input.Bio); err != nil {
			return err
		}
	}
	return nil
}

func (serv *AdminService) UpdateFavourInfo(input *entity.FavourUpdate, favourId int) error {
	if input.CategoryId < 0 {
		return errors.New("invalid category id")
	} else if input.Price < 0 {
		return errors.New("invalid price")
	}

	if input.CategoryId != 0 {
		if err := serv.favourRepo.UpdateCategoryId(favourId, input.CategoryId); err != nil {
			return err
		}
	}

	if input.Title != "" {
		if err := serv.favourRepo.UpdateFavourTitle(favourId, input.Title); err != nil {
			return err
		}
	}

	if !input.Duration.IsZero() {
		if err := serv.favourRepo.UpdateFavourDuration(favourId, input.Duration); err != nil {
			return err
		}
	}

	if input.Price != 0 {
		if err := serv.favourRepo.UpdateFavourPrice(favourId, input.Price); err != nil {
			return err
		}
	}

	return nil
}
