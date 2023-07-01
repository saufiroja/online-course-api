package oauth

import (
	"github.com/saufiroja/online-course-api/interfaces"
	"github.com/saufiroja/online-course-api/models/entity"
	"gorm.io/gorm"
)

type oauthAccessTokenRepository struct {
	db *gorm.DB
}

func NewOauthAccessTokenRepository(db *gorm.DB) interfaces.OauthAccessTokenRepository {
	return &oauthAccessTokenRepository{db}
}

// DeleteAccessToken implements interfaces.OauthAccessTokenRepository.
func (r *oauthAccessTokenRepository) DeleteAccessToken(oauth entity.OauthAccessToken) error {
	err := r.db.Model(&entity.OauthAccessToken{}).Delete(&oauth).Error
	if err != nil {
		return err
	}

	return nil
}

// FindOneByAccessToken implements interfaces.OauthAccessTokenRepository.
func (r *oauthAccessTokenRepository) FindOneByAccessToken(accessToken string) (*entity.OauthAccessToken, error) {
	var oauth entity.OauthAccessToken
	err := r.db.Model(&entity.OauthAccessToken{}).Where("token = ?", accessToken).First(&oauth).Error
	if err != nil {
		return nil, err
	}

	return &oauth, nil
}

// InsertAccessToken implements interfaces.OauthAccessTokenRepository.
func (r *oauthAccessTokenRepository) InsertAccessToken(oauth entity.OauthAccessToken) (*entity.OauthAccessToken, error) {
	err := r.db.Model(&entity.OauthAccessToken{}).Create(&oauth).Error
	if err != nil {
		return nil, err
	}

	return &oauth, nil
}
