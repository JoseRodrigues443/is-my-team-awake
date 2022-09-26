package lib

import (
	"errors"
	"fmt"
)

const (
	minHour   = 0
	maxHour   = 23
	minMinute = 0
	maxMinute = 59
)

type HourOfDay struct {
	Hour   int
	Minute int
}

func NewHourOfDay(hour int, minute int) (*HourOfDay, error) {
	if hour < minHour || hour > maxHour || minute < minMinute || minute > maxMinute {
		return nil, errors.New(fmt.Sprintf("Time Out of bounds, Hour should be between %d and %d, and minute should be between %d and %d...", minHour, maxHour, minMinute, maxMinute))
	}
	return &HourOfDay{hour, minute}, nil
}
