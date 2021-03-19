package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/streadway/amqp"
)

const OnStartDelay = 42

var RabbitQueue amqp.Queue
var RabbitChannel *amqp.Channel

func main() {

	log.Println("Starting App")

	startApi()

	time.Sleep(time.Duration(OnStartDelay) * time.Second)

	CreateAndSetupRabbitMqConnection()

	for {

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