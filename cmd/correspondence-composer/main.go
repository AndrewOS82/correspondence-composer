package main

import (
	"context"
	"fmt"

	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	"correspondence-composer/config"
	"correspondence-composer/gateways/kafkaclient"
	rulesgateway "correspondence-composer/gateways/rulesengine"
	"correspondence-composer/gateways/s3client"
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
	s3 := setupS3Client(config.S3)
	storageclient := s3client.New(config.S3, s3)
	composer := service.Composer{
		DataFetcher: &usecases.DataFetcher{},
		RuleExecutor: &usecases.RuleExecutor{
			RulesEngine: rulesEngineGateway,
		},
		Uploader: &usecases.Uploader{
			StorageClient: storageclient,
		},
		// XmlGenerator: &usecases.XmlGenerator{},
	}

	kafka := kafkaclient.New(config.Kafka)
	// nolint
	kafka.Subscribe("correspondence.test.one", func(key string, value string) {
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

func setupS3Client(config s3client.Config) *s3.Client {
	ctx := context.Background()
	// TO DO: set up dynamic config
	cfg, err := awsconfig.LoadDefaultConfig(ctx, awsconfig.WithRegion(config.AWSRegion))
	if err != nil {
		fmt.Printf("Could not load AWS config")
	}

	return s3.NewFromConfig(cfg)
}
