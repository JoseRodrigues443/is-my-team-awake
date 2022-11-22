package tests

import (
	"testing"

	"github.com/JoseRodrigues443/is-my-team-awake/lib"
)

func CreateTimeOfDay(t *testing.T) {
	result, _ := lib.NewHourOfDay(8, 0)
	if result == nil {
		t.Error("result should be a valid time of day, got", result)
	}
}
