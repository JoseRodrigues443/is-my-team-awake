package tests

import (
	"reflect"
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

	if reflect.ValueOf(got).IsZero() {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestAddMemberStore(t *testing.T) {
	repo := store.NewRepo()
	to_add := lib.TeamMember{
		Name:     "Test 1",
		Location: "Asia/Kolkata",
	}
	total_before := repo.Total()

	repo.AddMember(to_add)

	total_after := repo.Total()

	if total_before == total_after {
		t.Errorf("got %d, wanted %d", total_after, total_before+1)
	}
}

func TestAddMemberRepeatedStore(t *testing.T) {
	repo := store.NewRepo()
	first, _ := repo.First()
	to_add := lib.TeamMember{
		Name:     first.Name,
		Location: "Asia/Kolkata",
	}
	total_before := repo.Total()

	repo.AddMember(to_add)

	total_after := repo.Total()

	if total_before < total_after {
		t.Errorf("got %d, wanted %d", total_after, total_before)
	}
}
