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

func (c *Consume) ReadMultiplicationMessage() {
	kafkaConfig := config.KafkaConfig{}

	reader := kafkaConfig.Connect("multiply")

	message, err := reader.ReadMessage(context.Background())

	if err != nil {
		fmt.Printf("failed to read message: %s", err)
		return
	}

	fmt.Println("Message is:...", message)

	//err = reader.Close()
	//
	//if err != nil {
	//	fmt.Printf("failed to close reader: %s", err)
	//	return
	//}
}
