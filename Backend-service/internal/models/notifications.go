package models

import "time"

type Notifications struct {
	ID         int64     `json:"id" db:"id"`
	UserID     int64     `json:"user_id" db:"user_id"`
	Title      string    `json:"title" db:"title"`
	Message    string    `json:"message" db:"message"`
	IsRead     bool      `json:"is_read" db:"is_read"`
	CreatedBy  string    `json:"created_by" db:"created_by"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	ModifiedBy string    `json:"modified_by" db:"modified_by"`
	ModifiedAt time.Time `json:"modified_at" db:"modified_at"`
}
