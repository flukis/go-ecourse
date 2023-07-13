package orderdetail

import (
	"e-course/domain"
	"fmt"
)

type orderDetailUsecase struct {
	repo domain.OrderDetailRepository
}

// Create implements domain.OrderDetailUsecase.
func (uc *orderDetailUsecase) Create(entity domain.OrderDetail) {
	_, err := uc.repo.Create(entity)
	if err != nil {
		fmt.Println(err)
	}
}

func NewOrderDetailUsecase(repo domain.OrderDetailRepository) domain.OrderDetailUsecase {
	return &orderDetailUsecase{repo}
}
