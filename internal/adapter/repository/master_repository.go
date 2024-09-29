package repository

import (
	"beauty_salon/internal/domain/entity"
	"log"

	"github.com/jmoiron/sqlx"
)

type MasterRepository struct {
	db *sqlx.DB
}

func NewMasterRepository(db *sqlx.DB) *MasterRepository {
	return &MasterRepository{db: db}
}

func (repo *MasterRepository) GetAllMasters() ([]entity.MasterResponse, error) {
	var masters []entity.MasterResponse
	query := "SELECT first_name, second_name, email, title, bio FROM positions JOIN masters ON (positions.id=position_id) JOIN users ON (user_id=users.id)"
	if err := repo.db.Select(&masters, query); err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return masters, nil
}

func (repo *MasterRepository) GetMasterById(id int) (entity.MasterResponse, error) {
	var master entity.MasterResponse
	query := "SELECT first_name, second_name, email, title, bio FROM positions JOIN masters ON (positions.id=position_id) JOIN users ON (user_id=users.id) WHERE masters.id=$1"
	if err := repo.db.Get(&master, query, id); err != nil {
		log.Println(err.Error())
		return master, err
	}
	return master, nil
}
