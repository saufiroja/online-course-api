package oauth

import (
	"github.com/saufiroja/online-course-api/interfaces"
	"github.com/saufiroja/online-course-api/models/entity"
	"gorm.io/gorm"
)

type OauthRefreshTokenRepository struct {
	db *gorm.DB
}

func NewOauthRefreshTokenRepository(db *gorm.DB) interfaces.OauthRefreshTokenRepository {
	return &OauthRefreshTokenRepository{db}
}

// DeleteRefreshToken implements interfaces.OauthRefreshTokenRepository.
func (r *OauthRefreshTokenRepository) DeleteRefreshToken(oauth entity.OauthRefreshToken) error {
	err := r.db.Model(&entity.OauthRefreshToken{}).Delete(&oauth).Error
	if err != nil {
		return err
	}

	return nil
}

// FindOneByOauthAccessTokenID implements interfaces.OauthRefreshTokenRepository.
func (r *OauthRefreshTokenRepository) FindOneByOauthAccessTokenID(oauthAccessTokenID int) (*entity.OauthRefreshToken, error) {
	var oauth entity.OauthRefreshToken

	err := r.db.Model(&entity.OauthRefreshToken{}).
		Where("oauth_access_token_id = ?", oauthAccessTokenID).
		First(&oauth).Error
	if err != nil {
		return nil, err
	}

	return &oauth, nil
}

// FindOneByRefreshToken implements interfaces.OauthRefreshTokenRepository.
func (r *OauthRefreshTokenRepository) FindOneByRefreshToken(refreshToken string) (*entity.OauthRefreshToken, error) {
	var oauth entity.OauthRefreshToken
	err := r.db.Model(&entity.OauthRefreshToken{}).
		Preload("OauthAccessToken").
		Where("token = ?", refreshToken).
		First(&oauth).Error
	if err != nil {
		return nil, err
	}

	return &oauth, nil
}

// InsertRefreshToken implements interfaces.OauthRefreshTokenRepository.
func (r *OauthRefreshTokenRepository) InsertRefreshToken(oauth entity.OauthRefreshToken) (*entity.OauthRefreshToken, error) {
	err := r.db.Create(&oauth).Error
	if err != nil {
		return nil, err
	}

	return &oauth, nil
}
