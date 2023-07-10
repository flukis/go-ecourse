package oauth

import (
	"e-course/domain"
	"e-course/pkg/resp"
	"e-course/pkg/utils"
	"errors"
	"fmt"
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
	adminUsecase domain.AdminUsecase
}

// Refresh implements domain.OauthUsecase.
func (uc *oAuthUsecase) Refresh(data domain.RefreshTokenRequestBody) (*domain.LoginResponse, *resp.ErrorResp) {
	// check oaut refresh token
	oauthRefreshToken, err := uc.refreshToken.FindOneByToken(data.RefreshToken)
	if err != nil {
		return nil, err
	}

	if oauthRefreshToken.ExpiredAt.Before(time.Now()) {
		return nil, &resp.ErrorResp{
			Code: 400,
			Err:  errors.New("refresh token is expired, please login"),
		}
	}

	var user domain.UserResponse
	isAdmin := false
	expTime := time.Now().Add(24 * 365 * time.Hour)

	if *oauthRefreshToken.OauthAccessToken.OauthClientID == 2 {
		admin, _ := uc.adminUsecase.FindOneByID(int(oauthRefreshToken.UserID))

		user.ID = admin.ID
		user.Email = admin.Email
		user.Name = admin.Name
		isAdmin = true
	} else {
		commonUser, _ := uc.userUsecase.FindOneByID(int(oauthRefreshToken.UserID))

		user.ID = commonUser.ID
		user.Email = commonUser.Email
		user.Name = commonUser.Name
	}

	claims := &domain.ClaimResponse{
		ID:      user.ID,
		Name:    user.Name,
		Email:   user.Email,
		IsAdmin: isAdmin,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	jwtKey := []byte(os.Getenv("JWT_SECRET"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, errSign := token.SignedString(jwtKey)

	if errSign != nil {
		return nil, &resp.ErrorResp{
			Code: 500,
			Err:  errSign,
		}
	}

	// insert ke oauth access token
	newDataOauthAccessToken := domain.OauthAccessToken{
		OauthClientID: oauthRefreshToken.OauthAccessToken.OauthClientID,
		UserID:        oauthRefreshToken.ID,
		Token:         tokenString,
		Scope:         "*",
		ExpiredAt:     &expTime,
	}

	saveOauthAccessToken, err := uc.accessToken.Create(newDataOauthAccessToken)
	if err != nil {
		return nil, err
	}

	expTimeRefreshToken := time.Now().Add(24 * 266 * time.Hour)

	refreshTokenString, errGen := utils.GenerateRefreshToken()
	if errGen != nil {
		return nil, &resp.ErrorResp{
			Code: 500,
			Err:  errors.New("internal server error"),
		}
	}
	dataOauthRefreshToken := domain.OauthRefreshToken{
		OauthAccessTokenID: &saveOauthAccessToken.ID,
		UserID:             user.ID,
		Token:              refreshTokenString.String(),
		ExpiredAt:          &expTimeRefreshToken,
	}

	saveOauthRefreshToken, err := uc.refreshToken.Create(dataOauthRefreshToken)
	if err != nil {
		return nil, err
	}

	// delete old refresh token
	err = uc.refreshToken.Delete(*oauthRefreshToken)
	if err != nil {
		return nil, err
	}

	// delete old access token
	err = uc.accessToken.Delete(*oauthRefreshToken.OauthAccessToken)
	if err != nil {
		return nil, err
	}

	return &domain.LoginResponse{
		AccessToken:  tokenString,
		RefreshToken: saveOauthRefreshToken.Token,
		Type:         "Bearer",
		ExpiredAt:    expTime.Format(time.RFC3339),
		Scope:        "*",
	}, err
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

	// check if admin
	if oauthClient.Name == "web-admin" {
		dataEmail, err := uc.adminUsecase.FindOneByEmail(data.Email)
		if err != nil {
			return nil, &resp.ErrorResp{
				Code: 400,
				Err:  errors.New("admin with that email is not found or invalid"),
			}
		}

		user.ID = dataEmail.ID
		user.Email = dataEmail.Email
		user.Name = dataEmail.Name
		user.Password = dataEmail.Password
	} else {
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
	}

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

	if oauthClient.ClientID == "2" {
		claims.IsAdmin = true
	}

	fmt.Println(claims)

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
	refreshTokenString, errGen := utils.GenerateRefreshToken()
	if errGen != nil {
		return nil, &resp.ErrorResp{
			Code: 500,
			Err:  errors.New("internal server error"),
		}
	}
	dataRefreshToken := domain.OauthRefreshToken{
		OauthAccessTokenID: &resAccessToken.ID,
		UserID:             user.ID,
		Token:              refreshTokenString.String(),
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
	adminUsecase domain.AdminUsecase,
) domain.OauthUsecase {
	return &oAuthUsecase{
		client:       client,
		accessToken:  accessToken,
		refreshToken: refreshToken,
		userUsecase:  userUsecase,
		adminUsecase: adminUsecase,
	}
}
