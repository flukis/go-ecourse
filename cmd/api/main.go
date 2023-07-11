package main

import (
	adminUC "e-course/internals/admin"
	adminHTTPHandler "e-course/internals/admin/http"
	adminRepo "e-course/internals/admin/mysql"
	forgotPasswordUC "e-course/internals/forgot_password"
	forgotPasswordHTTPHandler "e-course/internals/forgot_password/http"
	forgotPasswordRepo "e-course/internals/forgot_password/mysql"
	oauthUC "e-course/internals/oauth"
	oauthHTTPHandler "e-course/internals/oauth/http"
	oauthRepo "e-course/internals/oauth/mysql"
	productUC "e-course/internals/product"
	productHandler "e-course/internals/product/http"
	productRepo "e-course/internals/product/mysql"
	productCategoryUC "e-course/internals/product_category"
	productCategoryHandler "e-course/internals/product_category/http"
	productCategoryRepo "e-course/internals/product_category/mysql"
	"e-course/internals/register"
	userHTTPHandler "e-course/internals/register/http"
	userUC "e-course/internals/user"
	userRepo "e-course/internals/user/mysql"
	mysql "e-course/pkg/db/mysql"
	email "e-course/pkg/mail/sendgrid"
	media "e-course/pkg/media/cloudinary"

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
	adminRepository := adminRepo.NewMysqlAdminRepository(db)
	productCategoryRepository := productCategoryRepo.NewMysqlProductCategoryRepository(db)
	productRepository := productRepo.NewMysqlProductRepository(db)

	mediaUsecase := media.NewMediaUsecase()
	mailUsecase := email.NewMailUsecase()
	userUsecase := userUC.NewUserUsacase(userRepository)
	registerUsecase := register.NewRegisterUsecase(userUsecase, mailUsecase)
	adminUsecase := adminUC.NewAdminUsecase(adminRepository)
	forgotpasswordUsecase := forgotPasswordUC.NewForgotPasswordUsecase(forgotpasswordRepository, userUsecase, mailUsecase)
	oauthUseCase := oauthUC.NewOAuthUsecase(
		oauthClientRepository,
		oauthAccessRepository,
		oauthRefreshRepository,
		userUsecase,
		adminUsecase,
	)
	productCategoryUsecase := productCategoryUC.NewProductCategoryUsecase(productCategoryRepository, mediaUsecase)
	productUsecase := productUC.NewProductUsecase(productRepository, mediaUsecase)

	oauthHTTPHandler.NewOAuthHandler(oauthUseCase).Route(
		&r.RouterGroup,
	)
	userHTTPHandler.NewRegisterHandler(registerUsecase).Route(
		&r.RouterGroup,
	)
	forgotPasswordHTTPHandler.NewForgotPasswordhandler(forgotpasswordUsecase).Route(
		&r.RouterGroup,
	)
	adminHTTPHandler.NewAdminHandler(adminUsecase).Route(
		&r.RouterGroup,
	)
	productCategoryHandler.NewProductCategoryHandler(productCategoryUsecase).Route(
		&r.RouterGroup,
	)
	productHandler.NewProductHandler(productUsecase).Route(
		&r.RouterGroup,
	)

	r.Run(":8081")
}
