package helper

import (
	"time"
)

const (
	timeFormat = "2006-01-02T15:04:05.999Z07:00"
)

func DecodeTime(input string) (time.Time, error) {
	t, err := time.Parse(timeFormat, input)
	return t, err
}
