package cmd

import (
	"fmt"

	"github.com/JoseRodrigues443/is-my-team-awake/lib"
	"github.com/JoseRodrigues443/is-my-team-awake/logger"
	"github.com/JoseRodrigues443/is-my-team-awake/store"
	"github.com/spf13/cobra"
)

var (
	list = &cobra.Command{
		Use:   "list",
		Short: "list if its awake",
		Long:  ``,
		Run:   cmdList,
	}
)

func cmdList(ccmd *cobra.Command, args []string) {
	config := lib.GetConfig()
	repo := store.NewRepo()
	members := repo.GetAll()
	logger.Log.Printf("cmd.List members: %d", len(members))
	isAwake := "Sleeping"
	for _, v := range members {
		if v.IsAwake(config) {
			isAwake = "Awake"
		}
		// fmt.Printf("%s - ", v.Name)
		// fmt.Printf("%s - ", v.Location)
		// fmt.Printf("%s - ", isAwake)
		fmt.Printf("%s - %s - %s", v.Name, v.Location, isAwake)
		fmt.Println("")
	}
}
