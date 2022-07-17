package internal

import (
	"context"
	"github.com/raymondgitonga/consumer-service/config"
)

type Consume struct{}

type Consumer interface {
	ReadMultiplicationMessage()
}

func (c *Consume) ReadMultiplicationMessage(ctx context.Context) {
	kafkaConfig := config.KafkaConfig{}

	kafkaConfig.Connect("multiply", ctx)

}
