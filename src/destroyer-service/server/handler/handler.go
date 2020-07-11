package handler

import (
	"context"
	"encoding/json"
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/victor-nach/pulsar-microservice/src/destroyer-service/server/repo"
	"github.com/victor-nach/pulsar-microservice/src/destroyer-service/pb"
)

type Handler struct {
	Repo	repo.RepoInterface
	Producer	pulsar.Producer
}

// AcquireTargets - trigger and event to save targets
// passes the list of targets in the eventpayload 
func (h *Handler) AcquireTargets(ctx context.Context, req *pb.Targets) (*pb.Response, error) {
	payload, _ := json.Marshal(req)
	messageId, _ := h.Producer.Send(ctx, &pulsar.ProducerMessage{
		Payload: payload,
	})
	return &pb.Response{MessageId: string(messageId.Serialize())}, nil
}

// ListTargets - get the list of all the targets saved in the db 
func (h *Handler) ListTargets(ctx context.Context, req *pb.Request) (*pb.Targets, error) {
	targets, err := h.Repo.GetTargets()
	if err != nil {
		return nil, err
	}
	return &pb.Targets{Data: targets}, nil
}