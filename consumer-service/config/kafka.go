package config

import (
	"github.com/segmentio/kafka-go"
)

type KafkaConfig struct{}

func (c KafkaConfig) Connect(topic string) *kafka.Reader {
	config := kafka.ReaderConfig{
		Brokers:     []string{"localhost:29092"},
		GroupID:     "",
		GroupTopics: nil,
		Topic:       topic,
		Partition:   0,
		MaxBytes:    10,
	}

	reader := kafka.NewReader(config)

	return reader
}
