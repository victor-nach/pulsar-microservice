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

	req := &pb.Targets{
		Data: targets,
	}
	_, err = r.Client.AcquireTargets(context.Background(), req)
	if err != nil {
		u.ResErr(u.Res{Ctx: c, Msg: "Error occurred calling destroyer service", Err: err})
		return
	}
	u.ResSuccess(u.Res{Ctx: c, Status: 201, Msg: "Successfully created targets!"})
}

// ListTargets - make a grpc call on the destroyer service
func (r *Handler) ListTargets(c *gin.Context) {
	//var targets []*pb.Target
	response, err := r.Client.ListTargets(context.Background(), &pb.Request{})
	if err != nil {
		u.ResErr(u.Res{Ctx: c, Msg: "Error occurred calling destroyer service", Err: err})
		return
	}
	// unmarshall data into targets
	u.ResSuccess(u.Res{Ctx: c, Status: 200, Msg: "Successfully retreived targets!", Data: response})
}
