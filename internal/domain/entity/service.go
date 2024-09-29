package entity

import (
	"beauty_salon/internal/domain"
	"time"
)

type Category struct {
	Id          int
	Title       string
	Description string
	Slug        string
}

type Service struct {
	Id         int
	CategoryId int
	Title      string
	Duration   time.Duration
	Price      float64
}

type FavourResponse struct {
	Category string                `db:"category_title"`
	Title    string                `db:"service_title"`
	Duration domain.CustomDuration `db:"duration"`
	Price    float64               `db:"price"`
}
