package cmd

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/troyApart/fluid-character-generator/server"
)

var RootCmd = &cobra.Command{
	Run: startServer,
}

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.WarnLevel)
}

func startServer(cmd *cobra.Command, args []string) {
	stop := make(chan os.Signal)
	defer close(stop)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	srv, err := server.New()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Fatal("unable to initialize server")
	}

	// create a server manager and start all the servers
	serverManager := server.NewServerManager(srv)
	serverManager.ShutdownTimeout = time.Duration(1) * time.Millisecond
	go func() {
		log.Info("starting server")
		err := serverManager.Start()
		if err != nil {
			log.WithFields(log.Fields{
				"error": err,
			}).Fatal("start error")
		}
	}()

	// wait for sigint or sigterm
	<-stop

	// stop traffic
	log.Printf("shutting down servers")
	err = serverManager.Stop()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Fatal("error occurred when stopping")
	}
}
