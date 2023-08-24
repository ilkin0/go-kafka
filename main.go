package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"log"
)

func main() {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"client.id":         "client.test.id",
		"acks":              "all",
	})

	if err != nil {
		fmt.Printf("Failed to create Producer: %s\n", err)
	}

	delivery_chan := make(chan kafka.Event, 1000)
	topic := "GoKafka"

	err = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte("FOO_TEST")},
		delivery_chan)

	if err != nil {
		log.Fatal(err)
	}
}
