package config

import (
	"fmt"
	"log"
	"os"

	"github.com/Bachelor-project-f20/eventToGo"
	etgNats "github.com/Bachelor-project-f20/eventToGo/nats"
	"github.com/nats-io/nats.go"
)

func extractFromEnvForQueue(config *ConfigValues) {
	if v := os.Getenv("NATS_BROKER_URL"); v != "" {
		config.MessageBrokerType = eventToGo.NATS
		config.MessageBrokerConnection = v
	}
}

func setupEventEmitterAndListener(config *ConfigValues, result *ConfigResult) error {
	if config.MessageBrokerType == eventToGo.NATS {
		log.Println("Setting up nats")
		encodedConn, err := setupNatsConn(config)
		if err != nil {
			log.Fatalf("Error connecting to Nats: %v \n", err)
			return err
		}
		if config.UseEmitter {
			result.EventEmitter, err = etgNats.NewNatsEventEmitter(encodedConn, config.Exchange, config.QueueType)
			if err != nil {
				log.Fatalf("Error creating event emitter: %v \n", err)
			}
		}
		if config.UseListener {
			result.EventListener, err = etgNats.NewNatsEventListener(encodedConn, config.Exchange, config.QueueType)
			if err != nil {
				log.Fatalf("Error creating event listener: %v \n", err)
			}
		}

	}
	log.Println("setupEventEmitterAndListener DONE")
	return nil
}

func setupNatsConn(config *ConfigValues) (*nats.EncodedConn, error) {

	natsConn, err := nats.Connect(config.MessageBrokerConnection)

	if err != nil {
		fmt.Println("Connection to Nats failed")
		return nil, err
	}

	encodedConn, err := nats.NewEncodedConn(natsConn, "json")

	if err != nil {
		fmt.Println("Creation of encoded connection failed")
		return nil, err
	}

	return encodedConn, nil

}
