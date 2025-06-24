package request

import "time"

type CreateCampaignRequest struct {
	CampaignID int64     `json:"campaign_id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	ImageUrl   string    `json:"image_url,omitempty"`
	CreatedBy  string    `json:"created_by"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedBy string    `json:"modified_by"`
	ModifiedAt time.Time `json:"modified_at"`
	Target     float64   `json:"target_amount"`
	Current    float64   `json:"current_amount"`
}

type UpdateCampaignRequest struct {
	Title       string  `form:"title"`
	Description string  `form:"description"`
	Target      float64 `form:"target"`
}
