package dto

import (
	"database/sql/driver"
	"errors"
	"time"
)

type CustomTime struct {
	time.Time
}

const layout = "02.01.2006 15:04:05"

func (ct CustomTime) MarshalJSON() ([]byte, error) {
	time := ct.Time.Format(layout)
	return []byte(`"` + time + `"`), nil
}

func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	str := string(b[1 : len(b)-1])
	time, err := time.Parse(layout, str)
	if err != nil {
		return err
	}
	ct.Time = time
	return nil
}

func (ct *CustomTime) Scan(src interface{}) error {
	switch value := src.(type) {
	case time.Time:
		ct.Time = value
	default:
		return errors.New("invalid time")
	}
	return nil
}

func (ct CustomTime) Value() (driver.Value, error) {
	return ct.Time, nil
}
