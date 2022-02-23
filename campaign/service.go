package campaign

type Service interface {
	RegisterCampaign(input CampaignInput) (Campaign, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterCampaign(input CampaignInput) (Campaign, error) {
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
