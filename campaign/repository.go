package campaign

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Campaign, error)
	FindByUserID(id int) ([]Campaign, error)
	Delete(id int) (Campaign, error)
	CreateCampaign(campaign Campaign) (Campaign, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Campaign, error) {
	var camps []Campaign
	err := r.db.Find(&camps).Preload("CampaignImages", "campaign_images.is_primary = 1").Error
	if err != nil {
		return camps, err
	}
	return camps, nil
}

func (r *repository) FindByUserID(id int) ([]Campaign, error) {
	var campaigns []Campaign
	err := r.db.Where("user_id = ?", id).Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}

func (r *repository) Delete(id int) (Campaign, error) {
	var campaign Campaign
	err := r.db.Where("id = ?", id).Delete(&campaign).Error
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}

func (r *repository) CreateCampaign(campaign Campaign) (Campaign, error) {
	err := r.db.Create(&campaign).Error
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}
