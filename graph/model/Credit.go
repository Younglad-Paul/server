package model

import "time"

type Credit struct {
	ID     string   `json:"id"`
	UserID string   `json:"userID"`
	Amount *float64 `json:"amount,omitempty"`
	Timestamp   time.Time `bson:"timestamp"`
}
