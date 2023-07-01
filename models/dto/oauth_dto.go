package dto

import "github.com/golang-jwt/jwt/v4"

type LoginRequestBody struct {
	Email        string `json:"email" binding:"email"`
	Password     string `json:"password" binding:"required"`
	ClientID     string `json:"client_id" binding:"required"`
	ClientSecret string `json:"client_secret" binding:"required"`
}

type RefreshTokenRequestBody struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Type         string `json:"type"`
	ExpiredAt    string `json:"expired_at"`
	Scope        string `json:"scope"`
}

type UserResponse struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ClaimsReponse struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	IsAdmin bool   `json:"is_admin,omitempty"`
	jwt.RegisteredClaims
}

type MapClaimsResponse struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	IsAdmin       bool   `json:"is_admin,omitempty"`
	jwt.MapClaims `json:"omitempty"`
}
