package repository

import (
	"beauty_salon/internal/domain/entity"

	"github.com/jmoiron/sqlx"
)

type FavourRepository struct {
	db *sqlx.DB
}

func NewFavourRepository(db *sqlx.DB) *FavourRepository {
	return &FavourRepository{db: db}
}

func (repo *FavourRepository) GetAllFavours() ([]entity.FavourResponse, error) {
	var favours []entity.FavourResponse
	query := "SELECT categories.title AS category_title, services.title AS service_title, duration, price FROM services JOIN categories ON (category_id = categories.id)"
	if err := repo.db.Select(&favours, query); err != nil {
		return nil, err
	}
	return favours, nil
}

func (repo *FavourRepository) GetFavourById(id int) (entity.FavourResponse, error) {
	var favour entity.FavourResponse
	query := "SELECT categories.title AS category_title, services.title AS service_title, duration, price FROM services JOIN categories ON (category_id = categories.id) WHERE services.id=$1"
	if err := repo.db.Get(&favour, query, id); err != nil {
		return favour, err
	}
	return favour, nil
}
