package tests

import (
	"testing"

	"github.com/JoseRodrigues443/is-my-team-awake/lib"
)

func NewHourOfDay(hour int, minute int) *lib.HourOfDay {
	hourOfDay, _ := lib.NewHourOfDay(hour, minute)
	return hourOfDay
}

var timeDayCaseTests = []struct {
	name  string
	start *lib.HourOfDay
	end   *lib.HourOfDay
	want  bool
}{
	{"Is Equals", NewHourOfDay(11, 0), NewHourOfDay(11, 0), true},
	{"Is Different", NewHourOfDay(11, 0), NewHourOfDay(11, 0), true},
}

func TestCreateTimeOfDay(t *testing.T) {
	result, _ := lib.NewHourOfDay(8, 0)
	if result == nil {
		t.Error("result should be a valid time of day, got", result)
	}
}
func TestIsEqualsBasic(t *testing.T) {
	main, _ := lib.NewHourOfDay(8, 0)
	toCompare, _ := lib.NewHourOfDay(9, 0)
	got := main.IsEqualsTo(toCompare)
	if got {
		t.Error("result should be false, got", got)
	}
}

func TestIsEquals(t *testing.T) {
	for _, tt := range timeDayCaseTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.start.IsEqualsTo(tt.end)
			if got != tt.want {
				t.Error("result should be false, got", got)
			}
		})
	}
}
