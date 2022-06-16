package campaign

type CampaignFormatter struct {
	UserID int    `json:"user_id"`
	ID     int    `json:"id"`
	Name   string `json:"name"`
	//Description 	string `json:"description"`
	ShortDescription string `json:"short_description"`
	ImageURL         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
}

func FormatCampaign(campaign Campaign) CampaignFormatter {
	CampaignFormatter := CampaignFormatter{}
	CampaignFormatter.UserID = campaign.UserID
	CampaignFormatter.ID = campaign.ID
	CampaignFormatter.Name = campaign.Name
	CampaignFormatter.ShortDescription = campaign.ShortDescription
	CampaignFormatter.GoalAmount = campaign.GoalAmount
	CampaignFormatter.CurrentAmount = campaign.CurrentAmount
	CampaignFormatter.ImageURL = ""

	if len(campaign.CampaignImages) > 0 {
		CampaignFormatter.ImageURL = campaign.CampaignImages[0].FileName
	}
	return CampaignFormatter
}

func FormatCampaigns(campaigns []Campaign) []CampaignFormatter {
	campaignFormatter := []CampaignFormatter{}
	for _, campaign := range campaigns {
		campaignFormatter = append(campaignFormatter, FormatCampaign(campaign))
	}
	return campaignFormatter
}
