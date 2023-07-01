package oauth

import (
	"github.com/saufiroja/online-course-api/interfaces"
	"github.com/saufiroja/online-course-api/models/entity"
	"gorm.io/gorm"
)

type oauthClientRepository struct {
	db *gorm.DB
}

func NewOauthClientRepository(db *gorm.DB) interfaces.OauthClientRepository {
	return &oauthClientRepository{db}
}

// FindByClientIDAndClientSecret implements interfaces.OauthClientRepository.
func (r *oauthClientRepository) FindByClientIDAndClientSecret(clientID string, clientSecret string) (*entity.OauthClient, error) {
	var oauthClient entity.OauthClient
	err := r.db.Model(&entity.OauthClient{}).
		Where("client_id = ?", clientID).
		Where("client_secret = ?", clientSecret).
		First(&oauthClient).Error

	if err != nil {
		return nil, err
	}

	return &oauthClient, nil
}
