package main

import (
	"fmt"
	"log"
	"testing"
	/* "os"
	"os/signal"
	"syscall" */

	"github.com/Shopify/sarama"

	"kafka"
)

func TestKafka(t *testing.T) {
	topic := "test"

	// producer
	producer, err := kafka.ConnectProducer([]string{"localhost:9093"})
	if err != nil {
		log.Fatal("Error connecting Kafka")
	}
	defer producer.Close()

	message := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder("test"),
	}
	partition, offset, err := producer.SendMessage(message)
	if err != nil {
		log.Fatal("Error connect producer")
	}
	fmt.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", topic, partition, offset)

	// consumer
	worker, err := kafka.ConnectConsumer([]string{"localhost:9093"})
	if err != nil {
		log.Fatal("Error connect consumer")
	}
	consumer, err := worker.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		log.Fatal("Error connect consumer")
	}

	// sigchan := make(chan os.Signal, 1)
	// signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	done := make(chan int)

	go func() {
		for {
			select {
			case err := <-consumer.Errors():
				fmt.Println(err)
			case msg := <-consumer.Messages():
				fmt.Printf("topic: %s, value: %s\n", string(msg.Topic), string(msg.Value))
				/* case <-sigchan:
				fmt.Println("Interrupt is detected")
				doneCh <- 1 */
			}
		}
	}()

	<-done

	if err := worker.Close(); err != nil {
		log.Fatal("Error")
	}
}
