package main

import (
	"go-api-template/src/internal/app/api-my-domain-com/api/cli"
	"go-api-template/src/internal/app/api-my-domain-com/api/server"
	"go-api-template/src/internal/app/api-my-domain-com/api/server/config"
	"log"
	"os"
)

func main() {
	// Reading the values in the command line keys
	cliConfig, err := cli.NewConfig()
	if err != nil {
		log.Fatalf(" -[exit]- cli NewConfig() error [%v]\n", err)
	}

	var logFile *os.File

	// We write the log to a file, not to the screen
	if len(cliConfig.LogFile) > 0 {
		logFile, err = os.OpenFile(cliConfig.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			log.Fatalf(" -[exit]- error OpenFile log file (%v): %v", cliConfig.LogFile, err)
		}

		log.SetOutput(logFile)

		defer func() {
			err = logFile.Close()
			if err != nil {
				log.Fatalf(" -[exit]- error Closing log file (%v): %v", cliConfig.LogFile, err)
			}
		}()
	}

	appConfig := new(config.Config)
	err = appConfig.Load(cliConfig.ConfigFile)
	if err != nil {
		log.Fatalf(" -[exit]- appConfig.Load() error [%v]\n", err)
	}

	srv := server.NewHTTPServer(&appConfig.Hosts.API, logFile)
	srv.Run()
}
