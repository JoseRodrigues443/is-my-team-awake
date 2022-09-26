package lib

import (
	"time"

	"github.com/JoseRodrigues443/is-my-team-awake/logger"
)

type TeamMember struct {
	Name     string
	Location string
}

func (t *TeamMember) IsAwake(c *Config) bool {
	loc, _ := time.LoadLocation(t.Location)

	// set timezone,
	now := time.Now().In(loc)
	logger.Log.Printf("This is our first log message in Go.")

	start := time.Date(now.Year(), now.Month(), now.Day(), c.StartOfDay.Hour, c.StartOfDay.Minute, 0, 0, loc)

	end := time.Date(now.Year(), now.Month(), now.Day(), c.EndOfDay.Hour, c.EndOfDay.Minute, 0, 0, loc)

	return now.After(start) && now.Before(end)
}
