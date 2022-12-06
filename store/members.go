package store

import (
	"strings"

	"github.com/JoseRodrigues443/is-my-team-awake/lib"
	utils "github.com/JoseRodrigues443/is-my-team-awake/utils/logger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Repo struct {
	members *gorm.DB
}

func NewRepo() *Repo {

	var path = "./db/is-my-team-awake.db"

	var fullPath = utils.GetProjectRootPath(path)

	db, err := gorm.Open(sqlite.Open(utils.GetProjectRootPath(fullPath)), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&lib.TeamMember{})

	return &Repo{
		members: db,
	}
}

func generateKey(name string) string {
	return strings.ReplaceAll(strings.ToLower(name), " ", "_")
}

func (r *Repo) First() (lib.TeamMember, error) {
	var to_return lib.TeamMember
	result := r.members.First(&to_return)
	return to_return, result.Error
}

func (r *Repo) GetAll() ([]lib.TeamMember, error) {
	var to_return []lib.TeamMember
	result := r.members.Find(&to_return)
	return to_return, result.Error
}

func (r *Repo) GetMemberByName(name string) (lib.TeamMember, error) {
	var to_return lib.TeamMember
	result := r.members.First(&to_return, "name = ?", name)
	return to_return, result.Error
}

func (r *Repo) AddMember(member lib.TeamMember) (lib.TeamMember, error) {
	result := r.members.Create(&member)
	return member, result.Error
}

func (r *Repo) RemoveMember(name string) (lib.TeamMember, error) {
	var to_return lib.TeamMember
	result := r.members.Where("name = ?", "name").Delete(&to_return)
	return to_return, result.Error
}

func (r *Repo) Total() int64 {
	var to_return int64
	r.members.Count(&to_return)
	return to_return
}
