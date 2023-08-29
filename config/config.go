package config

import (
	"os"

	"correspondence-composer/gateways/kafkaclient"
	"correspondence-composer/gateways/rulesengine"
	"correspondence-composer/utils/log"
)

type Config struct {
	Env         string
	RulesEngine rulesengine.Config
	Kafka       kafkaclient.Config
}

func GetConfig(logger log.Logger) Config {
	LoadEnvFile(logger)

	env := getEnvOrDefault("ENVIRONMENT", "development")

	return newConfig(env)
}

func newConfig(env string) Config {
	return Config{
		Env: env,
		RulesEngine: rulesengine.Config{
			Username:        os.Getenv("RULES_ENGINE_USERNAME"),
			Password:        os.Getenv("RULES_ENGINE_PASSWORD"),
			AuthEndpoint:    os.Getenv("RULES_ENGINE_AUTH_ENDPOINT"),
			AuthClientCode:  os.Getenv("RULES_ENGINE_AUTH_CLIENT_CODE"),
			ExecuteEndpoint: os.Getenv("RULES_ENGINE_EXECUTE_ENDPOINT"),
			ExecuteClient:   os.Getenv("RULES_ENGINE_EXECUTE_CLIENT"),
		},
		Kafka: kafkaclient.Config{
			BootstrapServer:  os.Getenv("KAFKA_BOOTSTRAP_SERVER"),
			SecurityProtocol: os.Getenv("KAFKA_SECURITY_PROTOCOL"),
			GroupID:          os.Getenv("KAFKA_GROUP_ID"),
			SASLMechanism:    os.Getenv("KAFKA_SASL_MECHANISM"),
			SASLUsername:     os.Getenv("KAFKA_SASL_USERNAME"),
			SASLPassword:     os.Getenv("KAFKA_SASL_PASSWORD"),
		},
	}
}
