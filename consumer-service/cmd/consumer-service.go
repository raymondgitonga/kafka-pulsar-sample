package main

import (
	"context"
	"github.com/raymondgitonga/consumer-service/internal"
)

func main() {
	consumer := internal.Consume{}
	consumer.ReadMultiplicationMessage(context.Background())
}
