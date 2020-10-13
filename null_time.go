package gormpgtime

import (
	"database/sql/driver"
	"time"
)

type NullTime struct {
	Time  time.Time
	Valid bool
}

func (nt *NullTime) Scan(value interface{}) error {
	var err error
	nt.Time, err = time.Parse("15:04:05", value.(string))
	nt.Valid = err == nil
	return nil
}

func (nt NullTime) Value() (driver.Value, error) {
	if !nt.Valid { //
		return nil, nil
	}
	return nt.Time.Format("15:04:05"), nil
}

func (nt *NullTime) Set(t *time.Time) NullTime {
	if t != nil {
		nt.Time = *t
		nt.Valid = true
	}
	return *nt
}
