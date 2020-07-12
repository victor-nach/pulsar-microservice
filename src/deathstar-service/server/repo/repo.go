package repo

import (
	//"fmt"
	//"log"
	//"os"

	"context"
	"encoding/json"
	"log"

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
	var payload []interface{}
	data, _ := json.Marshal(targets.Data)
	_ = json.Unmarshal(data, &payload)

	_, err := r.Collection.InsertMany(context.TODO(), payload)
	if err != nil {
		log.Println("Failed to save targets to the db", err)
	}
	return nil
}
