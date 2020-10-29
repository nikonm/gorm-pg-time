package gormpgtime

import (
	"strings"
	"time"
)

// Convert string timezone offset to *time.Location
// tz format example "+02:30", "-4:30", "03:30"
func TimeToLocation(tz string) (*time.Location, error) {
	sign := "+"
	tm := strings.TrimPrefix(tz, "+")
	if strings.HasPrefix(tz, "-") {
		sign = "-"
		tm = strings.TrimPrefix(tz, "-")
	}
	if !strings.Contains(tm, ":") {
		tm += ":00"
	}
	zn, err := time.Parse("15:04", tm)
	if err != nil {
		return nil, err
	}
	secOffset := time.Duration(zn.Hour()*3600+zn.Minute()*60+zn.Second()) * time.Second

	if sign == "+" {
		return time.FixedZone("", int(secOffset.Seconds())), nil
	}
	return time.FixedZone("", int(-secOffset.Seconds())), nil
}
