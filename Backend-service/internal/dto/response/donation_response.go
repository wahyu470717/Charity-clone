package response

import "time"

type DonationResponse struct {
	ID          int64     `json:"id"`
	Amount      float64   `json:"amount"`
	IsAnonymous bool      `json:"is_anonymous"`
	CreatedAt   time.Time `json:"created_at"`
}
