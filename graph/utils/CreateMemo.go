package utils

import (
	"backend/graph/model"
	"context"
	"math/rand"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

func GenerateMemo(generate int) string {
	const strings = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())
	gen := make([]byte, generate)
	for i := range gen {
		gen[i] = strings[rand.Intn(len(strings))]
	}

	prefix := "0x000"
	suffix := "EA"

	return prefix + string(gen) + suffix
}

func CreateMemo(ctx context.Context, collection *mongo.Collection, userID string) error {
	memo := &model.Memo{
		UserID: userID,
		Memo:   GenerateMemo(28),
	}

	_, err := collection.InsertOne(ctx, memo)
	return err
}
