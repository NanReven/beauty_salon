package dto

type ServiceResponse struct {
	Category string         `db:"category_title"`
	Title    string         `db:"service_title"`
	Duration CustomDuration `db:"duration"`
	Price    float64        `db:"price"`
}
