package config

import (
	"context"
	"github.com/segmentio/kafka-go"
)

type KafkaConfig struct{}

func (c KafkaConfig) Connect(topic string, message string) error {
	writer := &kafka.Writer{
		Addr:                   kafka.TCP("localhost:9092"),
		Topic:                  topic,
		AllowAutoTopicCreation: true,
	}

	err := writer.WriteMessages(
		context.Background(),
		kafka.Message{
			Key:   []byte("Key-A"),
			Value: []byte(message),
		},
	)

	return err
}
