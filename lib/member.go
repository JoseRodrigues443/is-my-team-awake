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
	theirTime := t.TheirTime(c)

	return theirTime.Hour >= c.StartOfDay.Hour && theirTime.Hour <= c.EndOfDay.Hour
}

func (t *TeamMember) TheirTime(c *Config) *HourOfDay {
	loc, error := time.LoadLocation(t.Location)

	if error != nil {
		logger.Log.Printf("Not valid location: %s", t.Location)
	}

	// set timezone,
	now := time.Now().In(loc)

	return &HourOfDay{
		Hour:   now.Hour(),
		Minute: now.Minute(),
	}
}
