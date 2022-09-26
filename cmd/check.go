package cmd

import (
	"fmt"
	"os"

	"github.com/JoseRodrigues443/is-my-team-awake/lib"
	"github.com/JoseRodrigues443/is-my-team-awake/store"
	"github.com/spf13/cobra"
)

var (
	check = &cobra.Command{
		Use:   "check",
		Short: "check if its awake",
		Long:  ``,
		Run:   cmdCheck,
	}
)

func cmdCheck(ccmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "No username is specified. Please specify a team member.")
		return
	}
	name := args[0]
	config := lib.GetConfig()
	repo := store.NewRepo()
	member, error := repo.GetMemberByName(name)
	if error != nil {
		fmt.Printf("The user %s was not found. Try the list command to find the correct name.\n", name)
		return
	}
	theirTime := member.TheirTime(config)
	isAwake := "Sleeping"
	if member.IsAwake(config) {
		isAwake = "Awake"
	}
	fmt.Printf("The Team member %s is %s, where he lives its %d:%d at %s \n", member.Name, isAwake, theirTime.Hour, theirTime.Minute, member.Location)
}
