package main

import (
	"github.com/streadway/amqp"
	"os"
)

const QueueName = "blocks"

func CreateAndSetupRabbitMqConnection() error {

	conn, err := amqp.Dial(os.Getenv("RABBIT_CONN_STR"))

	//conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	channel, err := conn.Channel()
	failOnError(err, "Failed to open a channel")

	queue, err := channel.QueueDeclare(
		QueueName, // name
		true,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)

	RabbitQueue = queue
	RabbitChannel = channel

	return err
}


