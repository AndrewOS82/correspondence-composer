package main

import (
	"fmt"

	"correspondence-composer/config"
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

	// This is a placeholder. The correspondence type will be determined by incoming kafka message.
	correspondenceType := "anniversary"

	fmt.Println("Running composer process...")
	err := composer.RunProcess(correspondenceType)
	if err != nil {
		fmt.Printf("Error running composer process: %v\n", err)
	}
	fmt.Println("Composer finished successfully")
}
