package campaign

type CampaignInput struct {
	UserID           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	Description      string `json:"description"`
	Perk             string `json:"perk"`
	GoalAmount       int    `json:"goal_amount"`
}
