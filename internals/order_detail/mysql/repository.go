package repository

import (
	"e-course/domain"
	"e-course/pkg/resp"

	"gorm.io/gorm"
)

type orderDetailRepository struct {
	db *gorm.DB
}

// Create implements domain.OrderDetailRepository.
func (r *orderDetailRepository) Create(entity domain.OrderDetail) (*domain.OrderDetail, *resp.ErrorResp) {
	if err := r.db.Create(&entity).Error; err != nil {
		return nil, &resp.ErrorResp{
			Code: 500,
			Err:  err,
		}
	}
	return &entity, nil
}

func NewOrderDetailRepository(db *gorm.DB) domain.OrderDetailRepository {
	return &orderDetailRepository{db}
}
