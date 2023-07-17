package dashboard

import "e-course/domain"

type dashboardUsecase struct {
	user    domain.UserUsecase
	admin   domain.AdminUsecase
	product domain.ProductUsecase
	order   domain.OrderUsecase
}

// GetDashboard implements domain.DashboardUsecase.
func (uc *dashboardUsecase) GetDashboard() domain.DashboardResponseBody {
	var data domain.DashboardResponseBody

	data.TotalAdmin = uc.admin.TotalCountAdmin()
	data.TotalUser = uc.user.TotalCountUser()
	data.TotalProduct = uc.product.TotalCountProduct()
	data.TotalOrder = uc.order.TotalCountOrder()

	return data
}

func NewDasboardUsecase(
	user domain.UserUsecase,
	admin domain.AdminUsecase,
	product domain.ProductUsecase,
	order domain.OrderUsecase) domain.DashboardUsecase {
	return &dashboardUsecase{user, admin, product, order}
}
