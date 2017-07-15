package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/FlorentinDUBOIS/bouncer/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
