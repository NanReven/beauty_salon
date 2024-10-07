package repository

import (
	"beauty_salon/internal/domain/entity"
	"fmt"

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
		return nil, fmt.Errorf("failed to get list of all masters: %w", err)
	}
	return masters, nil
}

func (repo *MasterRepository) GetMasterById(id int) (entity.MasterResponse, error) {
	var master entity.MasterResponse
	query := "SELECT first_name, second_name, email, title, bio FROM positions JOIN masters ON (positions.id=position_id) JOIN users ON (user_id=users.id) WHERE masters.id=$1"
	if err := repo.db.Get(&master, query, id); err != nil {
		return master, fmt.Errorf("failed to get master wih id %d: %w", id, err)
	}
	return master, nil
}

func (repo *MasterRepository) GetMasterName(userId int) (string, error) {
	var masterName string
	query := "SELECT CONCAT(first_name, ' ', second_name) FROM users WHERE id=$1"
	if err := repo.db.Get(&masterName, query, userId); err != nil {
		return "", fmt.Errorf("failed to get master name wih user id %d: %w", userId, err)
	}
	return masterName, nil
}

func (repo *MasterRepository) UpdateUserId(masterId, userId int, slugified string) error {
	query := "UPDATE masters SET user_id=$1, slug=$2 WHERE id=$3"
	if _, err := repo.db.Exec(query, userId, slugified, masterId); err != nil {
		return fmt.Errorf("failed to update user id %d for mastter with id %d: %w", userId, masterId, err)
	}
	return nil
}

func (repo *MasterRepository) UpdatePositionId(masterId, positionId int) error {
	query := "UPDATE masters SET position_id=$1 WHERE id=$2"
	if _, err := repo.db.Exec(query, positionId, masterId); err != nil {
		return fmt.Errorf("failed to update posititon of master wih id %d: %w", masterId, err)
	}
	return nil
}

func (repo *MasterRepository) UpdateBio(masterId int, bio string) error {
	query := "UPDATE masters SET bio=$1 WHERE id=$2"
	if _, err := repo.db.Exec(query, bio, masterId); err != nil {
		return fmt.Errorf("failed to update bio of master wih id %d: %w", masterId, err)
	}
	return nil
}
