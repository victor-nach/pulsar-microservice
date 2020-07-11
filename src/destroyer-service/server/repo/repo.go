package repo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/victor-nach/pulsar-microservice/src/destroyer-service/pb"
)

type Repository struct {
	Collection *mongo.Collection
}

type RepoInterface interface {
	GetTargets() ([]*pb.Target, error)
}

// GetTargets - returns a list of all the targets saved in the db
func (r *Repository) GetTargets() ([]*pb.Target, error) {
	var targets []*pb.Target
	cur, err := r.Collection.Find(context.TODO(), bson.D{{}}, options.Find().SetLimit(100))
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())
	if err = cur.Err(); err != nil {
		return nil, err
	}
	for cur.Next(context.TODO()) {
		var target pb.Target
		cur.Decode(&target)
		targets = append(targets, &target)
	}
	return targets, nil
}
