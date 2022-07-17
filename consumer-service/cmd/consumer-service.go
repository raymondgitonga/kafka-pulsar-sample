package main

import "github.com/raymondgitonga/consumer-service/internal"

func main() {

	for {
		consumer := internal.Consume{}
		consumer.ReadMultiplicationMessage()
	}

}
