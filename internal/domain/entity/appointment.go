package entity

import "time"

const (
	StatusPending   = "pending"
	StatusConfirmed = "confirmed"
	StatusRejected  = "rejected"
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
