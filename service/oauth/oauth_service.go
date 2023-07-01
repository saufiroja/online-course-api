package oauth

import (
	"time"

	"github.com/saufiroja/online-course-api/config"
	"github.com/saufiroja/online-course-api/interfaces"
	"github.com/saufiroja/online-course-api/models/dto"
	"github.com/saufiroja/online-course-api/utils"
)

type service struct {
	accessTokenRepository  interfaces.OauthAccessTokenRepository
	clientRepository       interfaces.OauthClientRepository
	refreshTokenRepository interfaces.OauthRefreshTokenRepository
	adminService           interfaces.AdminService
	userService            interfaces.UserService
	conf                   *config.AppConfig
}

func NewOauthService(
	accessTokenRepository interfaces.OauthAccessTokenRepository,
	clientRepository interfaces.OauthClientRepository,
	refreshTokenRepository interfaces.OauthRefreshTokenRepository,
	adminService interfaces.AdminService,
	userService interfaces.UserService,
	conf *config.AppConfig,
) interfaces.OauthService {
	return &service{
		accessTokenRepository:  accessTokenRepository,
		clientRepository:       clientRepository,
		refreshTokenRepository: refreshTokenRepository,
		adminService:           adminService,
		userService:            userService,
		conf:                   conf,
	}
}

// LogOut implements interfaces.OauthService.
func (s *service) LogOut(accessToken string) error {
	oauthAccessToken, err := s.accessTokenRepository.FindOneByAccessToken(accessToken)
	if err != nil {
		return utils.HandlerError(404, "access token not found")
	}

	oauthRefreshToken, err := s.refreshTokenRepository.FindOneByOauthAccessTokenID(int(oauthAccessToken.ID))
	if err != nil {
		return utils.HandlerError(404, "refresh token not found")
	}

	err = s.refreshTokenRepository.DeleteRefreshToken(*oauthRefreshToken)
	if err != nil {
		return utils.HandlerError(500, "failed to delete refresh token")
	}

	err = s.accessTokenRepository.DeleteAccessToken(*oauthAccessToken)
	if err != nil {
		return utils.HandlerError(500, "failed to delete access token")
	}

	return nil
}

// Login implements interfaces.OauthService.
func (s *service) Login(input *dto.LoginRequestBody) (*dto.LoginResponse, error) {
	oauthClient, err := s.clientRepository.FindByClientIDAndClientSecret(input.ClientID, input.ClientSecret)
	if err != nil {
		return nil, utils.HandlerError(404, "client not found")
	}

	var userResponse dto.UserResponse

	if oauthClient.Name == "web-admin" {
		dataAdmin, err := s.adminService.FindOneAdminByEmail(input.Email)
		if err != nil {
			return nil, utils.HandlerError(404, "admin not found")
		}

		userResponse = dto.UserResponse{
			ID:       dataAdmin.ID,
			Email:    dataAdmin.Email,
			Name:     dataAdmin.Name,
			Password: dataAdmin.Password,
		}
	} else {
		users, err := s.userService.FindUserByEmail(input.Email)
		if err != nil {
			return nil, utils.HandlerError(404, "user not found")
		}

		userResponse = dto.UserResponse{
			ID:       users.ID,
			Email:    users.Email,
			Name:     users.Name,
			Password: users.Password,
		}
	}

	err = utils.ComparePassword(userResponse.Password, input.Password)
	if err != nil {
		return nil, utils.HandlerError(401, "password not match")
	}

	accessToken, expires, err := utils.GenerateAccessToken(&userResponse, oauthClient, s.conf)
	if err != nil {
		return nil, utils.HandlerError(500, "generate access token failed")
	}

	oauthAccessToken, err := s.accessTokenRepository.InsertAccessToken(accessToken)
	if err != nil {
		return nil, utils.HandlerError(500, "insert access token failed")
	}

	oauthRefreshToken, err := utils.GenerateRefreshToken(&userResponse, oauthAccessToken)
	if err != nil {
		return nil, utils.HandlerError(500, "generate refresh token failed")
	}

	refreshToken, err := s.refreshTokenRepository.InsertRefreshToken(oauthRefreshToken)
	if err != nil {
		return nil, utils.HandlerError(500, "insert refresh token failed")
	}

	res := dto.LoginResponse{
		AccessToken:  oauthAccessToken.Token,
		RefreshToken: refreshToken.Token,
		Type:         "Bearer",
		ExpiredAt:    expires.Format(time.RFC3339),
		Scope:        "*",
	}

	return &res, nil
}

// RefreshToken implements interfaces.OauthService.
func (s *service) RefreshToken(input *dto.RefreshTokenRequestBody) (*dto.LoginResponse, error) {
	oauthRefreshToken, err := s.refreshTokenRepository.FindOneByRefreshToken(input.RefreshToken)
	if err != nil {
		return nil, utils.HandlerError(404, "refresh token not found")
	}

	if oauthRefreshToken.ExpiredAt.Before(time.Now()) {
		return nil, utils.HandlerError(401, "refresh token expired")
	}

	var userResponse dto.UserResponse

	if *oauthRefreshToken.OauthAccessToken.OauthClientID == 2 {
		admin, err := s.adminService.FindOneAdminByID(int(oauthRefreshToken.UserID))
		if err != nil {
			return nil, utils.HandlerError(404, "admin not found")
		}

		userResponse = dto.UserResponse{
			ID:    admin.ID,
			Email: admin.Email,
			Name:  admin.Name,
		}
	} else {
		users, err := s.userService.FindUserById(int(oauthRefreshToken.UserID))
		if err != nil {
			return nil, utils.HandlerError(404, "user not found")
		}

		userResponse = dto.UserResponse{
			ID:    users.ID,
			Email: users.Email,
			Name:  users.Name,
		}
	}

	accessToken, err := utils.AccessToken(&userResponse, oauthRefreshToken, s.conf)
	if err != nil {
		return nil, utils.HandlerError(500, "refresh token failed")
	}

	oauthAccessToken, err := s.accessTokenRepository.InsertAccessToken(*accessToken)
	if err != nil {
		return nil, utils.HandlerError(500, "insert access token failed")
	}

	refreshToken, err := utils.RefreshToken(oauthAccessToken, oauthRefreshToken)
	if err != nil {
		return nil, utils.HandlerError(500, "refresh token failed")
	}

	resRefresh, err := s.refreshTokenRepository.InsertRefreshToken(refreshToken)
	if err != nil {
		return nil, utils.HandlerError(500, "insert refresh token failed")
	}

	err = s.refreshTokenRepository.DeleteRefreshToken(*oauthRefreshToken)
	if err != nil {
		return nil, utils.HandlerError(500, "delete refresh token failed")
	}

	err = s.accessTokenRepository.DeleteAccessToken(*oauthAccessToken)
	if err != nil {
		return nil, utils.HandlerError(500, "delete access token failed")
	}

	res := dto.LoginResponse{
		AccessToken:  oauthAccessToken.Token,
		RefreshToken: resRefresh.Token,
		Type:         "Bearer",
		ExpiredAt:    accessToken.ExpiredAt.Format(time.RFC3339),
		Scope:        "*",
	}

	return &res, nil
}
