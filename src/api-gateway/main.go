package main

import (
	"os"
	"github.com/victor-nach/pulsar-microservice/src/api-gateway/api/routes"
)

func main() {
	router := routes.Router()
	PORT, ok := os.LookupEnv("PORT")
	if !ok {
		PORT = "8000"
	}
	router.Run(":" + PORT)
}
