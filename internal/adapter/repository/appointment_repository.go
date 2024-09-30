package repository

import (
	"beauty_salon/internal/domain/entity"
	"errors"
	"time"

	"github.com/jmoiron/sqlx"
)

type AppointmentRepository struct {
	db *sqlx.DB
}

func NewAppointmentRepository(db *sqlx.DB) *AppointmentRepository {
	return &AppointmentRepository{db: db}
}

const (
	pendingStatus   = "pending"
	acceptedStatus  = "accepted"
	completedStatus = "completed"
	cancelledStatus = "cancelled"
)

func (repo *AppointmentRepository) CreateAppointment(userId int, appointment *entity.AppointmentInput, appointmentEnd time.Time, totalSum float64) (int, error) {
	var appointmentId int
	tx, err := repo.db.Begin()
	if err != nil {
		return 0, err
	}
	query := "INSERT INTO appointments (appointment_start, appointment_end, user_id, master_id, status, comment, total_sum) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id"
	row := tx.QueryRow(query, appointment.AppointmentStart.Time, appointmentEnd, userId, appointment.MasterId, pendingStatus, appointment.Comment, totalSum)
	if err := row.Scan(&appointmentId); err != nil {
		tx.Rollback()
		return 0, err
	}

	for _, serviceId := range appointment.Services {
		insert := "INSERT INTO appointmentservices (appointment_id, service_id) VALUES ($1, $2)"
		if _, err := tx.Exec(insert, appointmentId, serviceId); err != nil {
			tx.Rollback()
			return 0, err
		}
	}
	return appointmentId, tx.Commit()
}

func (repo *AppointmentRepository) GetAllAppointments(userId int) ([]entity.AppointmentResponse, error) {
	var appointments []entity.AppointmentResponse
	query := "SELECT appointments.id, appointment_start, appointment_end, CONCAT(first_name, ' ', second_name) AS master, status, comment, total_sum FROM appointments JOIN masters ON (master_id = masters.id) JOIN users ON (masters.user_id = users.id) WHERE appointments.user_id = $1"
	if err := repo.db.Select(&appointments, query, userId); err != nil {
		return appointments, err
	}
	return appointments, nil
}

func (repo *AppointmentRepository) GetAppointmentById(userId, appointmentId int) (entity.AppointmentResponse, error) {
	var appointment entity.AppointmentResponse
	query := "SELECT appointments.id, appointment_start, appointment_end, CONCAT(first_name, ' ', second_name) AS master, status, comment, total_sum FROM appointments JOIN masters ON (master_id = masters.id) JOIN users ON (masters.user_id = users.id) WHERE appointments.user_id = $1 AND appointments.id = $2"
	if err := repo.db.Get(&appointment, query, userId, appointmentId); err != nil {
		return appointment, err
	}
	return appointment, nil
}

func (repo *AppointmentRepository) GetFavoursByAppointmentId(appointmentId int) ([]entity.FavourResponse, error) {
	var services []entity.FavourResponse
	query := "SELECT categories.title AS category_title, services.title AS service_title, duration, price FROM services JOIN categories ON (category_id = categories.id) WHERE services.id IN (SELECT service_id FROM appointmentservices WHERE appointment_id = $1)"
	if err := repo.db.Select(&services, query, appointmentId); err != nil {
		return services, err
	}
	return services, nil
}

func (repo *AppointmentRepository) CancelAppointment(userId, appointmentId int) (string, error) {
	var status string
	query := "UPDATE appointments SET status=$1 WHERE id = $2 AND user_id = $3 RETURNING status"
	if err := repo.db.Get(&status, query, cancelledStatus, appointmentId, userId); err != nil {
		return "", errors.New("user has no appointment with this id")
	}
	return status, nil
}
