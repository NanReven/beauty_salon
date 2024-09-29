package domain

import (
	"database/sql/driver"
	"errors"
	"time"
)

const (
	timeLayout     = "02.01.2006 15:04:05"
	durationLayout = "15:04:05"
)

type CustomTime struct {
	time.Time
}

type CustomDuration struct {
	time.Time
}

func (ct CustomTime) MarshalJSON() ([]byte, error) {
	time := ct.Time.Format(timeLayout)
	return []byte(`"` + time + `"`), nil
}

func (cd CustomDuration) MarshalJSON() ([]byte, error) {
	time := cd.Time.Format(durationLayout)
	return []byte(`"` + time + `"`), nil
}

func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	str := string(b[1 : len(b)-1])
	time, err := time.Parse(timeLayout, str)
	if err != nil {
		return err
	}
	ct.Time = time
	return nil
}

func (cd *CustomDuration) UnmarshalJSON(b []byte) error {
	str := string(b[1 : len(b)-1])
	time, err := time.Parse(durationLayout, str)
	if err != nil {
		return err
	}
	cd.Time = time
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

func (cd *CustomDuration) Scan(src interface{}) error {
	switch value := src.(type) {
	case time.Time:
		cd.Time = value
	default:
		return errors.New("invalid duration")
	}
	return nil
}

func (ct CustomTime) Value() (driver.Value, error) {
	return ct.Time, nil
}

func (cd CustomDuration) Value() (driver.Value, error) {
	return cd.Time, nil
}
