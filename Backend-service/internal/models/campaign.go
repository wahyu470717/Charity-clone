package models

import (
	"time"
)

type Campaigns struct {
	CampaignID  int64     `json:"campaign_id" db:"campaign_id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	Target      float64   `json:"target_amount" db:"target_amount"`
	Current     float64   `json:"current_amount" db:"current_amount"`
	ImageURL    string    `json:"image_url" db:"image_url"`
	IsActive    bool      `json:"is_active" db:"is_active"`
	CreatedBy   string    `json:"created_by" db:"created_by"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	ModifiedBy  string    `json:"modified_by" db:"modified_by"`
	ModifiedAt  time.Time `json:"modified_at" db:"modified_at"`
}
