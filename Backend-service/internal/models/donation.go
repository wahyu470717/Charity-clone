package models

import (
	"time"
)

type Donation struct {
	ID          int64     `json:"id" db:"id"`
	UserID      int64     `json:"user_id" db:"user_id"`
	CampaignID  int64     `json:"campaign_id" db:"campaign_id"`
	Amount      float64   `json:"amount" db:"amount"`
	IsAnonymous bool      `json:"is_anonymous" db:"is_anonymous"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}
