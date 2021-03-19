package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/streadway/amqp"
)

var RabbitQueue amqp.Queue
var RabbitChannel *amqp.Channel

func main() {

	log.Println("Starting App")

	go func() {
		for true {
			handleRequests()
		}
	}()

	time.Sleep(40 * time.Second)

	CreateAndSetupRabbitMqConnection()

	for true {

		block, err := GetLastBlockFromNode()
		if err != nil {
			log.Println("Error while getting block", err)
			continue
		}

		if CurrentBlockNumber == block {
			log.Println("Received block is old: " + block)
			continue
		}

		PublishBlock(block)

		log.Println("Time to wait")

		delay := os.Getenv("INTERVAL")
		result, err := strconv.Atoi(delay)

		time.Sleep(time.Duration(result) * time.Second)
	}
}