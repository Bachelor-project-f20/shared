package config

import (
	"errors"
	"log"
	"os"
	"strconv"

	"github.com/Bachelor-project-f20/go-outbox"
)

func extractFromEnvForDb(config *ConfigValues) {
	if v := os.Getenv("DB_CONNECTION_SLEEP"); v != "" {
		if i1, err := strconv.Atoi(v); err == nil {
			config.DatabaseConnectionSleepDuration = i1
		}
	}
	if v := os.Getenv("MYSQL_URL"); v != "" {
		config.DatabaseType = outbox.MySQL
		config.DatabaseConnection = v
	} else if v := os.Getenv("POSTGRES_URL"); v != "" {
		config.DatabaseType = outbox.Postgres
		config.DatabaseConnection = v
	}
}

func setupOutbox(config *ConfigValues, result *ConfigResult) error {
	if !config.UseOutbox {
		return nil
	}
	if len(config.OutboxModels) == 0 {
		log.Fatalln("setupOutbox: No models to persist")
		return errors.New("No models to persist")
	}
	if result.EventEmitter == nil {
		log.Fatalln("setupOutbox: No event emitter")
		return errors.New("No event emitter")
	}

	out, err := outbox.NewOutbox(config.DatabaseType, config.DatabaseConnection, config.DatabaseConnectionSleepDuration, result.EventEmitter, config.OutboxModels...)
	if err != nil {
		log.Fatalf("Error creating Outbox: %v \n", err)
	}
	result.Outbox = out

	return nil
}
