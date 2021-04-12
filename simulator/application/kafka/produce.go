package kafka

import (
	"encoding/json"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	route2 "github.com/vitor9/fs-fc-code_delivery/application/route"
	"github.com/vitor9/fs-fc-code_delivery/infra/kafka"
	"log"
	"os"
	"time"
)

func Produce(msg *ckafka.Message) {
	producer := kafka.NewKafkaProducer()
	route := route2.NewRoute()

	json.Unmarshal(msg.Value, &route)
	route.LoadPositions()

	positions, err := route.ExportJsonPositions()
	if err != nil {
		log.Println(err.Error())
	}

	for _, p := range positions {
		kafka.Publish(p, os.Getenv("KafkaProduceTopic"), producer)
		// Para nao ficar mandando muito rapido as transacoes, colocamos para mandar a cada 500 milisegundos
		time.Sleep(time.Millisecond * 500)
	}
}