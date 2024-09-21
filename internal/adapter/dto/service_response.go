package dto

import "time"

type ServiceResponse struct {
	Category string    `db:"category_title"`
	Title    string    `db:"service_title"`
	Duration time.Time `db:"duration"`
	Price    float64   `db:"price"`
}
