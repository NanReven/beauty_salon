package dto

type AppointmentInput struct {
	AppointmentStart CustomTime `json:"appointment_start" binding:"required"`
	MasterId         int        `json:"master_id" binding:"required"`
	Comment          string     `json:"comment"`
	Services         []int      `json:"services" binding:"required"`
}
