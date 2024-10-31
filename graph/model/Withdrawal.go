package model

import "time"

type Withdrawal struct {
	UserID    string
	From      string
	Amount    float64
	Status    bool
	Timestamp time.Time
	
}
