package repo

import (
	//"fmt"
	//"log"
	//"os"

	"context"
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/victor-nach/pulsar-microservice/src/deathstar-service/pb"
	"go.mongodb.org/mongo-driver/mongo"
	//"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	Collection *mongo.Collection
}

type RepoInterface interface {
	SaveTarget(message pulsar.Message) error
}

func (r *Repository) SaveTarget (msg pulsar.Message) {
	var targets []pb.Target
	r.Collection.InsertMany(context.TODO(), targets)
}
