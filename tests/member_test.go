package tests

import (
	"testing"
	"time"

	"github.com/JoseRodrigues443/is-my-team-awake/lib"
)

const name = "Jonh Doe"

var config *lib.Config = newConfig()

type mockedClock struct{}

// It's mocking the time.Now() function to use the day one of 3:00 of January of 2022 UTC
func (mockedClock) Now() time.Time {
	hour := 3
	return time.Date(2022, time.January, 1, hour, 0, 0, 0, time.UTC)
}

func newConfig() *lib.Config {
	config := &lib.Config{
		StartOfDay: NewHourOfDay(8, 0),
		EndOfDay:   NewHourOfDay(18, 0), // work time ended
		Clock:      mockedClock{},
	}
	return config
}

var memberCaseTests = []struct {
	location  string
	isAwake   bool
	theirTime *lib.HourOfDay
}{
	{"Asia/Shanghai", true, NewHourOfDay(11, 0)},    // Asia/Shanghai UTC+08:00 then its 03:00 + 8 hours = 11:00 its awake
	{"Europe/London", false, NewHourOfDay(3, 0)},    // Europe/London UTC+00:00 then its 03:00 + 0 hours = 03:00 its asleep
	{"Atlantic/Azores", false, NewHourOfDay(2, 0)},  // Atlantic/Azores UTC-01:00 then its 03:00 - 1 hours = 02:00 its asleep
	{"Pacific/Honolulu", true, NewHourOfDay(17, 0)}, // Pacific/Honolulu UTC-10:00 then its 03:00 - 10 hours = 17:00 its awake
	{"Australia/Perth", true, NewHourOfDay(11, 0)},  // Same UTC as Asia/Shanghai, then should be 11:00 its awake
}

func TestIsAwakeBasic(t *testing.T) {
	teamMember := lib.TeamMember{
		Name:     name,
		Location: "Asia/Shanghai",
	}
	want := true
	got := teamMember.IsAwake(config)
	if got != want {
		t.Errorf("result should be %t, got %t", want, got)
	}
}

func TestIsAwake(t *testing.T) {
	for _, tt := range memberCaseTests {
		t.Run(tt.location, func(t *testing.T) {
			teamMember := lib.TeamMember{
				Name:     name,
				Location: tt.location,
			}
			got := teamMember.IsAwake(config)
			if got != tt.isAwake {
				t.Errorf("IsAwake(%s, %s) got %v, want %v", name, tt.location, got, tt.isAwake)
			}
		})
	}
}

func TestTheirTimeBasic(t *testing.T) {
	teamMember := lib.TeamMember{
		Name:     name,
		Location: "Asia/Shanghai",
	}
	want := NewHourOfDay(11, 0)
	result := teamMember.TheirTime(config)
	if !result.IsEqualsTo(want) {
		t.Errorf("result should be %d, got %d", want, result)
	}
}

func TestTheirTime(t *testing.T) {
	for _, tt := range memberCaseTests {
		t.Run(tt.location, func(t *testing.T) {
			teamMember := lib.TeamMember{
				Name:     name,
				Location: tt.location,
			}
			got := teamMember.TheirTime(config)
			if !got.IsEqualsTo(tt.theirTime) {
				t.Errorf("TheirTime(%s, %s) got %v, want %v", name, tt.location, got, tt.theirTime)
			}
		})
	}
}
