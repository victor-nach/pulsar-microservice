package handler

import (
	"context"
	"encoding/json"
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/victor-nach/pulsar-microservice/src/deathstar-service/server/repo"
	"github.com/victor-nach/pulsar-microservice/src/deathstar-service/pb"
)

type Handler struct {
	Repo		repo.RepoInterface
	Producer	pulsar.Producer
}

func (h *Handler) AcquireTargets(ctx context.Context, req *pb.Targets) (*pb.Response, error) {
	payload, _ := json.Marshal(req)
	messageId, _ := h.Producer.Send(ctx, &pulsar.ProducerMessage{
		Payload: payload,
	})
	return &pb.Response{MessageId: string(messageId.Serialize())}, nil
}