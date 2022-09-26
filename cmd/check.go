package cmd

import (
	"fmt"
	"os"

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

}
