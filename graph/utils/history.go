package utils

import (
	"context"
	"time"

	"backend/graph/model"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateHistoryRecord(ctx context.Context, collection *mongo.Collection, userID, entityType, actionType, changedData string) error {
	history := &model.History{
		ID:          uuid.New().String(),
		UserID:      userID,
		EntityType:  entityType,
		ActionType:  actionType,
		ChangedData: changedData,
		Timestamp:   time.Now(),
	}

	_, err := collection.InsertOne(ctx, history)
	return err
}
