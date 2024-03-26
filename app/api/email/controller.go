package email

import (
	"encoding/json"
	"fmt"
	"log"
	"notification-service/app/helpers"

	amqp "github.com/rabbitmq/amqp091-go"
)

var RABBIT_MQ_CONNECTION = helpers.GetEnv("RABBIT_MQ_CONNECTION")

func EmailController(ch *amqp.Channel) {
	sendEmailInfoMessage(ch)
}

func sendEmailInfoMessage(ch *amqp.Channel) {
	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		log.Fatalf("failed to declare a queue. Error: %s", err)
	}

	err = ch.QueueBind(
		q.Name,              // queue name
		"email.status.info", // routing key
		"notification",      // exchange
		false,
		nil)
	if err != nil {
		log.Fatalf("failed to queue bind. Error: %s", err)
	}

	messages, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatalf("failed to register a consumer. Error: %s", err)
	}

	// var forever chan struct{}

	go func() {
		for message := range messages {
			fmt.Println("EMAIL", message)
			type Message struct {
				Status string `json:"message"`
			}
			var data Message

			err := json.Unmarshal(message.Body, &data)
			if err != nil {
				log.Printf("Failed to marshal JSON: %v", err)
			}

			if data.Status == EMAIL_CONFIRMATION_COMPLETE {
				log.Print("Correct!")
			}

			log.Printf("received a message: %s", message.Body)
			log.Print(data.Status)

		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	// <-forever
}
