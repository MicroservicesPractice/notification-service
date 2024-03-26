package api

import (
	"log"
	"notification-service/app/api/email"
	"notification-service/app/api/sms"
	"notification-service/app/helpers"

	amqp "github.com/rabbitmq/amqp091-go"
)

var RABBIT_MQ_CONNECTION = helpers.GetEnv("RABBIT_MQ_CONNECTION")

func RabbitMqControllers() {
	log.Println("Start")
	conn, err := amqp.Dial(RABBIT_MQ_CONNECTION)
	if err != nil {
		log.Fatalf("unable to open connect to RabbitMQ server. Error: %s", err)
	}

	defer func() {
		_ = conn.Close() // Закрываем подключение в случае удачной попытки подключения
	}()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("failed to open a channel. Error: %s", err)
	}

	defer func() {
		_ = ch.Close() // Закрываем подключение в случае удачной попытки подключения
	}()

	err = ch.ExchangeDeclare(
		"notification", // name
		"fanout",       // type
		true,           // durable
		false,          // auto-deleted
		false,          // internal
		false,          // no-wait
		nil,            // arguments
	)
	if err != nil {
		log.Fatalf("failed to declare exchange. Error: %s", err)
	}

	var forever chan struct{}

	email.EmailController(ch)
	sms.SmsController(ch)

	<-forever
}
