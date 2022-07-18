package service

import (
	"context"
	"fmt"
	"github.com/raymondgitonga/producer-service/config"
	"github.com/raymondgitonga/producer-service/internal/repositiory"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
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

	calc, err := repositiory.NewVariables(generateRandomNumber(), generateRandomNumber())

	if err != nil {
		log.Printf("Couldn't set new variables %s", err)
		return
	}

	result := fmt.Sprintf("%d", calc.Multiply())

	err = kafkaConfig.Connect("times", result, ctx)

	if err != nil {
		log.Fatalf("Error sending message %s", err)
	}
}

func generateRandomNumber() int {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 50

	randomNum := rand.Intn(max-min+1) + min
	return randomNum
}
