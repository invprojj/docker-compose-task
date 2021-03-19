package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/streadway/amqp"
	"log"
	"strconv"
	"time"
)

const CounterKey = "BLOCKSSENT"

var RabbitQueue amqp.Queue
var RabbitChannel *amqp.Channel

var db sql.DB
var err error


func main() {

	log.Println("Start App")

	startApi()

	time.Sleep(40 * time.Second)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	CreateDBConnection(psqlInfo)
	db.Exec(SqlCreateTable)

	CreateAndSetupRabbitMqConnection()

	msgs, err := RabbitChannel.Consume(
		RabbitQueue.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)

			res := string(d.Body[:])

			log.Println(res)

			value, err := strconv.ParseInt(hexaNumberToInteger(res), 16, 64)

			if err != nil {
				log.Printf ("Conversion failed: %s\n", err)
			} else {
				log.Printf ("Hexadecimal '%s' is integer %d (%X)",
					d.Body, value, value)
			}

			InsertBlockPostgre(value)

			Incr(CounterKey)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
