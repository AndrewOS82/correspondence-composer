package kafkaclient

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type EventConsumer func(key string, value string)

type EventPublisher interface {
	Publish(key string, value string) error
	Close()
}

type Kafka struct {
	Config *Config
}

type Config struct {
	BootstrapServer  string
	SecurityProtocol string
	GroupID          string
	SASLMechanism    string
	SASLUsername     string
	SASLPassword     string
}

func New(config Config) Kafka {
	return Kafka{Config: &config}
}

func (k *Kafka) Subscribe(topic string, consumer EventConsumer) error {
	kafkaConsumer, err := kafka.NewConsumer(convertConfigToKafkaConfig(k.Config))
	defer func() {
		_ = kafkaConsumer.Close()
	}()

	if err != nil {
		fmt.Printf("failed to create consumer: %s", err)
		return err
	}

	err = kafkaConsumer.SubscribeTopics([]string{topic}, nil)

	if err != nil {
		fmt.Printf("failed to subscribe: %s", err)
		return err
	}

	// Set up a channel for handling Ctrl-C, etc
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	fmt.Printf("%+v\n", k.Config)

	fmt.Println("entering message pump")

	run := true
	for run {
		select {
		case sig := <-sigchan:
			fmt.Printf("caught signal %v: Terminating\n", sig)
			fmt.Println("setting run=false")
			run = false
		default:
			//todo: consider making the timeout configurable
			event := kafkaConsumer.Poll(1000)

			if event == nil {
				continue
			}

			err := k.processEvent(event, consumer)

			if err != nil {
				fmt.Printf("error when processing received event: %v \n", err.Error())
				run = false
			}
		}
	}

	fmt.Println("exiting message pump")

	return nil
}

func (k *Kafka) processEvent(event kafka.Event, consumer EventConsumer) error {
	fmt.Print("processEvent \n")

	switch e := event.(type) {
	case *kafka.Message:
		fmt.Print("message \n")

		//todo: better handling of bytearrays, consider leveraging Avro, etc
		consumer(string(e.Key), string(e.Value))

	case kafka.Error:
		// Errors should generally be considered informational, the client will try to
		// automatically recover. But in this example we choose to terminate
		// the application if all brokers are down.

		fmt.Printf("error in kafka pipeline: %v\n", e.Code())

		if e.Code() == kafka.ErrAllBrokersDown {
			return e
		}

	case kafka.OffsetsCommitted:
		fmt.Print("offsetsCommitted \n")
		// You likely won't want this in production, but this event is helpful in debugging
		// commits while testing.

		if e.Error != nil {
			fmt.Printf("commit offset err: %v", e.Error)
		} else {
			for _, offset := range e.Offsets {
				fmt.Printf("offset committed: %s\n", offset)
			}
		}
	default:
		fmt.Printf("default: %v\n", e.String())

		// This captures all the other Poll() events and are effectively ignored.
	}

	return nil
}

func (k *Kafka) Publish(topic string) (EventPublisher, error) {

	fmt.Printf("%+v\n", k.Config)

	producer, err := kafka.NewProducer(convertConfigToKafkaConfig(k.Config))

	if err != nil {
		fmt.Printf("error when creating producer: %v\n", err.Error())
		return nil, err
	}

	return &kafkaProducer{
			topic:    topic,
			producer: producer,
		},
		nil
}

// -------- utility functions ----------

func convertConfigToKafkaConfig(config *Config) *kafka.ConfigMap {
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
