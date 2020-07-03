package repo

import (
	//"context"
	//"fmt"
	//"log"
	//"os"

	"go.mongodb.org/mongo-driver/mongo"
	//"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	Collection *mongo.Collection
}

type RepoInterface interface {
	//GetAll(ctx context.Context) ([]*Consignment, error)
}
