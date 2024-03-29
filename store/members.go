package store

import (
	"errors"
	"strings"

	"github.com/JoseRodrigues443/is-my-team-awake/lib"
)

var ErrNotFound = errors.New("Member was not found")

type Repo struct {
	membersById map[string]lib.TeamMember
}

func NewRepo() *Repo {
	membersById := make(map[string]lib.TeamMember)

	membersById["jonh_doe"] = lib.TeamMember{
		Name:     "Jonh Doe",
		Location: "Asia/Shanghai",
	}

	membersById["janet_doe"] = lib.TeamMember{
		Name:     "Janet Doe",
		Location: "Asia/Kolkata",
	}

	membersById["李祖阳"] = lib.TeamMember{
		Name:     "李祖阳",
		Location: "Europe/London",
	}

	return &Repo{
		membersById: membersById,
	}
}

func generateKey(name string) string {
	return strings.ReplaceAll(strings.ToLower(name), " ", "_")
}

func (r *Repo) GetAll() []lib.TeamMember {
	v := make([]lib.TeamMember, 0)
	for _, tx := range r.membersById {
		v = append(v, tx)
	}
	return v
}

func (r *Repo) GetMemberByName(name string) (*lib.TeamMember, error) {
	key := generateKey(name)
	member, ok := r.membersById[key]
	if ok {
		return &member, nil
	}
	return nil, ErrNotFound
}

func (r *Repo) AddMember(member lib.TeamMember) {
	key := generateKey(member.Name)
	r.membersById[key] = member
}

func (r *Repo) RemoveMember(name string) {
	key := generateKey(name)
	delete(r.membersById, key)
}
