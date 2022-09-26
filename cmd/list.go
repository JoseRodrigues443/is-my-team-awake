package cmd

import (
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
	repo := store.NewRepo()
	members := repo.GetAll()
	for _, v := range members {
		println(v.Name, " - ", v.Location)
	}
}
