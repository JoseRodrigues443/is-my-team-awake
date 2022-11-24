package cmd

import (
	"fmt"

	"github.com/JoseRodrigues443/is-my-team-awake/lib"
	"github.com/JoseRodrigues443/is-my-team-awake/store"
	"github.com/spf13/cobra"
)

var (
	check       = buildCheckCommand()
	NameToCheck string
)

func buildCheckCommand() *cobra.Command {
	toReturn := &cobra.Command{
		Use:   "check",
		Short: "check if its awake",
		Long:  ``,
		Run:   cmdCheck,
	}
	toReturn.Flags().StringVarP(&NameToCheck, "name", "n", "", "Name of the team member to check (required)")
	_ = toReturn.MarkFlagRequired("name")

	return toReturn
}

func cmdCheck(ccmd *cobra.Command, args []string) {
	config := lib.GetConfig()
	repo := store.NewRepo()
	member, error := repo.GetMemberByName(NameToCheck)
	if error != nil {
		fmt.Printf("The user %s was not found. Try the list command to find the correct name.\n", NameToCheck)
		return
	}
	theirTime := member.TheirTime(config)
	isAwake := "Sleeping"
	if member.IsAwake(config) {
		isAwake = "Awake"
	}
	fmt.Printf("The Team member %s is %s, where he lives its %d:%d at %s \n", member.Name, isAwake, theirTime.Hour, theirTime.Minute, member.Location)
}
