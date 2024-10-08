package entity

import (
	"beauty_salon/internal/domain"
	"errors"
	"time"
)

var (
	ErrAppointmentNotFound  = errors.New("appointment not found")
	ErrAppointmentCancelled = errors.New("appointment already cancelled")
)

type Appointment struct {
	Id       int
	Start    time.Time
	End      time.Time
	UserId   int
	MasterId int
	StatusId int
	Comment  string
	TotalSum float64
}

type AppointmentService struct {
	AppointmentId int
	ServiceId     int
}

type AppointmentInput struct {
	AppointmentStart domain.CustomTime `json:"appointment_start" binding:"required"`
	MasterId         int               `json:"master_id" binding:"required"`
	Comment          string            `json:"comment"`
	Services         []int             `json:"services" binding:"required"`
}

type AppointmentResponse struct {
	Id               int               `db:"id"`
	AppointmentStart domain.CustomTime `db:"appointment_start"`
	AppointmentEnd   domain.CustomTime `db:"appointment_end"`
	Master           string            `db:"master"`
	Status           string            `db:"status"`
	Comment          string            `db:"comment"`
	Services         []FavourResponse
	TotalSum         float64 `db:"total_sum"`
}
