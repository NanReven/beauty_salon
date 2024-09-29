package repository

import (
	"beauty_salon/internal/domain/entity"
	"errors"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type AppointmentRepository struct {
	db *sqlx.DB
}

func NewAppointmentRepository(db *sqlx.DB) *AppointmentRepository {
	return &AppointmentRepository{db: db}
}

func (repo *AppointmentRepository) CreateAppointment(userId int, appointment *entity.AppointmentInput) (int, error) {
	appointmentEnd := appointment.AppointmentStart.Time
	var appointmentId int
	var totalSum float64

	tx, err := repo.db.Begin()
	if err != nil {
		return 0, err
	}

	for _, id := range appointment.Services {
		var serviceDuration time.Time
		var servicePrice float64
		query := fmt.Sprintf("SELECT duration, price FROM %s WHERE id=$1", servicesTable)
		row := tx.QueryRow(query, id)
		if err := row.Scan(&serviceDuration, &servicePrice); err != nil {
			tx.Rollback()
			return 0, err
		}
		duration := time.Duration(serviceDuration.Hour()*3600+serviceDuration.Minute()*60) * time.Second
		appointmentEnd = appointmentEnd.Add(duration)
		totalSum += servicePrice
	}

	query := fmt.Sprintf("INSERT INTO %s (appointment_start, appointment_end, user_id, master_id, status, comment, total_sum) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id", appointmentsTable)
	row := tx.QueryRow(query, appointment.AppointmentStart.Time, appointmentEnd, userId, appointment.MasterId, "pending", appointment.Comment, totalSum)
	if err := row.Scan(&appointmentId); err != nil {
		tx.Rollback()
		return 0, err
	}

	for _, serviceId := range appointment.Services {
		insert := fmt.Sprintf("INSERT INTO %s (appointment_id, service_id) VALUES ($1, $2)", appointmentServicesTable)
		if _, err := tx.Exec(insert, appointmentId, serviceId); err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	return appointmentId, tx.Commit()
}

func (repo *AppointmentRepository) GetAllAppointments(userId int) ([]entity.AppointmentResponse, error) {
	var appointments []entity.AppointmentResponse
	query := fmt.Sprintf("SELECT appointments.id, appointment_start, appointment_end, CONCAT(first_name, ' ', second_name) AS master, status, comment, total_sum FROM %s JOIN %s ON (master_id = masters.id) JOIN %s ON (masters.user_id = users.id) WHERE appointments.user_id = $1", appointmentsTable, mastersTable, usersTable)
	if err := repo.db.Select(&appointments, query, userId); err != nil {
		return appointments, err
	}
	for appointmentIndex, appointment := range appointments {
		var services []entity.FavourResponse
		servicesQuery := fmt.Sprintf("SELECT categories.title AS category_title, services.title AS service_title, duration, price FROM %s JOIN %s ON (category_id = categories.id) WHERE services.id IN (SELECT service_id FROM %s WHERE appointment_id = $1)", servicesTable, categoriesTable, appointmentServicesTable)
		if err := repo.db.Select(&services, servicesQuery, appointment.Id); err != nil {
			return appointments, err
		}
		appointments[appointmentIndex].Services = services
	}
	return appointments, nil
}

func (repo *AppointmentRepository) GetAppointmentById(userId, appointmentId int) (entity.AppointmentResponse, error) {
	var appointment entity.AppointmentResponse
	var services []entity.FavourResponse
	query := fmt.Sprintf("SELECT appointments.id, appointment_start, appointment_end, CONCAT(first_name, ' ', second_name) AS master, status, comment, total_sum FROM %s JOIN %s ON (master_id = masters.id) JOIN %s ON (masters.user_id = users.id) WHERE appointments.user_id = $1 AND appointments.id = $2", appointmentsTable, mastersTable, usersTable)

	if err := repo.db.Get(&appointment, query, userId, appointmentId); err != nil {
		return appointment, err
	}

	servicesQuery := fmt.Sprintf("SELECT categories.title AS category_title, services.title AS service_title, duration, price FROM %s JOIN %s ON (category_id = categories.id) WHERE services.id IN (SELECT service_id FROM %s WHERE appointment_id = $1)", servicesTable, categoriesTable, appointmentServicesTable)
	if err := repo.db.Select(&services, servicesQuery, appointmentId); err != nil {
		return appointment, err
	}

	appointment.Services = services
	return appointment, nil
}

func (repo *AppointmentRepository) CancelAppointment(userId, appointmentId int) (string, error) {
	var status string
	query := fmt.Sprintf("UPDATE %s SET status='cancelled' WHERE id = $1 AND user_id = $2 RETURNING status", appointmentsTable)
	if err := repo.db.Get(&status, query, appointmentId, userId); err != nil {
		return "", errors.New("user has no appointment with this id")
	}
	return status, nil
}
