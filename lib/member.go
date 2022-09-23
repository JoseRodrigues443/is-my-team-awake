package lib

import (
	"log"
	"time"
)

type TeamMember struct {
	Name     string
	Location string
}

func (t *TeamMember) IsAwake(c *Config) bool {
	loc, _ := time.LoadLocation(t.Location)

	// set timezone,
	now := time.Now().In(loc)
	log.Print("This is our first log message in Go.")

	start := time.Date(now.Year(), now.Month(), now.Day(), c.StartOfDay.UTC().Hour(), c.StartOfDay.UTC().Minute(), 0, 0, loc)

	end := time.Date(now.Year(), now.Month(), now.Day(), c.EndOfDay.UTC().Hour(), c.EndOfDay.UTC().Minute(), 0, 0, loc)

	return now.After(start) && now.Before(end)
}
