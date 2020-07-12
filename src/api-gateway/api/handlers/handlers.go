package handlers

import (
	"context"

	"github.com/gin-gonic/gin"
	u "github.com/victor-nach/pulsar-microservice/src/api-gateway/api/utils"
	"github.com/victor-nach/pulsar-microservice/src/api-gateway/pb"
)

type Handler struct {
	Client pb.DestroyerServiceClient
}

// HomePage - return a welcome message on the root route
func (r *Handler) HomePage(c *gin.Context) {
	u.ResSuccess(u.Res{Ctx: c, Status: 201, Msg: "Welcome to Pulsar microservice 2020"})
}

// AcquireTargets - make a grpc call on the destroyer service
func (r *Handler) AcquireTargets(c *gin.Context) {
	var targets []*pb.Target
	var err error
	// get targets from request body
	if err := c.ShouldBindJSON(&targets); err != nil {
		u.ResErr(u.Res{Ctx: c, Msg: "Invalid request", Err: err})
		return
	}
	// connect to the destroyer server
	conn := r.Client
	if err != nil {
		u.ResErr(u.Res{Ctx: c, Msg: "Destroyer service not available", Err: err})
		return
	}
	req := &pb.Targets{
		Data: targets,
	}
	_, err = conn.AcquireTargets(context.Background(), req)
	if err != nil {
		u.ResErr(u.Res{Ctx: c, Msg: "Error occurred calling destroyer service", Err: err})
		return
	}
	u.ResSuccess(u.Res{Ctx: c, Status: 201, Msg: "Successfully created targets!"})
}
