package interfaces

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saufiroja/online-course-api/models/dto"
	"github.com/saufiroja/online-course-api/models/entity"
)

// interface repository
type OauthAccessTokenRepository interface {
	InsertAccessToken(oauth entity.OauthAccessToken) (*entity.OauthAccessToken, error)
	DeleteAccessToken(oauth entity.OauthAccessToken) error
	FindOneByAccessToken(accessToken string) (*entity.OauthAccessToken, error)
}

type OauthClientRepository interface {
	FindByClientIDAndClientSecret(clientID, clientSecret string) (*entity.OauthClient, error)
}

type OauthRefreshTokenRepository interface {
	InsertRefreshToken(oauth entity.OauthRefreshToken) (*entity.OauthRefreshToken, error)
	FindOneByRefreshToken(refreshToken string) (*entity.OauthRefreshToken, error)
	FindOneByOauthAccessTokenID(oauthAccessTokenID int) (*entity.OauthRefreshToken, error)
	DeleteRefreshToken(oauth entity.OauthRefreshToken) error
}

// interface service
type OauthService interface {
	Login(input *dto.LoginRequestBody) (*dto.LoginResponse, error)
	RefreshToken(input *dto.RefreshTokenRequestBody) (*dto.LoginResponse, error)
	LogOut(accessToken string) error
}

// interface handler
type OauthHandler interface {
	Login(c *fiber.Ctx) error
	RefreshToken(c *fiber.Ctx) error
}
