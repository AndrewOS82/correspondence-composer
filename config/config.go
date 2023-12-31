package config

import (
	"os"

	"correspondence-composer/gateways/kafkaclient"
	"correspondence-composer/gateways/rulesengine"
	"correspondence-composer/gateways/s3client"
	"correspondence-composer/utils/log"
)

type Config struct {
	Env                  string
	Kafka                kafkaclient.Config
	PolicyAPIAuthToken   string
	PolicyAPIBaseURL     string
	PolicyDataSampleFile string
	RulesConfigFile      string
	IncomingKafkaTopic   string
	RulesEngine          rulesengine.Config
	S3                   s3client.Config
	LogLevel             string
}

func GetConfig(logger log.Logger) Config {
	LoadEnvFile(logger)

	env := getEnvOrDefault("ENVIRONMENT", "development")

	return newConfig(env)
}

func newConfig(env string) Config {
	return Config{
		Env: env,
		S3: s3client.Config{
			AWSRegion: os.Getenv("AWS_REGION"),
			AWSBucket: os.Getenv("AWS_BUCKET"),
		},
		RulesEngine: rulesengine.Config{
			Username:        os.Getenv("RULES_ENGINE_USERNAME"),
			Password:        os.Getenv("RULES_ENGINE_PASSWORD"),
			AuthEndpoint:    os.Getenv("RULES_ENGINE_AUTH_ENDPOINT"),
			ClientCode:      os.Getenv("RULES_ENGINE_CLIENT_CODE"),
			ExecuteEndpoint: os.Getenv("RULES_ENGINE_EXECUTE_ENDPOINT"),
		},
		Kafka: kafkaclient.Config{
			BootstrapServer:  os.Getenv("KAFKA_BOOTSTRAP_SERVER"),
			SecurityProtocol: os.Getenv("KAFKA_SECURITY_PROTOCOL"),
			GroupID:          os.Getenv("KAFKA_GROUP_ID"),
			SASLMechanism:    os.Getenv("KAFKA_SASL_MECHANISM"),
			SASLUsername:     os.Getenv("KAFKA_SASL_USERNAME"),
			SASLPassword:     os.Getenv("KAFKA_SASL_PASSWORD"),
		},
		PolicyDataSampleFile: getEnvOrDefault("POLICY_DATA_SAMPLE_FILE", "./gateways/policyapi/sample_policy_data.json"),
		PolicyAPIAuthToken:   os.Getenv("POLICY_API_TOKEN"),
		PolicyAPIBaseURL:     os.Getenv("ENTERPRISE_API_BASE_URL"),
		RulesConfigFile:      getEnvOrDefault("BUSINESS_RULES_CONFIG_FILE", "./config/business_rules.json"),
		IncomingKafkaTopic:   os.Getenv("INCOMING_KAFKA_TOPIC"),
		LogLevel:             os.Getenv("LOG_LEVEL"),
	}
}
