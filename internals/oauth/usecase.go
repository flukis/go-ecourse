package oauth

import (
	"e-course/domain"
	"e-course/pkg/resp"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type oAuthUsecase struct {
	client       domain.OauthClientRepository
	accessToken  domain.OauthAccessTokenRepository
	refreshToken domain.OauthRefreshTokenRepository
	userUsecase  domain.UserUsecase
}

// Login implements domain.OAuthUsecase.
func (uc *oAuthUsecase) Login(data domain.LoginRequestBody) (*domain.LoginResponse, *resp.ErrorResp) {
	// TODO: CHECK IS ADMIN
	// check client id and client secret is valid
	oauthClient, err := uc.client.FindByClientIDnClientSecret(
		data.ClientID,
		data.ClientSecret,
	)

	if err != nil {
		return nil, err
	}

	var user domain.UserResponse
	dataUser, err := uc.userUsecase.FindByEmail(data.Email)
	if err != nil {
		return nil, &resp.ErrorResp{
			Code: 400,
			Err:  errors.New("user with that email is not found or invalid"),
		}
	}

	user.ID = dataUser.ID
	user.Email = dataUser.Email
	user.Name = dataUser.Name
	user.Password = dataUser.Password

	jwtKey := []byte(os.Getenv("JWT_SECRET"))
	errBcrypt := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(data.Password),
	)
	if errBcrypt != nil {
		return nil, &resp.ErrorResp{
			Code: 401,
			Err:  errors.New("wrong password"),
		}
	}

	// create expired for JWT
	expTime := time.Now().Add(24 * 365 * time.Hour)
	claims := &domain.ClaimResponse{
		ID:      user.ID,
		Name:    user.Name,
		Email:   user.Name,
		IsAdmin: false,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	tokenString, _ := token.SignedString(jwtKey)

	// insert data oauth access token
	dataAccesToken := domain.OauthAccessToken{
		OauthClientID: &oauthClient.ID,
		UserID:        user.ID,
		Token:         tokenString,
		Scope:         "*",
		ExpiredAt:     &expTime,
	}

	resAccessToken, err := uc.accessToken.Create(dataAccesToken)
	if err != nil {
		return nil, err
	}
	// insert data refresh access token
	dataRefreshToken := domain.OauthRefreshToken{
		OauthAccessTokenID: &oauthClient.ID,
		UserID:             user.ID,
		Token:              tokenString,
		ExpiredAt:          &expTime,
	}
	resRefreshToken, err := uc.refreshToken.Create(dataRefreshToken)
	if err != nil {
		return nil, err
	}

	return &domain.LoginResponse{
		AccessToken:  resAccessToken.Token,
		RefreshToken: resRefreshToken.Token,
		Type:         "Bearer",
		ExpiredAt:    expTime.Format(time.RFC3339),
		Scope:        "*",
	}, nil
}

func NewOAuthUsecase(
	client domain.OauthClientRepository,
	accessToken domain.OauthAccessTokenRepository,
	refreshToken domain.OauthRefreshTokenRepository,
	userUsecase domain.UserUsecase,
) domain.OauthUsecase {
	return &oAuthUsecase{
		client:       client,
		accessToken:  accessToken,
		refreshToken: refreshToken,
		userUsecase:  userUsecase,
	}
}
