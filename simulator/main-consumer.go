package main

import (
	"github.com/joho/godotenv"
	"github.com/vitor9/fs-fc-code_delivery/infra/kafka"

	"log"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {
	// Na hora de rodar o metodo main, antes do Kafka conseguir se conectar e publicar a msg
	// a funcao main acaba. Por isso, fizemos uma gambiarra
	// de loop infinito para deixar o go sempre de pe
	producer := kafka.NewKafkaProducer()

	kafka.Publish("Ola","readtest", producer)
	for  {
		_ = 1
	}

	//route := route2.Route{
	//	ID:        "1",
	//	ClientID:  "1",
	//}
	//route.LoadPositions()
	//stringjson, _ := route.ExportJsonPositions()
	//fmt.Println(stringjson[0])
}
