package main

import (
	"fmt"
	"log"

	"github.com/IBM/sarama"
)

func main() {
	// Configure the producer
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	// Connect to Kafka brokers
	brokers := []string{"localhost:9092"}
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		log.Fatalln("Failed to start Sarama producer:", err)
	}
	defer producer.Close()

	// Your topic and message
	topic := "new_topic"
	message := "Hello, Kafka!"

	// Send the message to Kafka
	msg := &sarama.ProducerMessage{
		Topic:     topic,
		Value:     sarama.StringEncoder(message),
		Partition: -1,
	}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Println("Failed to send message:", err)
		return
	}

	fmt.Printf("Message sent successfully! Topic: %s, Partition: %d, Offset: %d\n", topic, partition, offset)
}
