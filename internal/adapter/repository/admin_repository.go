package repository

import (
	"beauty_salon/internal/domain/entity"
	"log"

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
		log.Println("wtf no rows create master")
		return 0, err
	}
	return id, nil
}
