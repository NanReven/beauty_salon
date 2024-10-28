package entity

import "errors"

type Position struct {
	Id    int
	Title string
}

var (
	ErrMasterNotFound     = errors.New("master not found")
	ErrInvalidMasterInput = errors.New("invalid request data")
)

type Master struct {
	Id         int
	UserId     int    `json:"user_id" binding:"required"`
	PositionId int    `json:"position_id" binding:"required"`
	Bio        string `json:"bio" binding:"required"`
	Slug       string
}

type MasterResponse struct {
	FirstName  string `db:"first_name"`
	SecondName string `db:"second_name"`
	Email      string `db:"email"`
	Position   string `db:"title"`
	Bio        string `db:"bio"`
}

type MasterUpdate struct {
	UserId     int    `json:"user_id"`
	PositionId int    `json:"position_id"`
	Bio        string `json:"bio"`
}

type AppointmentReply struct {
	AppointmentId int `json:"appointment_id"`
	Status        string
}
