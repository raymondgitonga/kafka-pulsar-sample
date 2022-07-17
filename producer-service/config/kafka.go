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

	defer func() {
		err := writer.Close()
		if err != nil {
			fmt.Println("Error closing producer: ", err)
			return
		}
		fmt.Println("Producer closed")
	}()

	err := writer.WriteMessages(
		ctx,
		kafka.Message{
			Key:   []byte("Key-A"),
			Value: []byte(message),
		},
	)

	if err != nil {
		return errors.New("Context cancelled called: " + err.Error())
	}

	return err
}
