package main

import (
	log "github.com/sirupsen/logrus"

	"notification-service/app/config/initializers"
	"notification-service/app/helpers"
)

var SERVER_PORT = helpers.GetEnv("SERVER_PORT")

func init() {
	helpers.CheckRequiredEnvs()

	initializers.InitLogger()
}

func main() {
	initializers.RmqInit()

	log.Infof("Server has been started on port %v", SERVER_PORT)
}
