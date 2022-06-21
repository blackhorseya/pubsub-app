package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/google/uuid"
	"time"
)

const (
	kafkaURL = "localhost:9092"
)

func main() {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": kafkaURL})
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	go func() {
		for e := range producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %s\n", ev.Key)
				}
			}
		}
	}()

	topic := "topic1"
	for i := 0; ; i++ {
		key := fmt.Sprintf("Key-%d", i)
		msg := &kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Key:            []byte(key),
			Value:          []byte(fmt.Sprint(uuid.New())),
		}
		_ = producer.Produce(msg, nil)
		time.Sleep(1 * time.Second)
	}
}
