package utils

import (
	"context"
	"time"

	"backend/graph/model"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateCredit(ctx context.Context, collection *mongo.Collection, userID string) error {
	initialAmount := 0.00
	credit := &model.Credit{
		ID:        uuid.New().String(),
		UserID:    userID,
		Amount:    &initialAmount,
		Timestamp: time.Now(),
	}

	_, err := collection.InsertOne(ctx, credit)
	return err
}
