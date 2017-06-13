package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	version = "0.0.1"
	githash = "HEAD"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("bouncer %s\n (%s)", version, githash)
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
