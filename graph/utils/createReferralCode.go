package utils

import (
	"context"
	"math/rand"
	"time"

	"backend/graph/model"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

func GenerateRandomString(n int) string {
	const lettersAndDigits = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	for i := range b {
		b[i] = lettersAndDigits[rand.Intn(len(lettersAndDigits))]
	}
	return string(b)
}

func CreateReferralCode(ctx context.Context, collection *mongo.Collection, userID string) error {
	ref := &model.Referral{
		ID:     uuid.New().String(),
		UserID: userID,
		Link:   GenerateRandomString(8),
		Count:  0,
	}

	_, err := collection.InsertOne(ctx, ref)
	return err
}
