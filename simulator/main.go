package main

import (
	"fmt"
	"github.com/joho/godotenv"
	kafka2 "github.com/vitor9/fs-fc-code_delivery/application/kafka"
	"github.com/vitor9/fs-fc-code_delivery/infra/kafka"
	//kafka2 "github.com/vitor9/fs-fc-code_delivery/application/kafka"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}
func main() {
	msgChan := make(chan *ckafka.Message)
	consumer := kafka.NewKafkaConsumer(msgChan)
	go consumer.Consume()

	for msg := range msgChan {
		fmt.Println(string(msg.Value))
		go kafka2.Produce(msg)
	}

}
