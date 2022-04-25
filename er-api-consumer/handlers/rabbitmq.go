package handlers

import (
	"encoding/json"
	"er-api-consumer/models"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/streadway/amqp"
)

func CreateMessage() {
	r, err := GetData()
	if err != nil {
		log.Println(err)
	}
	var msg = models.Message{
		Time: time.Now(),
		Rate: struct {
			EUR float64
			TRY float64
		}{
			EUR: r.Rate.EUR,
			TRY: r.Rate.TRY,
		},
	}

	fmt.Println(msg)
	RabbitMessager(msg)
}

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

func RabbitMessager(msg models.Message) {
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

	body, _ := json.Marshal(msg)
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", body)
}
