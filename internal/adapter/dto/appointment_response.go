package dto

type AppointmentResponse struct {
	Id               int        `db:"id"`
	AppointmentStart CustomTime `db:"appointment_start"`
	AppointmentEnd   CustomTime `db:"appointment_end"`
	Master           string     `db:"master"`
	Status           string     `db:"status"`
	Comment          string     `db:"comment"`
	Services         []FavourResponse
	TotalSum         float64 `db:"total_sum"`
}
