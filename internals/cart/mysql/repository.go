package repository

import (
	"e-course/domain"
	"e-course/pkg/resp"
	"e-course/pkg/utils"

	"gorm.io/gorm"
)

type mysqlCartRepository struct {
	db *gorm.DB
}

// DeleteByUserId implements domain.CartRepository.
func (r *mysqlCartRepository) DeleteByUserId(userId int) *resp.ErrorResp {
	var carts domain.Cart
	if err := r.db.Where("user_id = ?", userId).Delete(&carts).Error; err != nil {
		return &resp.ErrorResp{
			Code: 500,
			Err:  err,
		}
	}

	return nil
}

// FindAllByUserId implements domain.CartRepository.
func (r *mysqlCartRepository) FindAllByUserId(userId int, offset int, limit int) []domain.Cart {
	var carts []domain.Cart

	r.db.Scopes(utils.Paginate(offset, limit)).Preload("User").Preload("Product").Where("user_id = ?", userId).Find(&carts)

	return carts
}

// FindOneById implements domain.CartRepository.
func (r *mysqlCartRepository) FindOneById(id int) (*domain.Cart, *resp.ErrorResp) {
	var cart domain.Cart
	if err := r.db.Where("id = ?", id).Preload("User").Preload("Product").First(&cart).Error; err != nil {
		return nil, &resp.ErrorResp{
			Code: 500,
			Err:  err,
		}
	}
	return &cart, nil
}

// Create implements domain.CartRepository.
func (r *mysqlCartRepository) Create(data domain.Cart) (*domain.Cart, *resp.ErrorResp) {
	if err := r.db.Create(&data).Error; err != nil {
		return nil, &resp.ErrorResp{
			Code: 500,
			Err:  err,
		}
	}
	return &data, nil
}

// Delete implements domain.CartRepository.
func (r *mysqlCartRepository) Delete(data domain.Cart) *resp.ErrorResp {
	if err := r.db.Delete(&data).Error; err != nil {
		return &resp.ErrorResp{
			Code: 500,
			Err:  err,
		}
	}

	return nil
}

// Update implements domain.CartRepository.
func (r *mysqlCartRepository) Update(data domain.Cart) (*domain.Cart, *resp.ErrorResp) {
	if err := r.db.Save(&data).Error; err != nil {
		return nil, &resp.ErrorResp{
			Code: 500,
			Err:  err,
		}
	}
	return &data, nil
}

func NewMysqlCartRepository(db *gorm.DB) domain.CartRepository {
	return &mysqlCartRepository{db}
}
