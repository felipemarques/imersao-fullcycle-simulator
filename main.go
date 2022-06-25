package main

import (
	"fmt"
	"log"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	kafka2 "github.com/felipemarques/imersao-fullcycle-simulator/application/kafka"
	"github.com/felipemarques/imersao-fullcycle-simulator/infra/kafka"
	"github.com/joho/godotenv"
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

	// producer := kafka.NewKafkaProducer()
	// kafka.Publish("ola", "readtest", producer)
	// for {
	// 	_ = 1
	// }

	// router := route2.Route{
	// 	ID:       "1",
	// 	ClientID: "1",
	// }
	// router.LoadPositions()
	// stringJson, _ := router.ExportJsonPositions()
	// fmt.Println(stringJson[1])
}
