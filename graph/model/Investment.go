package model


type Investment struct {
	ID     string   `json:"id"`
	UserID string   `json:"userID"`
	Name   *string  `json:"name,omitempty"`
	Amount *float64 `json:"amount,omitempty"`
}
