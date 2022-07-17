package service

import (
	"fmt"
	"github.com/raymondgitonga/producer-service/config"
	"github.com/raymondgitonga/producer-service/internal/repositiory"
	"log"
)

type Produce struct{}

type Producer interface {
	SendMultiplicationMessage()
}

func (p Produce) SendMultiplicationMessage() {
	kafkaConfig := config.KafkaConfig{}
	calc, err := repositiory.NewVariables(float64(4), float64(6))

	if err != nil {
		log.Fatalf("Couldn't set new variables %s", err)
	}

	result := fmt.Sprintf("%f", calc.Multiply())

	err = kafkaConfig.Connect("multiply", result)

	if err != nil {
		log.Fatalf("Error sending message %s", err)
	}

}
