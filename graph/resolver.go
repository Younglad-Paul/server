package graph

import (
	"go.mongodb.org/mongo-driver/mongo"
)

// Resolver holds your resolver implementations.
type Resolver struct {
	MongoClient *mongo.Client
}

// NewResolver returns a new Resolver with the given MongoDB client.
func NewResolver(client *mongo.Client) *Resolver {
	return &Resolver{
		MongoClient: client,
	}
}
