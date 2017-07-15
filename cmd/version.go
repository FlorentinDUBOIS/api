package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	semvers = "0.0.1"
	githash = "HEAD"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Get the bouncer version",
	Run:   version,
}

func init() {
	RootCmd.AddCommand(versionCmd)
}

func version(pCmd *cobra.Command, pArgs []string) {
	fmt.Printf("version %s (%s)\n", semvers, githash)
}
