package model

type Referral struct {
	ID     string `json:"id"`
	UserID string `json:"userID"`
	Link   string `json:"link,omitempty"`
	Count  float64 `json:"count"`
}
