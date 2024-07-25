package model

import "time"

type Investment struct {
	ID     string   `json:"id"`
	UserID string   `json:"userID"`
	Amount *float64 `json:"amount,omitempty"`
	Name   *string  `json:"name,omitempty"`
	Timestamp   time.Time `bson:"timestamp"`

}
