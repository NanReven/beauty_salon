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
	if id <= 0 {
		return entity.MasterResponse{}, entity.ErrInvalidMasterInput
	}
	return uc.repo.GetMasterById(id)
}

func (uc *MasterService) GetMasterName(userId int) (string, error) {
	if userId <= 0 {
		return "", entity.ErrInvalidMasterInput
	}
	return uc.repo.GetMasterName(userId)
}

func (uc *MasterService) ReplyToAppointment(input *entity.AppointmentReply, masterId int) error {
	if input.AppointmentId <= 0 {
		return entity.ErrInvalidAppointmentInput
	} else if err := uc.repo.GetMasterAppointment(masterId, input.AppointmentId); err != nil {
		return err
	}
	return uc.repo.ReplyToAppointment(input)
}
