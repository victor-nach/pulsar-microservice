package repo

import (
	//"fmt"
	//"log"
	//"os"

	"context"
	"github.com/victor-nach/pulsar-microservice/src/deathstar-service/pb"
	"go.mongodb.org/mongo-driver/mongo"
	//"go.mongodb.org/mongo-driver/mongo/options"
)

// Repository - Repo 
type Repository struct {
	Collection *mongo.Collection
}

// RepoInterface - Repository interface 
type RepoInterface interface {
	SaveTarget(targets pb.Targets) error
}

// SaveTarget - Save targets to the db
func (r *Repository) SaveTarget(targets pb.Targets) error {
	r.Collection.InsertMany(context.TODO(), []interface{}{targets.Data})
	return nil
}
