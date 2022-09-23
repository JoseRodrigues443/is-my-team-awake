package tests

import (
	"testing"
	"time"

	"github.com/JoseRodrigues443/is-my-team-awake/lib"
)

func TestIsAwake(t *testing.T) {
	teamMember := &lib.TeamMember{
		Name:     "Jonh Doe",
		Location: "Asia/Shanghai",
	}
	config := &lib.Config{
		StartOfDay: time.Now(),
		EndOfDay:   time.Now().Add(time.Hour),
	}
	result := teamMember.IsAwake(config)
	if result {
		t.Error("result should be false, got", result)
	}
}
