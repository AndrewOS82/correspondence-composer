package kafkaclient

import (
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type kafkaProducer struct {
	topic    string
	producer *kafka.Producer
}

func (kp *kafkaProducer) Publish(key string, value string) error {
	err := kp.producer.Produce(
		&kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &kp.topic,
				Partition: kafka.PartitionAny,
			},
			Key:           []byte(key),
			Value:         []byte(value),
			Timestamp:     time.Time{},
			TimestampType: 0,
			Opaque:        nil,
			Headers:       []kafka.Header{},
		},
		nil)

	if err != nil {
		fmt.Printf("error when publishing event: %v", err.Error())
		return err
	}

	return nil
}

func (kp *kafkaProducer) Close() {
	kp.producer.Close()
}
