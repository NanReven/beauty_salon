package repository

import (
	"beauty_salon/internal/domain"
	"beauty_salon/internal/domain/entity"
	"fmt"

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
		return nil, fmt.Errorf("failed to get list of all favours: %w", err)
	}
	return favours, nil
}

func (repo *FavourRepository) GetFavourById(id int) (entity.FavourResponse, error) {
	var favour entity.FavourResponse
	query := "SELECT categories.title AS category_title, services.title AS service_title, duration, price FROM services JOIN categories ON (category_id = categories.id) WHERE services.id=$1"
	if err := repo.db.Get(&favour, query, id); err != nil {
		return favour, fmt.Errorf("failed to get favour with id %d: %w", id, err)
	}
	return favour, nil
}

func (repo *FavourRepository) UpdateCategoryId(favourId, categoryId int) error {
	query := "UPDATE services SET category_id=$1 WHERE id=$2"
	if _, err := repo.db.Exec(query, categoryId, favourId); err != nil {
		return fmt.Errorf("failed to update category of favour with id %d: %w", favourId, err)
	}
	return nil
}

func (repo *FavourRepository) UpdateFavourTitle(favourId int, title string) error {
	query := "UPDATE services SET title=$1 WHERE id=$2"
	if _, err := repo.db.Exec(query, title, favourId); err != nil {
		return fmt.Errorf("failed to update title of favour with id %d: %w", favourId, err)
	}
	return nil
}

func (repo *FavourRepository) UpdateFavourDuration(favourId int, duration domain.CustomDuration) error {
	query := "UPDATE services SET duration=$1 WHERE id=$2"
	if _, err := repo.db.Exec(query, duration, favourId); err != nil {
		return fmt.Errorf("failed to update duration of favour with id %d: %w", favourId, err)
	}
	return nil
}

func (repo *FavourRepository) UpdateFavourPrice(favourId int, price float64) error {
	query := "UPDATE services SET price=$1 WHERE id=$2"
	if _, err := repo.db.Exec(query, price, favourId); err != nil {
		return fmt.Errorf("failed to update price of favour with id %d: %w", favourId, err)
	}
	return nil
}
