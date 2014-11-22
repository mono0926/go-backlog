package util

import "time"

func ParseDateString(d string) time.Time {
	t, _ := time.Parse("2006-01-02T15:04:05Z", d)
	return t
}
