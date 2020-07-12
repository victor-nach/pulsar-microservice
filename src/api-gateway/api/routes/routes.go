package routes

import (
	"github.com/gin-gonic/gin"
	c "github.com/victor-nach/pulsar-microservice/src/api-gateway/api/handlers"
)

// Router - returns a gin router
func Router() *gin.Engine {
	router := gin.New()
	router.GET("/", c.HomePage)
	return router
}
