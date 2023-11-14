package main

import (
	"fmt"
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	kafkaServer := os.Getenv("KAFKA_SERVER")
	kafkaTopic := os.Getenv("KAFKA_TOPIC")
	groupID := os.Getenv("KAFKA_GROUP_ID")
	config := kafka.ConfigMap{"bootstrap.servers": kafkaServer, "group.id": groupID, "go.events.channel.enable": true}
	consumer, consumerCreateErr := kafka.NewConsumer(&config)
	if consumerCreateErr != nil {
		fmt.Println("consumer not created ", consumerCreateErr.Error())
		os.Exit(1)
	}
	subscriptionErr := consumer.Subscribe(kafkaTopic, nil)
	if subscriptionErr != nil {
		fmt.Println("Unable to subscribe to topic " + kafkaTopic + " due to error - " + subscriptionErr.Error())
		os.Exit(1)
	} else {
		fmt.Println("subscribed to topic ", kafkaTopic)
	}

	f, err := os.Create("./logs/logs.txt")

	if err != nil {
		log.Fatalf("Error in creating file")
	}

	defer f.Close()

	for {
		fmt.Println("waiting for event...")
		kafkaEvent := <-consumer.Events()
		if kafkaEvent != nil {
			switch event := kafkaEvent.(type) {
			case *kafka.Message:
				_, err := f.WriteString(fmt.Sprintf("%+v\n", string(event.Value)))
				if err != nil {
					fmt.Println("Error in writing to file ", err.Error())
				}
			case kafka.Error:
				fmt.Println("Consumer error ", event.String())
			case kafka.PartitionEOF:
				fmt.Println(kafkaEvent)
			default:
				fmt.Println(kafkaEvent)
			}
		} else {
			fmt.Println("Event was null")
		}
	}

}
