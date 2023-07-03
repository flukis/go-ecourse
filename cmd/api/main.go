package main

import (
	"e-course/internals/register"
	userHTTPHandler "e-course/internals/register/http"
	userUC "e-course/internals/user"
	userRepo "e-course/internals/user/mysql"
	mysql "e-course/pkg/db/mysql"

	oauthUC "e-course/internals/oauth"
	oauthHTTPHandler "e-course/internals/oauth/http"
	oauthRepo "e-course/internals/oauth/mysql"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := mysql.DB()

	userRepository := userRepo.NewMysqlUserRepository(db)
	userUsecase := userUC.NewUserUsacase(userRepository)

	registerUsecase := register.NewRegisterUsecase(userUsecase)

	userHTTPHandler.NewRegisterHandler(registerUsecase).Route(
		&r.RouterGroup,
	)

	oauthClientRepository := oauthRepo.NewOAuthClientRepository(db)
	oauthAccessRepository := oauthRepo.NewOauthAccessTokenRepository(db)
	oauthRefreshRepository := oauthRepo.NewOauthRefreshTokenRepository(db)

	oauthUseCase := oauthUC.NewOAuthUsecase(
		oauthClientRepository,
		oauthAccessRepository,
		oauthRefreshRepository,
		userUsecase,
	)

	oauthHTTPHandler.NewOAuthHandler(oauthUseCase).Route(
		&r.RouterGroup,
	)

	r.Run(":8081")
}
