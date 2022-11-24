package cmd

import (
	"fmt"

	"github.com/JoseRodrigues443/is-my-team-awake/lib"
	"github.com/JoseRodrigues443/is-my-team-awake/store"
	"github.com/spf13/cobra"
)

var (
	add           = buildAddCommand()
	NameToAdd     string
	LocationToAdd string
)

func buildAddCommand() *cobra.Command {
	toReturn := &cobra.Command{
		Use:   "add",
		Short: "add a team member",
		Long:  ``,
		Run:   cmdAdd,
	}
	toReturn.Flags().StringVarP(&NameToAdd, "name", "n", "", "Name of the team member (required)")
	toReturn.Flags().StringVarP(&LocationToAdd, "location", "l", "", "Time zone location (required)")

	_ = toReturn.MarkFlagRequired("name")
	_ = toReturn.MarkFlagRequired("location")

	return toReturn
}

func cmdAdd(ccmd *cobra.Command, args []string) {
	repo := store.NewRepo()
	teamMember := lib.TeamMember{
		Name:     NameToAdd,
		Location: LocationToAdd,
	}
	repo.AddMember(teamMember)
	fmt.Printf("Saved")
}
