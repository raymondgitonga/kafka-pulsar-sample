package config

import (
	"context"
	"errors"
	"fmt"
	"github.com/segmentio/kafka-go"
)

type KafkaConfig struct{}

func (c KafkaConfig) Connect(topic string, message string, ctx context.Context) error {
	writer := &kafka.Writer{
		Addr:                   kafka.TCP("localhost:9092"),
		Topic:                  topic,
		AllowAutoTopicCreation: true,
	}

	err := writer.WriteMessages(
		ctx,
		kafka.Message{
			Key:   []byte("Key-A"),
			Value: []byte(message),
		},
	)

	if err == context.Canceled {
		closeErr := writer.Close()
		if closeErr != nil {
			fmt.Printf("Failed to close writer: %s", closeErr)
		}
		return errors.New("Context cancelled called: " + err.Error())
	}

	return err
}
