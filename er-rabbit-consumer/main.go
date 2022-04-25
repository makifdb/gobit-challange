package main

import (
	"encoding/json"
	"er-rabbit-consumer/db"
	"er-rabbit-consumer/models"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func getRabbitUrl() string {
	RABBITMQ_DEFAULT_USER := os.Getenv("RABBITMQ_DEFAULT_USER")
	RABBITMQ_DEFAULT_PASS := os.Getenv("RABBITMQ_DEFAULT_PASS")
	RABBITMQ_HOST := os.Getenv("RABBITMQ_HOST")
	RABBITMQ_PORT := os.Getenv("RABBITMQ_PORT")

	return fmt.Sprintf("amqp://%s:%s@%s:%s/", RABBITMQ_DEFAULT_USER, RABBITMQ_DEFAULT_PASS, RABBITMQ_HOST, RABBITMQ_PORT)
}

func main() {
	conn, err := amqp.Dial(getRabbitUrl())
	failOnError(err, "Failed to connect RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"message", // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			var message models.Message
			err := json.Unmarshal(d.Body, &message)
			if err != nil {
				log.Printf("Error: %s", err)
			}
			var time time.Time
			time = message.Time

			db.WriteDatabase(time, message.Rate.EUR, message.Rate.TRY)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
