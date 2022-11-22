package tests

import (
	"testing"

	"github.com/JoseRodrigues443/is-my-team-awake/lib"
	"github.com/JoseRodrigues443/is-my-team-awake/store"
)

func TestStore(t *testing.T) {

	repo := store.NewRepo()
	got, _ := repo.GetMemberByName("Janet Doe")

	want := lib.TeamMember{
		Name:     "Janet Doe",
		Location: "Asia/Kolkata",
	}

	if got == nil {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
