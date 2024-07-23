package model

import "time"

type Notification struct {
	ID        string    `bson:"id"`
	UserID    string    `bson:"userid"`
	Message   string    `bson:"message"`
	Seen      bool      `bson:"seen"`
	Timestamp time.Time `bson:"timestamp"`
}
