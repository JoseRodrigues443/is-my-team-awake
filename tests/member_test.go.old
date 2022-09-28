package tests

import (
	"testing"

	"github.com/JoseRodrigues443/is-my-team-awake/lib"
)

func TestIsAwake(t *testing.T) {
	teamMember := &lib.TeamMember{
		Name:     "Jonh Doe",
		Location: "Asia/Shanghai",
	}
	startOfDay, _ := lib.NewHourOfDay(8, 0)
	endOfDay, _ := lib.NewHourOfDay(18, 0)
	config := &lib.Config{
		StartOfDay: startOfDay,
		EndOfDay:   endOfDay,
	}
	result := teamMember.IsAwake(config)
	if result {
		t.Error("result should be false, got", result)
	}
}
