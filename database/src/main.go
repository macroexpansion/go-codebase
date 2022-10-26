package main

import (
	"fmt"
	"log"

	"kafka"
)

func main() {
	fmt.Println("test")

	producer, err := kafka.ConnectProducer([]string{"localhost:9092"})
	_ = producer
	if err != nil {
		log.Fatal("Error connecting Kafka")
	}
	defer producer.Close()
}
