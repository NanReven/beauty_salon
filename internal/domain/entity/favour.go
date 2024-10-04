package entity

import (
	"beauty_salon/internal/domain"
)

type Category struct {
	Id          int
	Title       string
	Description string
	Slug        string
}

type Favour struct {
	Id         int
	CategoryId int                   `json:"category_id" binding:"required"`
	Title      string                `json:"title" binding:"required"`
	Duration   domain.CustomDuration `json:"duration" binding:"required"`
	Price      float64               `json:"price" binding:"required"`
}

type FavourResponse struct {
	Category string                `db:"category_title"`
	Title    string                `db:"service_title"`
	Duration domain.CustomDuration `db:"duration"`
	Price    float64               `db:"price"`
}

type FavourUpdate struct {
	CategoryId int                   `json:"category_id"`
	Title      string                `json:"title"`
	Duration   domain.CustomDuration `json:"duration"`
	Price      float64               `json:"price"`
}
