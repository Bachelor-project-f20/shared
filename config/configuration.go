package config

import (
	"encoding/json"
	"log"
	"os"

	"github.com/Bachelor-project-f20/eventToGo"
	"github.com/Bachelor-project-f20/go-outbox"
)

var (
	DBTypeDefault                  = outbox.MySQL
	DBConnectionDefault            = "root:root@/root?charset=utf8&parseTime=True&loc=Local"
	MessageBrokerTypeDefault       = eventToGo.NATS
	MessageBrokerDefaultConnection = "localhost:4222"
	ExchangeDefault                = "user"
	QueueTypeDefault               = "queue"
)

type ConfigValues struct {
	DatabaseType            outbox.DbType
	DatabaseConnection      string
	MessageBrokerType       eventToGo.BrokerType
	MessageBrokerConnection string
	Exchange                string
	QueueType               string
	UseEmitter              bool
	UseListener             bool
	UseOutbox               bool
	OutboxModels            []interface{}
}

type ConfigResult struct {
	EventEmitter  eventToGo.EventEmitter
	EventListener eventToGo.EventListener
	Outbox        outbox.Outbox
}

func ConfigService(configFilePath string, defaultValues ConfigValues) (ConfigResult, error) {
	config := setupDefaulValues(defaultValues)
	configResult := ConfigResult{}

	err := config.extractConfiguration(configFilePath)
	if err != nil {
		log.Panicln("Failed to extractConfigurations: ", err)
		return ConfigResult{}, err
	}

	err = setupEventEmitterAndListener(&config, &configResult)
	if err != nil {
		log.Panicln("Failed to setup message broker: ", err)
		return ConfigResult{}, err
	}

	err = setupOutbox(&config, &configResult)
	if err != nil {
		log.Panicln("Failed to setup outbox: ", err)
		return ConfigResult{}, err
	}

	return configResult, nil
}

func setupDefaulValues(defaultValues ConfigValues) ConfigValues {
	config := defaultValues
	config.DatabaseType = defaultValues.DatabaseType
	if defaultValues.DatabaseConnection == "" {
		config.DatabaseConnection = DBConnectionDefault
	}
	config.MessageBrokerType = defaultValues.MessageBrokerType
	if defaultValues.MessageBrokerConnection == "" {
		config.MessageBrokerConnection = MessageBrokerDefaultConnection
	}
	if defaultValues.Exchange == "" {
		config.Exchange = ExchangeDefault
	}
	if defaultValues.QueueType == "" {
		config.QueueType = QueueTypeDefault
	}
	config.UseEmitter = defaultValues.UseEmitter
	config.UseListener = defaultValues.UseListener
	config.UseOutbox = defaultValues.UseOutbox
	return config
}

func (config *ConfigValues) extractConfiguration(filename string) error {

	config.extractFromFile(filename)

	extractFromEnvForDb(config)
	extractFromEnvForQueue(config)

	return nil
}

func (config *ConfigValues) extractFromFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Println("Config file not found")
	}
	json.NewDecoder(file).Decode(config)
}
