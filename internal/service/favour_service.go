package service

import (
	"beauty_salon/internal/adapter/repository"
	"beauty_salon/internal/domain/entity"
)

type FavourService struct {
	repo repository.Favour
}

func NewFavourService(repo repository.Favour) *FavourService {
	return &FavourService{repo: repo}
}

func (uc *FavourService) GetAllFavours() ([]entity.FavourResponse, error) {
	return uc.repo.GetAllFavours()
}

func (uc *FavourService) GetFavourById(id int) (entity.FavourResponse, error) {
	return uc.repo.GetFavourById(id)
}
