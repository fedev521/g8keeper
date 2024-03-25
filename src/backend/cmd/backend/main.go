package main

import (
	"fmt"
	"io"
	"os"

	"github.com/fedev521/g8keeper/backend/internal/log"
	"github.com/fedev521/g8keeper/backend/internal/srv"
	"github.com/spf13/pflag"
)

const (
	exitError      = 1
	exitUnexpected = 125
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(exitUnexpected)
		}
	}()
	if err := run(os.Args, os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(exitError)
	}
}

func run(_ []string, _ io.Reader, _ io.Writer) error {
	// init viper and pflag
	configureDefaultSettings()

	// parse CLI arguments
	pflag.Parse()

	// load configuration
	config, err := loadConfiguration()
	if err != nil {
		return err
	}

	// now configuration is loaded, but not necessarily valid

	logger := log.NewLogger(config.Log) // create logger (log config is valid)
	log.SetStandardLogger(logger)       // override the global logger
	log.SetDefaultLogger(logger)        // set the internal default logger

	log.Debug("Loaded configuration")

	// validate configuration
	err = config.Validate()
	if err != nil {
		log.Error(err.Error())
		return fmt.Errorf("configuration is invalid: %w", err)
	}

	log.Info("App started", map[string]interface{}{
		"srv":     config.App,
		"tinksvc": config.TinkSvc,
	})

	log.Info("Setup completed")

	err = srv.StartServer(config.App, config.TinkSvc)
	if err != nil {
		log.Error(err.Error())
		return fmt.Errorf("could not start server: %w", err)
	}

	return nil
}
