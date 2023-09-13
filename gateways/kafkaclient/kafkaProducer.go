package kafkaclient

import (
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"

	"correspondence-composer/utils/log"
)

type kafkaProducer struct {
	topic    string
	producer *kafka.Producer
	logger   log.Logger
}

func (kp *kafkaProducer) Publish(key string, value string) error {
	kp.logger.Debugln("publishing event")

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
		kp.logger.ErrorWithFields(err, log.Fields{
			"msg":   "error when producing kafka message",
			"key":   key,
			"value": value,
			"topic": &kp.topic,
		})

		return err
	}

	return nil
}

func (kp *kafkaProducer) Close() {
	kp.producer.Close()
}
