package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/Shopify/sarama"
	"github.com/ozoncp/ocp-classroom-api/internal/producer"
)

func main() {

	consumer, err := sarama.NewConsumer([]string{producer.KafkaBroker}, nil)
	if err != nil {
		fmt.Println("Could not create consumer: ", err)
	}

	fmt.Println("Consumer created")

	subscribe(producer.KafkaTopic, consumer)

	for {
		time.Sleep(time.Second)
	}
}

func subscribe(topic string, consumer sarama.Consumer) {

	partitionList, err := consumer.Partitions(topic) //get all partitions on the given topic
	if err != nil {
		fmt.Println("Error retrieving partitionList ", err)
	}

	initialOffset := sarama.OffsetOldest //get offset for the oldest message on the topic
	for _, partition := range partitionList {
		pc, _ := consumer.ConsumePartition(topic, partition, initialOffset)
		go func(pc sarama.PartitionConsumer) {
			for message := range pc.Messages() {
				messageReceived(message)
			}
		}(pc)
	}
}

func messageReceived(message *sarama.ConsumerMessage) {

	var classroomEvent producer.ClassroomEvent

	bytes := sarama.StringEncoder(message.Value)

	err := json.Unmarshal([]byte(bytes), &classroomEvent)
	if err != nil {
		fmt.Println("Could not unmarshal message")
	}

	fmt.Println("ClassroomEvent:", classroomEvent)
}
