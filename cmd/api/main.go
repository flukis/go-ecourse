package main

import (
	forgotPasswordUC "e-course/internals/forgot_password"
	forgotPasswordHTTPHandler "e-course/internals/forgot_password/http"
	forgotPasswordRepo "e-course/internals/forgot_password/mysql"
	oauthUC "e-course/internals/oauth"
	oauthHTTPHandler "e-course/internals/oauth/http"
	oauthRepo "e-course/internals/oauth/mysql"
	"e-course/internals/register"
	userHTTPHandler "e-course/internals/register/http"
	userUC "e-course/internals/user"
	userRepo "e-course/internals/user/mysql"
	mysql "e-course/pkg/db/mysql"
	email "e-course/pkg/mail/sendgrid"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := mysql.DB()

	userRepository := userRepo.NewMysqlUserRepository(db)
	oauthClientRepository := oauthRepo.NewOAuthClientRepository(db)
	oauthAccessRepository := oauthRepo.NewOauthAccessTokenRepository(db)
	oauthRefreshRepository := oauthRepo.NewOauthRefreshTokenRepository(db)
	forgotpasswordRepository := forgotPasswordRepo.NewMysqlForgotPasswordRepository(db)

	mailUsecase := email.NewMailUsecase()
	userUsecase := userUC.NewUserUsacase(userRepository)
	registerUsecase := register.NewRegisterUsecase(userUsecase, mailUsecase)
	oauthUseCase := oauthUC.NewOAuthUsecase(
		oauthClientRepository,
		oauthAccessRepository,
		oauthRefreshRepository,
		userUsecase,
	)
	forgotpasswordUsecase := forgotPasswordUC.NewForgotPasswordUsecase(forgotpasswordRepository, userUsecase, mailUsecase)

	oauthHTTPHandler.NewOAuthHandler(oauthUseCase).Route(
		&r.RouterGroup,
	)
	userHTTPHandler.NewRegisterHandler(registerUsecase).Route(
		&r.RouterGroup,
	)
	forgotPasswordHTTPHandler.NewForgotPasswordhandler(forgotpasswordUsecase).Route(
		&r.RouterGroup,
	)

	r.Run(":8081")
}
