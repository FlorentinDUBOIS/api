package cmd

import (
	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// RootCmd of bouncer
var RootCmd = &cobra.Command{
	Use:   "bouncer",
	Short: "bouncer command line interface",
	Run:   root,
}

func init() {
	cobra.OnInitialize(configure)

	RootCmd.PersistentFlags().
		BoolP("verbose", "v", false, "Set the output to verbose")

	if err := viper.BindPFlags(RootCmd.PersistentFlags()); err != nil {
		log.Debug(err)
	}
}

func configure() {
	if viper.GetBool("verbose") {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	viper.SetEnvPrefix("bouncer")
	viper.AutomaticEnv()

	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/bouncer")
	viper.AddConfigPath("$HOME/.bouncer")
	viper.AddConfigPath(".")

	if err := viper.MergeInConfig(); err != nil {
		log.Debug(err)
	}
}

func root(pCmd *cobra.Command, pArgs []string) {
	if err := pCmd.Help(); err != nil {
		log.Fatal(err)
	}
}
