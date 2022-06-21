package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"time"
)

const (
	kafkaURL = "localhost:9092"
	topic    = "topic1"
)

func newKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}

func main() {
	writer := newKafkaWriter(kafkaURL, topic)
	defer writer.Close()

	fmt.Println("start producing ... !!")

	for i := 0; ; i++ {
		key := fmt.Sprintf("Key-%d", i)
		msg := kafka.Message{
			Key:   []byte(key),
			Value: []byte(fmt.Sprint(uuid.New())),
		}
		err := writer.WriteMessages(context.Background(), msg)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("produced", key)
		}
		time.Sleep(1 * time.Second)
	}
}
