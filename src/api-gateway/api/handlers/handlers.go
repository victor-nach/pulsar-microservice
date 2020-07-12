package handlers

import (
	"github.com/gin-gonic/gin"
	u "github.com/victor-nach/pulsar-microservice/src/api-gateway/api/utils"
)

// HomePage - return a welcome message on the root route
func HomePage(c *gin.Context) {
	u.ResSuccess(u.Res{Ctx: c, Status: 201, Msg: "Welcome to Pulsar microservice 2020"})
}
