package config

import (
	"log"
	"os"

	"github.com/Bachelor-project-f20/eventToGo"
	etgNats "github.com/Bachelor-project-f20/eventToGo/nats"
)

func extractFromEnvForQueue(config *ConfigValues) {
	if v := os.Getenv("NATS_BROKER_URL"); v != "" {
		config.MessageBrokerType = eventToGo.NATS
		config.MessageBrokerConnection = v
	}
}

func setupEventEmitterAndListener(config *ConfigValues, result *ConfigResult) error {
	if config.MessageBrokerType == eventToGo.NATS {
		err := setupNats(config, result)
		if err != nil {
			return err
		}
		return nil
	}
	log.Println("setupEventEmitterAndListener DONE")
	return nil
}

func setupNats(config *ConfigValues, result *ConfigResult) error {
	log.Println("Setting up nats")
	natsHandler := etgNats.NatsHandler{}
	if config.UseEmitter && config.UseListener {
		e, l, err := natsHandler.SetupEmitterAndListener(config.Exchange, config.QueueType, config.MessageBrokerConnection)
		if err != nil {
			return err
		}
		result.EventEmitter = e
		result.EventListener = l
	} else if config.UseEmitter {
		e, err := natsHandler.SetupEmitter(config.Exchange, config.QueueType, config.MessageBrokerConnection)
		if err != nil {
			log.Fatalf("Error creating event emitter: %v \n", err)
			return err
		}
		result.EventEmitter = e
	} else if config.UseListener {
		l, err := natsHandler.SetupListener(config.Exchange, config.QueueType, config.MessageBrokerConnection)
		if err != nil {
			log.Fatalf("Error creating event listener: %v \n", err)
			return err
		}
		result.EventListener = l
	}
	log.Println("SHARED: setup nats done")
	return nil
}
