package main

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"listener/event"
	"log"
	"math"
	"os"
	"time"
)

func main() {
	// try to connect to rabbit mq
	rabbitConn, err := connect()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	defer rabbitConn.Close()
	// start listening for incoming messages/jobs
	log.Println("listening for and consuming RabbitMQ messages")
	// create a consumer

	consumer, err := event.NewConsumer(rabbitConn)
	if err != nil {
		panic(err)
	}

	// watch the queue && consume events
	err = consumer.Listen([]string{"log.INFO", "log.WARNING", "log.ERROR"})
	if err != nil {
		log.Println(err)
	}
}

func connect() (*amqp.Connection, error) {
	var counts int64
	var backOff = 1 * time.Second

	var connection *amqp.Connection

	for {
		c, err := amqp.Dial("amqp://guest:guest@rabbitmq")
		if err != nil {
			fmt.Println("RabbitMQ not yet ready...")
			counts++
		} else {
			log.Println("Connected to Rabbit MQ")
			connection = c
			break
		}

		if counts > 5 {
			fmt.Println(err)
			return nil, err
		}

		backOff = time.Duration(math.Pow(float64(counts), 2)) * time.Second
		log.Println("Backing off...")

		time.Sleep(backOff)

	}

	return connection, nil
}
