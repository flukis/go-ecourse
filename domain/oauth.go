package domain

import (
	"e-course/pkg/resp"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

// ------- Oauth Client

type OauthClient struct {
	ID           int64          `json:"id"`
	ClientID     string         `json:"client_id"`
	ClientSecret string         `json:"client_secret"`
	Name         string         `json:"name"`
	Redirect     string         `json:"redirect"`
	Scope        string         `json:"scope"`
	CreatedAt    *time.Time     `json:"created_at"`
	UpdatedAt    *time.Time     `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at"`
}

type OauthClientRepository interface {
	FindByClientIDnClientSecret(clientId, clientSecret string) (*OauthClient, *resp.ErrorResp)
}

// ------- Oauth Client
// ------- Oauth Access Token

type OauthAccessToken struct {
	ID            int64          `json:"id"`
	OauthClient   *OauthClient   `gorm:"foreignKey:OauthClientID;reference:ID"`
	OauthClientID *int64         `json:"oauth_client_id"`
	UserID        int64          `json:"user_id"`
	Token         string         `json:"token"`
	Scope         string         `json:"scope"`
	ExpiredAt     *time.Time     `json:"expired_at"`
	CreatedAt     *time.Time     `json:"created_at"`
	UpdatedAt     *time.Time     `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at"`
}

type OauthAccessTokenRepository interface {
	Create(entity OauthAccessToken) (*OauthAccessToken, *resp.ErrorResp)
	Delete(OauthAccessToken) *resp.ErrorResp
	FindOneByAccessToken(accessToken string) (*OauthAccessToken, *resp.ErrorResp)
	FindOneByOauthAccessTokenID(id int) (*OauthAccessToken, *resp.ErrorResp)
}

// ------- Oauth Access Token
// ------- Oauth Refresh Token

type OauthRefreshToken struct {
	ID                 int64             `json:"id"`
	OauthAccessToken   *OauthAccessToken `gorm:"foreignKey:OauthAccessToken;reference:ID"`
	OauthAccessTokenID *int64            `json:"oauth_access_token_id"`
	UserID             int64             `json:"user_id"`
	Token              string            `json:"token"`
	ExpiredAt          *time.Time        `json:"expired_at"`
	CreatedAt          *time.Time        `json:"created_at"`
	UpdatedAt          *time.Time        `json:"updated_at"`
	DeletedAt          gorm.DeletedAt    `json:"deleted_at"`
}

type OauthRefreshTokenRepository interface {
	Create(entity OauthRefreshToken) (*OauthRefreshToken, *resp.ErrorResp)
	FindOneByToken(token string) (*OauthRefreshToken, *resp.ErrorResp)
	FindOneByRefreshToken(refreshToken string) (*OauthRefreshToken, *resp.ErrorResp)
	Delete(OauthRefreshToken) *resp.ErrorResp
}

// ------- Oauth Refresh Token
// ------- Request Body

type LoginRequestBody struct {
	Email        string `json:"email" binding:"email"`
	Password     string `json:"password" binding:"required"`
	ClientID     string `json:"client_id" binding:"required"`
	ClientSecret string `json:"client_secret" binding:"required"`
}

type RefreshTokenRequestBody struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

// ------- Request Body
// ------- Response Body

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Type         string `json:"Bearer"`
	ExpiredAt    string `json:"expired_at"`
	Scope        string `json:"scope"`
}

type UserResponse struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ClaimResponse struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	IsAdmin bool   `json:"is_admin,omitempty"`
	jwt.RegisteredClaims
}

type MapClaimResponse struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	IsAdmin string `json:"is_admin,omitempty"`
	jwt.MapClaims
}

// ------- Response Body
// ------- Oauth Usecase

type OauthUsecase interface {
	Login(data LoginRequestBody) (*LoginResponse, *resp.ErrorResp)
}
