package main

import (
	"math/rand"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/troyApart/fluid-character-generator/cmd"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	if err := cmd.RootCmd.Execute(); err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Fatal("unable to execute command")
	}
}
