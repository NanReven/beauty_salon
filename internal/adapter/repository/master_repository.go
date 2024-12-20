package repository

import (
	"beauty_salon/internal/domain/entity"
	"database/sql"
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
		if err == sql.ErrNoRows {
			return master, entity.ErrMasterNotFound
		}
		return master, fmt.Errorf("failed to get master wih id %d: %w", id, err)
	}
	return master, nil
}

func (repo *MasterRepository) GetMasterName(userId int) (string, error) {
	var masterName string
	query := "SELECT CONCAT(first_name, ' ', second_name) FROM users WHERE id=$1"
	if err := repo.db.Get(&masterName, query, userId); err != nil {
		if err == sql.ErrNoRows {
			return "", entity.ErrMasterNotFound
		}
		return "", fmt.Errorf("failed to get master name wih user id %d: %w", userId, err)
	}
	return masterName, nil
}

func (repo *MasterRepository) GetMasterAppointment(masterId int, appointmentId int) error {
	var id int
	query := "SELECT id FROM appointments WHERE masterId=$1 AND id=$2"
	if err := repo.db.Get(&id, query, masterId, appointmentId); err != nil {
		if err == sql.ErrNoRows {
			return entity.ErrAppointmentNotFound
		}
		return err
	}
	return nil
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

func (repo *MasterRepository) GetMasterEmail(masterId int) (string, error) {
	var email string
	query := "SELECT email FROM users WHERE id=(SELECT user_id FROM masters WHERE id=$1)"
	row := repo.db.QueryRow(query, masterId)
	if err := row.Scan(&email); err != nil {
		if err == sql.ErrNoRows {
			return "", entity.ErrMasterNotFound
		}
		return "", fmt.Errorf("failed to get email of master wih id %d: %w", masterId, err)
	}
	return email, nil
}

func (repo *MasterRepository) ReplyToAppointment(input *entity.AppointmentReply) error {
	query := "UPDATE appointments SET status=$1 WHERE id=$2"
	if _, err := repo.db.Exec(query, input.Status, input.AppointmentId); err != nil {
		return err
	}
	return nil
}
