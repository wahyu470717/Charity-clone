package response

type CampaignResponse struct {
	CampaignID  int64   `json:"campaign_id"`
	Title       string  `json:"title"`
	Content     string  `json:"content"`
	ImageUrl    string  `json:"image_url"`
	Target      float64 `json:"target_amount"`
	Current     float64 `json:"current_amount"`
	Description string  `json:"description"`
}
