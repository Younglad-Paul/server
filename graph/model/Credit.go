package model

import "time"

type Credit struct {
	ID     string   `bson:"id"`
	UserID string   `bson:"userID"`
	Amount *float64 `bson:"amount,omitempty"`
	Timestamp   time.Time `bson:"timestamp"`
}
