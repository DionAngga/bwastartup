package campaign

type CampsFormatter struct {
	UserID           int    `json:"user_id"`
	ID               int    `json:"id"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	ShortDescription string `json:"short_description"`
}

func FormatCamps(campaign Campaign) CampsFormatter {
	formatter := CampsFormatter{
		UserID:           campaign.UserID,
		ID:               campaign.ID,
		Name:             campaign.Name,
		Description:      campaign.Description,
		ShortDescription: campaign.ShortDescription,
	}

	return formatter
}
