package main

import (
	"log"
	"os"
	"os/signal"
	"sync"

	"github.com/IBM/sarama"
)

func main() {
	brokers := []string{"localhost:9091", "localhost:9092", "localhost:9093"} // Replace with your Kafka broker addresses

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
	partitions := make([]int32, 12)
	for i := 0; i < 12; i++ {
		partitions[i] = int32(i)
	}
	// Create a new partition consumer
	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		log.Fatalln("Failed to create partition consumer:", err)
	}
	defer partitionConsumer.Close()

	// Consume messages from the partition consumer in real-time
	consumeRealTimeMessages(partitionConsumer)
}

func consumeRealTimeMessages(partitionConsumer sarama.PartitionConsumer) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	for {
		select {
		case msg := <-partitionConsumer.Messages():
			// Process the received message
			log.Printf("Received real-time message: Partition: %d, Offset: %d, Key: %s, Value: %s\n",
				msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))

		case err := <-partitionConsumer.Errors():
			log.Println("Error while consuming message:", err)

		case <-signals:
			log.Println("Received interrupt signal. Stopping consumer...")
			return
		}
	}
}

func main() {
	config := sarama.NewConfig()
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatalln("Failed to start Sarama consumer:", err)
	}
	defer consumer.Close()

	// Manually assign partitions to the consumer
	topic := "my_topic"
	partitions := []int32{0, 1, 2} // Specify the partitions you want to consume from

	var assignedPartitions []sarama.PartitionConsumer
	for _, partition := range partitions {
		pc, err := consumer.ConsumePartition(topic, partition, sarama.OffsetOldest)
		if err != nil {
			log.Println("Failed to create message consumer:", err)
			return
		}
		assignedPartitions = append(assignedPartitions, pc)
	}

	for {
		select {
		case msg := <-merged(assignedPartitions...):
			log.Printf("Received message: %s\n", string(msg.Value))
			// Process the message
		case <-ctx.Done():
			return
		}
	}
}

// Merging multiple channel streams into one
func merged(cs ...<-chan *sarama.ConsumerMessage) <-chan *sarama.ConsumerMessage {
	var wg sync.WaitGroup
	out := make(chan *sarama.ConsumerMessage)

	// Function to forward messages from a channel to the output channel
	collect := func(c <-chan *sarama.ConsumerMessage) {
		for msg := range c {
			out <- msg
		}
		wg.Done()
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go collect(c)
	}

	// Wait for all channels to be closed before closing the output channel
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
