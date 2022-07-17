package config

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
)

type KafkaConfig struct{}

type Received struct {
	Message string
	Offset  int64
}

func (c KafkaConfig) Connect(topic string, ctx context.Context, msgChan chan Received) {
	defer close(msgChan)
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
	err := reader.SetOffset(1)

	if err != nil {
		log.Printf("Failed to set offset %s", err)
	}

	for {
		message, err := reader.ReadMessage(ctx)
		if err != nil {
			fmt.Println("Error reading message ", err)
			break
		}

		fmt.Println(string(message.Value))
		msgChan <- Received{
			Message: string(message.Value),
			Offset:  message.Offset,
		}
	}
	if err := reader.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}
