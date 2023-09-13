package main

import (
	"context"

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
	// TODO: We have some nice utils at Pg that can help us leverage context in error tracing
	// and logging via Datadog. Waiting to add those in once we have a Datadog integration.
	ctx, cancelRootCtx := context.WithCancel(context.Background())
	logger := log.New(log.Config{
		ServiceName: serviceName,
	})
	config := config.GetConfig(logger)

	logger.SetLevel(config.LogLevel)

	policyAPIClient, err := policyapi.NewClientWithResponses(config.PolicyAPIBaseURL)
	if err != nil {
		logger.ErrorWithFields(err, log.Fields{
			"baseURL": config.PolicyAPIBaseURL,
			"msg":     "error setting up policy API client",
		})
		return
	}
	mockPolicyData, err := policyapi.GetMockPolicyData(config.PolicyDataSampleFile)
	if err != nil {
		logger.ErrorWithFields(err, log.Fields{
			"msg": "error loading mock policy data",
		})
	}
	policyAPIGateway := policyapi.New(policyAPIClient, config.PolicyAPIAuthToken, mockPolicyData)

	rulesEngineGateway := rulesgateway.New(config.RulesEngine, logger)
	rulesConfig, err := usecases.GetRulesConfig(config.RulesConfigFile)
	if err != nil {
		logger.ErrorWithFields(err, log.Fields{
			"configFile": config.RulesConfigFile,
			"msg":        "error setting up rules config",
		})
		return
	}

	s3 := setupS3Client(logger, config.S3)
	storageclient := s3client.New(config.S3, s3)

	composer := service.Composer{
		Logger: logger,
		DataFetcher: &usecases.DataFetcher{
			Logger:    logger,
			PolicyAPI: policyAPIGateway,
		},
		RuleExecutor: &usecases.RuleExecutor{
			Logger:      logger,
			RulesConfig: rulesConfig,
			RulesEngine: rulesEngineGateway,
		},
		Uploader: &usecases.Uploader{
			StorageClient: storageclient,
		},
		XMLBuilder: &usecases.XMLBuilder{},
	}

	kafka := kafkaclient.New(config.Kafka, logger)
	// nolint
	kafka.Subscribe(config.IncomingKafkaTopic, func(key string, value string) {
		handleKafkaMessage(ctx, logger, key, value, composer)
	})

	cancelRootCtx()
}

func handleKafkaMessage(ctx context.Context, logger log.Logger, messageKey string, messageValue string, composer service.Composer) {
	logger.InfoWithFields("Running composer process", log.Fields{
		"key":   messageKey,
		"value": messageValue,
	})

	// I don't know what the real kafka events will look like. This is a mock just to get
	// and end-to-end demo of the app working with an incoming dynamic policy number.
	newEvent := &models.KafkaEvent{
		Name:       "anniversary",
		PolicyData: &models.Policy{PolicyNumber: messageValue},
	}
	err := composer.RunProcess(ctx, newEvent)
	if err != nil {
		logger.ErrorWithFields(err, log.Fields{
			"msg": "composer process failed",
		})
		return
	}
}

func setupS3Client(logger log.Logger, config s3client.Config) *s3.Client {
	ctx := context.Background()
	// TO DO: set up dynamic config
	cfg, err := awsconfig.LoadDefaultConfig(ctx, awsconfig.WithRegion(config.AWSRegion))
	if err != nil {
		logger.ErrorWithFields(err, log.Fields{
			"msg": "could not load AWS config",
		})
	}

	return s3.NewFromConfig(cfg)
}
