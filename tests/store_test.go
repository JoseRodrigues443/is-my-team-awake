package tests

import (
	"testing"

	"github.com/JoseRodrigues443/is-my-team-awake/lib"
	"github.com/JoseRodrigues443/is-my-team-awake/store"
)

func TestGetMemberStore(t *testing.T) {
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

func TestAddMemberStore(t *testing.T) {
	repo := store.NewRepo()
	to_add := lib.TeamMember{
		Name:     "Test 1",
		Location: "Asia/Kolkata",
	}
	total_before := len(repo.GetAll())

	repo.AddMember(to_add)

	total_after := len(repo.GetAll())

	if total_before == total_after {
		t.Errorf("got %d, wanted %d", total_after, total_before+1)
	}
}

func TestAddMemberRepeatedStore(t *testing.T) {
	repo := store.NewRepo()
	all := repo.GetAll()
	to_add := lib.TeamMember{
		Name:     all[0].Name,
		Location: "Asia/Kolkata",
	}
	total_before := len(all)

	repo.AddMember(to_add)

	total_after := len(repo.GetAll())

	if total_before < total_after {
		t.Errorf("got %d, wanted %d", total_after, total_before)
	}
}
