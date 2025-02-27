package util

import "time"

func ParseStringToTime(str string) (time.Time, error) {
	return time.Parse("2006-01-02", str)
}
