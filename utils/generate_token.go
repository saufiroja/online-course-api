package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/saufiroja/online-course-api/config"
	"github.com/saufiroja/online-course-api/models/dto"
	"github.com/saufiroja/online-course-api/models/entity"
)

func GenerateAccessToken(user *dto.UserResponse, oauthClient *entity.OauthClient, conf *config.AppConfig) (entity.OauthAccessToken, time.Time, error) {
	expirationTime := time.Now().Add(24 * 365 * time.Hour)

	secret := conf.Jwt.Secret

	claims := &dto.ClaimsReponse{
		ID:      user.ID,
		Email:   user.Email,
		Name:    user.Name,
		IsAdmin: false,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	if oauthClient.Name == "web-admin" {
		claims.IsAdmin = true
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return entity.OauthAccessToken{}, time.Time{}, err
	}

	oauthAccessToken := entity.OauthAccessToken{
		OauthClientID: &oauthClient.ID,
		UserID:        user.ID,
		Token:         tokenString,
		Scope:         "*",
		ExpiredAt:     &expirationTime,
	}

	return oauthAccessToken, expirationTime, nil
}

func GenerateRefreshToken(user *dto.UserResponse, oauthAccessToken *entity.OauthAccessToken) (entity.OauthRefreshToken, error) {
	expirationTimeOauthAccessToken := time.Now().Add(24 * 366 * time.Hour)

	oauthRefreshToken := entity.OauthRefreshToken{
		OauthAccessTokenID: &oauthAccessToken.ID,
		UserID:             user.ID,
		Token:              RandString(128),
		ExpiredAt:          &expirationTimeOauthAccessToken,
	}

	return oauthRefreshToken, nil
}

func AccessToken(user *dto.UserResponse, oauthRefresh *entity.OauthRefreshToken, conf *config.AppConfig) (*entity.OauthAccessToken, error) {
	secret := conf.Jwt.Secret
	expirationTime := time.Now().Add(24 * 365 * time.Hour)

	claims := &dto.ClaimsReponse{
		ID:      user.ID,
		Name:    user.Name,
		Email:   user.Email,
		IsAdmin: false,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	if *oauthRefresh.OauthAccessToken.OauthClientID == 2 {
		claims.IsAdmin = true
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}

	oauthAccessToken := entity.OauthAccessToken{
		OauthClientID: oauthRefresh.OauthAccessToken.OauthClientID,
		UserID:        user.ID,
		Token:         tokenString,
		Scope:         "*",
		ExpiredAt:     &expirationTime,
	}

	return &oauthAccessToken, nil
}

func RefreshToken(oauthAccess *entity.OauthAccessToken, oauthRefresh *entity.OauthRefreshToken) (entity.OauthRefreshToken, error) {
	expirationTimeOauthRefreshToken := time.Now().Add(24 * 366 * time.Hour)

	oauthRefreshToken := entity.OauthRefreshToken{
		OauthAccessTokenID: &oauthAccess.ID,
		UserID:             oauthRefresh.UserID,
		Token:              RandString(128),
		ExpiredAt:          &expirationTimeOauthRefreshToken,
	}

	return oauthRefreshToken, nil
}
