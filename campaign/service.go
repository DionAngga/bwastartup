package campaign

type Service interface {
	GetCampaigns(UserID int) ([]Campaign, error)
	GetCampaignByID(input CampaignInput) (Campaign, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetCampaigns(userID int) ([]Campaign, error) {
	if userID != 0 {
		campaigns, err := s.repository.FindByUserID(userID)
		if err != nil {
			return campaigns, err
		}
		return campaigns, nil
	}
	campaigns, err := s.repository.FindAll()
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}

func (s *service) GetCampaignByID(input CampaignInput) (Campaign, error) {
	camps := Campaign{}
	camps.UserID = input.UserID
	camps.Name = input.Name
	camps.Description = input.Description
	camps.ShortDescription = input.ShortDescription
	camps.Perk = input.Perk
	camps.GoalAmount = input.GoalAmount
	newcampaign, err := s.repository.CreateCampaign(camps)
	if err != nil {
		return newcampaign, err
	}
	return newcampaign, nil
}
