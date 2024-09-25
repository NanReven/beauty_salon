package dto

import (
	"errors"
	"fmt"
	"time"
)

type CustomTime struct {
	time.Time
}

const customTimeFormat = "2006-01-02 15:04:05"

func (ct CustomTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf(`"%s"`, ct.Format(customTimeFormat))
	return []byte(formatted), nil
}

func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	str := string(b[1 : len(b)-1])

	t, err := time.Parse(customTimeFormat, str)
	if err != nil {
		return errors.New("invalid time format, use YYYY-MM-DD HH:MM:SS")
	}

	ct.Time = t
	return nil
}

type AppointmentInput struct {
	AppointmentStart CustomTime `json:"appointment_start" binding:"required"`
	MasterId         int        `json:"master_id" binding:"required"`
	Comment          string     `json:"comment"`
	Services         []int      `json:"services" binding:"required"`
}
