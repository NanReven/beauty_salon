package repository

import (
	"beauty_salon/internal/domain/entity"

	"github.com/jmoiron/sqlx"
)

type AdminRepository struct {
	db *sqlx.DB
}

func NewAdminRepository(db *sqlx.DB) *AdminRepository {
	return &AdminRepository{db: db}
}

func (repo *AdminRepository) CreateMaster(input *entity.Master, slug string) (int, error) {
	var id int
	query := "INSERT INTO masters (user_id, position_id, bio, slug) VALUES ($1, $2, $3, $4) RETURNING id"
	row := repo.db.QueryRow(query, input.UserId, input.PositionId, input.Bio, slug)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (repo *AdminRepository) CreateFavour(input *entity.Favour) (int, error) {
	var id int
	query := "INSERT INTO services (category_id, title, duration, price) VALUES ($1, $2, $3, $4) RETURNING id"
	row := repo.db.QueryRow(query, input.CategoryId, input.Title, input.Duration, input.Price)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
