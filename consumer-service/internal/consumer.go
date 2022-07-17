package internal

import (
	"context"
	"fmt"
	"github.com/raymondgitonga/consumer-service/config"
)

type Consume struct{}

type Consumer interface {
	ReadMultiplicationMessage()
}

func (c *Consume) ReadMultiplicationMessage(ctx context.Context) {
	kafkaConfig := config.KafkaConfig{}
	msgChan := make(chan config.Received)
	go kafkaConfig.Connect("multiply", ctx, msgChan)

	for val := range msgChan {
		fmt.Println(val.Message)
	}
}
