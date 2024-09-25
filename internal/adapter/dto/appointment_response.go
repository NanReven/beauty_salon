package dto

type AppointmentResponse struct {
	Id               int    `db:"id"`
	AppointmentStart string `db:"appointment_start"`
	AppointmentEnd   string `db:"appointment_end"`
	Master           string `db:"master"`
	Status           string `db:"status"`
	Comment          string `db:"comment"`
	Services         []ServiceResponse
	TotalSum         float64 `db:"total_sum"`
}
