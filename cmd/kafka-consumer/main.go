package main

import (
	"encoding/json"
	"flag"
	"fmt"

	"github.com/Shopify/sarama"
	"github.com/ozoncp/ocp-classroom-api/internal/producer"
)

var broker = flag.String("broker", producer.KafkaBroker, "Kafka Apache broker endpoint to connect")
var topic = flag.String("topic", producer.KafkaTopic, "Kafka Apache topic to consume")

func main() {

	flag.Parse()

	consumer, err := sarama.NewConsumer([]string{*broker}, nil)
	if err != nil {
		fmt.Println("Could not create consumer: ", err)
	}

	fmt.Println("Consumer created")

	subscribe(*topic, consumer)

	var cmd string
	fmt.Println("Press enter to exit")
	_, _ = fmt.Scanln(&cmd)
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
