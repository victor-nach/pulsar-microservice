package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/victor-nach/pulsar-microservice/src/api-gateway/api/client"
	"github.com/victor-nach/pulsar-microservice/src/api-gateway/api/handlers"
)

// Router - returns a gin router
func Router() *gin.Engine {
	// create grpc client
	client := client.Cli
	handler := &handlers.Handler{Client: client}
	router := gin.New()
	router.GET("/", handler.HomePage)
	router.POST("/acquire-targets", handler.AcquireTargets)
	router.GET("/list-targets", handler.ListTargets)
	return router
}
