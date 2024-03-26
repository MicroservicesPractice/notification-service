package main

import (
	"notification-service/app/api"
	"notification-service/app/config/initializers"
	"notification-service/app/helpers"
)

func init() {
	helpers.CheckRequiredEnvs()

	initializers.InitLogger()
}

func main() {
	api.RabbitMqControllers()
}
