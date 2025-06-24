package request

type DonationRequest struct {
	CampaignID  int64   `json:"campaign_id" validate:"required"`
	Amount      float64 `json:"amount" validate:"required,gt=0"`
	IsAnonymous bool    `json:"is_anonymous"`
}
