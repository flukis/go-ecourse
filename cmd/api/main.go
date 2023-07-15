package main

import (
	adminUC "e-course/internals/admin"
	adminHTTPHandler "e-course/internals/admin/http"
	adminRepo "e-course/internals/admin/mysql"
	cartUC "e-course/internals/cart"
	cartHandler "e-course/internals/cart/http"
	cartRepo "e-course/internals/cart/mysql"
	classRoomUC "e-course/internals/class_room"
	classRoomHandler "e-course/internals/class_room/http"
	classRoomRepo "e-course/internals/class_room/mysql"
	discountUC "e-course/internals/discount"
	discountHandler "e-course/internals/discount/http"
	discountRepo "e-course/internals/discount/mysql"
	forgotPasswordUC "e-course/internals/forgot_password"
	forgotPasswordHTTPHandler "e-course/internals/forgot_password/http"
	forgotPasswordRepo "e-course/internals/forgot_password/mysql"
	oauthUC "e-course/internals/oauth"
	oauthHTTPHandler "e-course/internals/oauth/http"
	oauthRepo "e-course/internals/oauth/mysql"
	orderUC "e-course/internals/order"
	orderHandler "e-course/internals/order/http"
	orderRepo "e-course/internals/order/mysql"
	orderDetailUC "e-course/internals/order_detail"
	orderDetailRepo "e-course/internals/order_detail/mysql"
	"e-course/internals/payment"
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
	discountRepository := discountRepo.NewMysqlDiscountRepository(db)
	cartRepository := cartRepo.NewMysqlCartRepository(db)
	orderRepo := orderRepo.NewOrderRepository(db)
	classroomRepo := classRoomRepo.NewClassroomRepository(db)
	orderDetailRepo := orderDetailRepo.NewOrderDetailRepository(db)

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
	paymentUsecase := payment.NewPaymentUsecase()
	classroomUsecase := classRoomUC.NewClassroomUsecase(classroomRepo)
	productCategoryUsecase := productCategoryUC.NewProductCategoryUsecase(productCategoryRepository, mediaUsecase)
	productUsecase := productUC.NewProductUsecase(productRepository, mediaUsecase)
	discountUsecase := discountUC.NewDiscountUsecase(discountRepository, mediaUsecase)
	cartUsecase := cartUC.NewCartUsecase(cartRepository)
	orderDetailUsecase := orderDetailUC.NewOrderDetailUsecase(orderDetailRepo)
	orderUsecase := orderUC.NewOrderUsecase(orderRepo, cartUsecase, discountUsecase, orderDetailUsecase, paymentUsecase, productUsecase)

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
	discountHandler.NewDiscountHandler(discountUsecase).Route(
		&r.RouterGroup,
	)
	cartHandler.NewCartHandler(cartUsecase).Route(
		&r.RouterGroup,
	)
	orderHandler.NewOrderHandler(orderUsecase).Route(
		&r.RouterGroup,
	)
	classRoomHandler.NewClassroomHandler(classroomUsecase).Route(
		&r.RouterGroup,
	)

	r.Run(":8081")
}
