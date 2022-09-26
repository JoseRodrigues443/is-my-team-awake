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
	loc, error := time.LoadLocation(t.Location)

	if error != nil {
		logger.Log.Printf("Not valid location: %s", t.Location)
		return false
	}

	// set timezone,
	now := time.Now().In(loc)

	start := time.Date(now.Year(), now.Month(), now.Day(), c.StartOfDay.Hour, c.StartOfDay.Minute, 0, 0, loc)

	end := time.Date(now.Year(), now.Month(), now.Day(), c.EndOfDay.Hour, c.EndOfDay.Minute, 0, 0, loc)

	return now.After(start) && now.Before(end)
}
