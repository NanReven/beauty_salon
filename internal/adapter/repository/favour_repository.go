package repository

import (
	"beauty_salon/internal/adapter/dto"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type FavourRepository struct {
	db *sqlx.DB
}

func NewFavourRepository(db *sqlx.DB) *FavourRepository {
	return &FavourRepository{db: db}
}

func (repo *FavourRepository) GetAllFavours() ([]dto.FavourResponse, error) {
	var favours []dto.FavourResponse
	query := fmt.Sprintf("SELECT categories.title AS category_title, services.title AS service_title, duration, price FROM %s JOIN %s ON (category_id = categories.id)", servicesTable, categoriesTable)
	if err := repo.db.Select(&favours, query); err != nil {
		return nil, err
	}
	return favours, nil
}

func (repo *FavourRepository) GetFavourById(id int) (dto.FavourResponse, error) {
	var favour dto.FavourResponse
	query := fmt.Sprintf("SELECT categories.title AS category_title, services.title AS service_title, duration, price FROM %s JOIN %s ON (category_id = categories.id) WHERE services.id=$1", servicesTable, categoriesTable)
	if err := repo.db.Get(&favour, query, id); err != nil {
		return favour, err
	}
	return favour, nil
}
