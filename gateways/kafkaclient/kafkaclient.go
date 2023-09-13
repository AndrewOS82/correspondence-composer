package kafkaclient

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"

	"correspondence-composer/utils/log"
)

type EventConsumer func(key string, value string)

type EventPublisher interface {
	Publish(key string, value string) error
	Close()
}

type Kafka struct {
	config Config
	logger log.Logger
}

type Config struct {
	BootstrapServer  string
	SecurityProtocol string
	GroupID          string
	SASLMechanism    string
	SASLUsername     string
	SASLPassword     string
}

func New(config Config, logger log.Logger) Kafka {
	return Kafka{
		config: config,
		logger: logger,
	}
}

func (k *Kafka) Subscribe(topic string, consumer EventConsumer) error {
	kafkaConsumer, err := kafka.NewConsumer(convertConfigToKafkaConfig(k.config))
	defer func() {
		_ = kafkaConsumer.Close()
	}()

	if err != nil {
		k.logger.ErrorWithFields(err, log.Fields{
			"msg": "failed to create consumer",
		})

		return err
	}

	err = kafkaConsumer.SubscribeTopics([]string{topic}, nil)

	if err != nil {
		k.logger.ErrorWithFields(err, log.Fields{
			"msg":   "failed to subscribe",
			"topic": topic,
		})
		return err
	}

	// Set up a channel for handling Ctrl-C, etc
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	k.logger.Debug("entering kafka select loop")

	run := true
	for run {
		select {
		case sig := <-sigchan:
			k.logger.Infof("caught signal %v: Terminating\n", sig)
			run = false
		default:
			//todo: consider making the timeout configurable
			event := kafkaConsumer.Poll(1000)

			if event == nil {
				continue
			}

			err := k.processEvent(event, consumer)

			if err != nil {
				k.logger.ErrorWithFields(err, log.Fields{
					"msg": "error when processing received event",
				})

				run = false
			}
		}
	}

	k.logger.Debug("exiting kafka select loop")

	return nil
}

func (k *Kafka) processEvent(event kafka.Event, consumer EventConsumer) error {
	k.logger.Debugln("processEvent")

	switch e := event.(type) {
	case *kafka.Message:
		k.logger.Debugln("message")

		//todo: better handling of bytearrays, consider leveraging Avro, etc
		consumer(string(e.Key), string(e.Value))

	case kafka.Error:
		// Errors should generally be considered informational, the client will try to
		// automatically recover. But in this example we choose to terminate
		// the application if all brokers are down.

		if e.Code() == kafka.ErrAllBrokersDown {
			k.logger.ErrorWithFields(e, log.Fields{
				"msg": "Cannot connect to kafka ",
			})

			return e
		}

		k.logger.InfoWithFields(e.String(), log.Fields{
			"msg": "recoverable kafka runtime error received, continuing processing",
		})

	case kafka.OffsetsCommitted:
		k.logger.Debugln("offsetsCommitted")

		// You likely won't want this in production, but this event is helpful in debugging
		// commits while testing.

		if e.Error != nil {
			k.logger.ErrorWithFields(e.Error, log.Fields{
				"msg": "error when committing offset to kafka",
			})

			return e.Error

		}

		for _, offset := range e.Offsets {

			k.logger.InfoWithFields("offset committed", log.Fields{
				"offset": offset,
			})
		}

	default:
		k.logger.Debugf("default: %v\n", e.String())

		// This captures all the other Poll() events and are effectively ignored.
	}

	return nil
}

func (k *Kafka) Publish(topic string) (EventPublisher, error) {
	producer, err := kafka.NewProducer(convertConfigToKafkaConfig(k.config))

	if err != nil {
		k.logger.ErrorWithFields(err, log.Fields{
			"msg": "error when creating producer",
		})

		return nil, err
	}

	return &kafkaProducer{
			topic:    topic,
			producer: producer,
		},
		nil
}

// -------- utility functions ----------

func convertConfigToKafkaConfig(config Config) *kafka.ConfigMap {
	kafkaConfig := kafka.ConfigMap{
		"bootstrap.servers": config.BootstrapServer,
		"security.protocol": config.SecurityProtocol,
		"group.id":          config.GroupID,
	}

	// nolint
	if len(config.SASLMechanism) > 0 {
		kafkaConfig.SetKey("sasl.mechanisms", config.SASLMechanism)
		kafkaConfig.SetKey("sasl.username", config.SASLUsername)
		kafkaConfig.SetKey("sasl.password", config.SASLPassword)
	}

	return &kafkaConfig
}
