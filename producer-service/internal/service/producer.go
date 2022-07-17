package service

import (
	"context"
	"fmt"
	"github.com/raymondgitonga/producer-service/config"
	"github.com/raymondgitonga/producer-service/internal/repositiory"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type Produce struct{}

type Producer interface {
	SendMultiplicationMessage()
}

func (p Produce) SendMultiplicationMessage() {
	ctx, cancel := context.WithCancel(context.Background())
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGKILL)

	kafkaConfig := config.KafkaConfig{}

	// go routine for getting signals asynchronously
	go func() {
		sig := <-signals
		fmt.Println("Got signal: ", sig)
		cancel()
	}()

	calc, err := repositiory.NewVariables(float64(9), float64(6))

	if err != nil {
		log.Fatalf("Couldn't set new variables %s", err)
	}

	result := fmt.Sprintf("%f", calc.Multiply())

	err = kafkaConfig.Connect("multiply", result, ctx)

	if err != nil {
		log.Fatalf("Error sending message %s", err)
	}

}
