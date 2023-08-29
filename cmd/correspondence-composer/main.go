package main

import (
	"fmt"

	"correspondence-composer/config"
	"correspondence-composer/gateways/kafkaclient"
	rulesgateway "correspondence-composer/gateways/rulesengine"
	"correspondence-composer/service"
	"correspondence-composer/usecases"
	"correspondence-composer/utils/log"
)

const serviceName = "correspondence-composer"

func main() {
	logger := log.New(log.Config{
		ServiceName: serviceName,
	})
	config := config.GetConfig(logger)

	rulesEngineGateway := rulesgateway.New(config.RulesEngine)
	composer := service.Composer{
		DataFetcher: &usecases.DataFetcher{},
		RuleExecutor: &usecases.RuleExecutor{
			RulesEngine: rulesEngineGateway,
		},
		// XmlGenerator: &usecases.XmlGenerator{},
	}

	kafka := kafkaclient.New(config.Kafka)
	// nolint
	kafka.Subscribe("correspondence.test.one", func(key string, value string){
		handleKafkaMessage(key, value, composer)
	})
}

func handleKafkaMessage(key string, value string, composer service.Composer) {
	fmt.Printf("Running composer process with value: [%v]\n", value)
	err := composer.RunProcess(value)
	if err != nil {
		fmt.Printf("Error running composer process: %v\n", err)
	}
	fmt.Printf("Composer finished successfully. Completed message: [%v] [%v]\n", key, value)
}
