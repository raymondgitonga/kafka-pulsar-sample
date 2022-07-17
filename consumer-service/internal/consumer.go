package internal

import (
	"context"
	"fmt"
	"github.com/raymondgitonga/consumer-service/config"
	"os"
	"os/signal"
	"syscall"
)

type Consume struct{}

type Consumer interface {
	ReadMultiplicationMessage()
}

func (c *Consume) ReadMultiplicationMessage(ctx context.Context) {
	kafkaConfig := config.KafkaConfig{}
	msgChan := make(chan config.Received)
	ctx, cancel := context.WithCancel(ctx)
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGKILL)

	// go routine for getting signals asynchronously
	go func() {
		sig := <-signals
		fmt.Println("Got signal: ", sig)
		cancel()
	}()

	go kafkaConfig.Connect("multiply", ctx, msgChan)

	for val := range msgChan {
		fmt.Println(val.Message)
	}
}
