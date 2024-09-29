package service

import (
	"beauty_salon/internal/adapter/dto"
	"beauty_salon/internal/adapter/repository"
)

type FavourService struct {
	repo repository.Favour
}

func NewFavourService(repo repository.Favour) *FavourService {
	return &FavourService{repo: repo}
}

func (uc *FavourService) GetAllFavours() ([]dto.FavourResponse, error) {
	return uc.repo.GetAllFavours()
}

func (uc *FavourService) GetFavourById(id int) (dto.FavourResponse, error) {
	return uc.repo.GetFavourById(id)
}
