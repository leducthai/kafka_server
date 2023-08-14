package main

import (
	"log"
	"os"
	"os/signal"
	"sync"

	"github.com/IBM/sarama"
)

func main() {
	brokers := []string{"localhost:9091", "localhost:9092"} // Replace with your Kafka broker addresses

	// Configure the consumer
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	// Create a new consumer
	consumer, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		log.Fatalln("Failed to create consumer:", err)
	}
	defer consumer.Close()

	topic := "real-time-topic" // Replace with the topic name used for producing messages
	partitions := make([]int32, 6)
	for i := 0; i < 6; i++ {
		partitions[i] = int32(i)
	}

	// Create a wait group to track the goroutines
	var wg sync.WaitGroup

	// Create a channel to handle the interrupt signal
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	// Create a channel to receive merged messages
	mergedMessages := make(chan *sarama.ConsumerMessage)

	// Create a channel to send the stop signal
	stopSignal := make(chan struct{})

	// Start processing messages
	for _, partition := range partitions {
		pc, err := consumer.ConsumePartition(topic, partition, sarama.OffsetOldest)
		if err != nil {
			log.Println("Failed to create message consumer:", err)
			return
		}
		wg.Add(1)
		go collect(pc, mergedMessages, stopSignal, &wg)
	}

	for {
		select {
		case msg, ok := <-mergedMessages:
			if !ok {
				// The mergedMessages channel is closed, exit the loop
				return
			}
			// Process the received message
			log.Printf("Received real-time message: Partition: %d, Offset: %d, Key: %s, Value: %s\n",
				msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))

		case <-signals:
			log.Println("Received interrupt signal. Stopping consumer...")
			// Close the stopSignal channel to signal goroutines to stop
			close(stopSignal)
		}
		select {
		case <-stopSignal:
			return
		default:
			continue
		}
	}
}

// Function to forward messages from a partition consumer to the output channel
func collect(c sarama.PartitionConsumer, out chan<- *sarama.ConsumerMessage, stopSignal <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case msg := <-c.Messages():
			out <- msg
		case err := <-c.Errors():
			log.Println("error: ", err)
		case <-stopSignal:
			// Stop signal received, exit the loop
			return
		}
	}
}
