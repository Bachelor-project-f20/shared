package config

import (
	"errors"
	"log"
	"os"

	"github.com/Bachelor-project-f20/eventToGo"
	etgNats "github.com/Bachelor-project-f20/eventToGo/nats"
	etgSNS "github.com/Bachelor-project-f20/eventToGo/sns"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
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
	if config.MessageBrokerType == eventToGo.SNS {
		err := setupSNS(config, result)
		if err != nil {
			return err
		}
		return nil
	}
	log.Println("setupEventEmitterAndListener DONE")
	return nil
}

func setupSNS(config *ConfigValues, result *ConfigResult) error {
	log.Println("Setting up SNS")

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{Credentials: credentials.AnonymousCredentials, Endpoint: aws.String("http://localhost:9911"), Region: aws.String("us-east-1"), DisableSSL: aws.Bool(true)},
	}))

	snsClient := sns.New(sess)

	snsHandler := etgSNS.SNSHandler{}
	if config.Events == nil {
		return errors.New("Must supply a list of events(topics) to create and subscribe to via config object")
	}

	if config.UseEmitter && config.UseListener {
		e, l, err := snsHandler.SetupEmitterAndListener(snsClient, config.Events...)
		if err != nil {
			return err
		}
		result.EventEmitter = e
		result.EventListener = l
	} else if config.UseEmitter {
		e, err := snsHandler.SetupEmitter(snsClient, config.Events...)
		if err != nil {
			log.Fatalf("Error creating event emitter: %v \n", err)
			return err
		}
		result.EventEmitter = e
	} else if config.UseListener {
		l, err := snsHandler.SetupListener(snsClient, config.Events...)
		if err != nil {
			log.Fatalf("Error creating event listener: %v \n", err)
			return err
		}
		result.EventListener = l
	}
	log.Println("SHARED: setup SNS done")
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
