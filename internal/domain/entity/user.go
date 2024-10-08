package entity

import "errors"

var (
	ErrUserNotFound = errors.New("user not found")
)

type User struct {
	Id         int
	FirstName  string `json:"first_name" binding:"required" db:"first_name"`
	SecondName string `json:"second_name" binding:"required" db:"second_name"`
	Email      string `json:"email" binding:"required" db:"email"`
	Password   string `json:"password" binding:"required" db:"password_hash"`
	Role       string `json:"role" db:"role"`
}
