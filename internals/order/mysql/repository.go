package repository

import (
	"e-course/domain"
	"e-course/pkg/resp"
	"e-course/pkg/utils"

	"gorm.io/gorm"
)

type orderRepository struct {
	db *gorm.DB
}

// Create implements domain.OrderRepository.
func (r *orderRepository) Create(data domain.Order) (*domain.Order, *resp.ErrorResp) {
	if err := r.db.Create(&data).Error; err != nil {
		return nil, &resp.ErrorResp{
			Code: 500,
			Err:  err,
		}
	}
	return &data, nil
}

// FindAllByUserId implements domain.OrderRepository.
func (r *orderRepository) FindAllByUserId(userId int, offset int, limit int) []domain.Order {
	var orders []domain.Order

	r.db.Scopes(utils.Paginate(offset, limit)).Preload("OrderDetails.Product").Where("user_id = ?", userId).Find(&orders)

	return orders
}

// FindOneByExternalId implements domain.OrderRepository.
func (r *orderRepository) FindOneByExternalId(externalId string) (*domain.Order, *resp.ErrorResp) {
	var order domain.Order
	if err := r.db.Where("external_id = ?", externalId).Preload("OrderDetails.Product").First(&order).Error; err != nil {
		return nil, &resp.ErrorResp{
			Code: 500,
			Err:  err,
		}
	}
	return &order, nil
}

// FindOneById implements domain.OrderRepository.
func (r *orderRepository) FindOneById(id int) (*domain.Order, *resp.ErrorResp) {
	var order domain.Order
	if err := r.db.Where("id = ?", id).Preload("OrderDetails.Product").First(&order).Error; err != nil {
		return nil, &resp.ErrorResp{
			Code: 500,
			Err:  err,
		}
	}
	return &order, nil
}

// TotalCountOrder implements domain.OrderRepository.
func (r *orderRepository) TotalCountOrder() int64 {
	var order domain.Order
	var totalOrder int64
	r.db.Model(&order).Count(&totalOrder)
	return totalOrder
}

// Update implements domain.OrderRepository.
func (r *orderRepository) Update(data domain.Order) (*domain.Order, *resp.ErrorResp) {
	if err := r.db.Save(&data).Error; err != nil {
		return nil, &resp.ErrorResp{
			Code: 500,
			Err:  err,
		}
	}
	return &data, nil
}

func NewOrderRepository(db *gorm.DB) domain.OrderRepository {
	return &orderRepository{db}
}
