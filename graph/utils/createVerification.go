package utils

import (
	"context"
	"fmt"

	"backend/graph/model"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateVerificationCode(ctx context.Context, collection *mongo.Collection, Email string) error {
	verify := &model.Verify{
		UniqueVerifier: uuid.New().String(),
		Email:          Email,
	}
	_, err := collection.InsertOne(ctx, verify)
	if err != nil {
		return fmt.Errorf("error passing email: %v", err)
	}

	return nil

}