package model

import "time"

type History struct {
	ID          string    `bson:"id"`
	UserID      string    `bson:"userid"`
	EntityType  string    `bson:"entity_type"`
	ActionType  string    `bson:"action_type"`
	ChangedData string    `bson:"changed_data"`
	Timestamp   time.Time `bson:"timestamp"`
}
