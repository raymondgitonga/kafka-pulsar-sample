package main

import "github.com/raymondgitonga/producer-service/internal/service"

func main() {

	for {
		producer := service.Produce{}
		producer.SendMultiplicationMessage()
	}

}
