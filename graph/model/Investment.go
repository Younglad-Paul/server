package model

import "time"

type Investment struct {
	ID        string    `bson:"id"`
	UserID    string    `bson:"userID"`
	Amount    *float64  `bson:"amount,omitempty"`
	Name      *string   `bson:"name,omitempty"`
	Timestamp time.Time `bson:"timestamp"`
}
