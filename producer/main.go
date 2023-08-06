package main

import (
	"fmt"
	"log"
	"time"

	"github.com/IBM/sarama"
)

func main() {
	brokers := []string{"localhost:9092", "localhost:9091"} // Replace with your Kafka broker addresses

	// Configure the producer
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	// Create a new producer
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		log.Fatalln("Failed to create producer:", err)
	}
	defer producer.Close()

	topic := "real-time-topic" // Replace with your desired topic name

	// Send real-time messages every 1 second
	for i := 1; i <= 10; i++ {
		message := fmt.Sprintf("Real-Time Message %d", i)

		// Create a new message
		msg := &sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.StringEncoder(message),
		}

		// Send the message to Kafka
		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			log.Println("Failed to send message:", err)
		} else {
			fmt.Printf("Message sent successfully! Topic: %s, Partition: %d, Offset: %d\n", topic, partition, offset)
		}

		time.Sleep(1 * time.Second) // Wait for 1 second before sending the next message
	}
}
