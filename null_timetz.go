package gormpgtime

import (
	"database/sql/driver"
	"fmt"
	"math"
	"strings"
	"time"
)

type NullTimeTZ struct {
	Time  time.Time
	Valid bool
}

func (nt *NullTimeTZ) Scan(value interface{}) error {
	err := nt.parsePGTime(value.(string))
	nt.Valid = err == nil
	return nil
}

func (nt *NullTimeTZ) parsePGTime(val string) error {
	var err error
	sign := "+"
	tz := strings.Split(val, "+")
	if len(tz) < 2 {
		sign = "-"
		tz = strings.Split(val, "-")
	}
	nt.Time, err = time.Parse("15:04:05", tz[0])
	if err != nil {
		return err
	}
	if len(tz) > 1 {
		loc, err := TimeToLocation(sign + tz[1])
		if err != nil {
			return err
		}
		nt.Time, _ = time.ParseInLocation("15:04:05", nt.Time.Format("15:04:05"), loc)
	}
	return nil
}

// Value implements the driver Valuer interface.
func (nt NullTimeTZ) Value() (driver.Value, error) {
	if !nt.Valid { //t.Time.Format("15:04:05")
		return nil, nil
	}
	v := nt.Time.Format("15:04:05") + nt.zoneToPG()
	return v, nil
}

func (nt *NullTimeTZ) Set(t *time.Time) NullTimeTZ {
	if t != nil {
		nt.Time = *t
		nt.Valid = true
	}
	return *nt
}

func (nt *NullTimeTZ) zoneToPG() string {
	_, offset := nt.Time.Zone()
	s := "+"
	if offset < 0 {
		s = "-"
		offset = int(math.Abs(float64(offset)))
	}
	tm := time.Time{}.Add(time.Duration(offset) * time.Second)
	return fmt.Sprintf("%s%.2d:%.2d", s, tm.Hour(), tm.Minute())
}
