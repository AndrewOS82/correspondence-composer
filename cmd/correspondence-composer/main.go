package main

import (
	"context"
	"fmt"

	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	"correspondence-composer/config"
	"correspondence-composer/gateways/kafkaclient"
	"correspondence-composer/gateways/policyapi"
	rulesgateway "correspondence-composer/gateways/rulesengine"
	"correspondence-composer/gateways/s3client"
	"correspondence-composer/models"
	"correspondence-composer/service"
	"correspondence-composer/usecases"
	"correspondence-composer/utils/log"
)

const serviceName = "correspondence-composer"

func main() {
	// I'm introducing ctx because it's an expected parameter for the policy API client. We have some
	// nice utils at Pg that can also help us leverage context in error tracing and logging via Datadog.
	// Waiting to add those in once we have the Datadog integration set up.
	ctx, cancelRootCtx := context.WithCancel(context.Background())
	logger := log.New(log.Config{
		ServiceName: serviceName,
	})
	config := config.GetConfig(logger)

	policyAPIClient, err := policyapi.NewClientWithResponses(config.PolicyAPIBaseURL)
	if err != nil {
		logger.ErrorWithFields(err, log.Fields{
			"baseURL": config.PolicyAPIBaseURL,
			"msg":     "error setting up policy API client",
		})
		return
	}
	policyAPIGateway := policyapi.New(policyAPIClient, config.PolicyAPIAuthToken)

	rulesEngineGateway := rulesgateway.New(config.RulesEngine)

	s3 := setupS3Client(config.S3)
	storageclient := s3client.New(config.S3, s3)

	composer := service.Composer{
		DataFetcher: &usecases.DataFetcher{
			Logger:    logger,
			PolicyAPI: policyAPIGateway,
		},
		RuleExecutor: &usecases.RuleExecutor{
			RulesEngine: rulesEngineGateway,
		},
		Uploader: &usecases.Uploader{
			StorageClient: storageclient,
		},
		XMLBuilder: &usecases.XMLBuilder{},
	}

	kafka := kafkaclient.New(config.Kafka)
	// nolint
	kafka.Subscribe("correspondence.test.one", func(key string, value string) {
		handleKafkaMessage(ctx, key, value, composer)
	})

	cancelRootCtx()
}

func handleKafkaMessage(ctx context.Context, messageKey string, messageValue string, composer service.Composer) {
	fmt.Printf("Running composer process with value: [%v]\n", messageValue)
	// I don't know what the real kafka events will look like. This is a mock just to get
	// and end-to-end demo of the app working with an incoming dynamic policy number.
	newEvent := &models.KafkaEvent{
		Name:       "anniversary",
		PolicyData: &models.Policy{PolicyNumber: messageValue},
	}
	err := composer.RunProcess(ctx, newEvent)
	if err != nil {
		fmt.Printf("Composer process failed. Error: %v\n", err)
		return
	}

	fmt.Printf("Composer finished successfully. Completed message: [%v] [%v]\n", messageKey, messageValue)
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
