package cmd

import (
	"github.com/JoseRodrigues443/is-my-team-awake/lib"
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
	for _, v := range members {
		isAwake := v.IsAwake(config)
		println(v.Name, " - ", v.Location, " - ", isAwake)
	}
}
