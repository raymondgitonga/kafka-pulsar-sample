package config

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type KafkaConfig struct{}

type Received struct {
	Message string
	Offset  int64
}

func (c KafkaConfig) Connect(topic string, ctx context.Context, msgChan chan Received) {
	defer close(msgChan)
	ctx, cancel := context.WithCancel(ctx)
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGKILL)

	// go routine for getting signals asynchronously
	go func() {
		sig := <-signals
		fmt.Println("Got signal: ", sig)
		cancel()
	}()

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
			if err == context.Canceled {
				fmt.Println("Signal interrupt error ", err)
				reader.Close()
				break
			}
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
