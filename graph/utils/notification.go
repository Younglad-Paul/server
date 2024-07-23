package utils

import (
	"backend/graph/model"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateNotification(ctx context.Context, collection *mongo.Collection, userID, message string) error {
	notification := &model.Notification{
		ID:        uuid.New().String(),
		UserID:    userID,
		Message:   message,
		Timestamp: time.Now(),
	}

	_, err := collection.InsertOne(ctx, notification)
	if err != nil {
		return fmt.Errorf("error inserting notification: %v", err)
	}

	return nil
}
