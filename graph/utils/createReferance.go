package utils

import (
	"context"

	"backend/graph/model"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateReference(ctx context.Context, collection *mongo.Collection, userID string) error {
	initialAmount := 0.00
	reference := &model.Reference{
		ID:     uuid.New().String(),
		UserID: userID,
		Amount: &initialAmount,
	}

	_, err := collection.InsertOne(ctx, reference)
	return err
}
