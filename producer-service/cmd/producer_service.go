package main

import "github.com/raymondgitonga/producer-service/internal/service"

func main() {
	count := 0
	for count < 10 {
		count++

		producer := service.Produce{}
		producer.SendMultiplicationMessage()
	}

}
