package mysql

import (
	"e-course/domain"
	"e-course/pkg/resp"
	"e-course/pkg/utils"

	"gorm.io/gorm"
)

type mysqlDiscountRepository struct {
	db *gorm.DB
}

// FindOneByCode implements domain.DiscountRepository.
func (m *mysqlDiscountRepository) FindOneByCode(code string) (*domain.Discount, *resp.ErrorResp) {
	var entity domain.Discount
	if err := m.db.Where("code = ?", code).First(&entity).Error; err != nil {
		return nil, &resp.ErrorResp{
			Code: 500,
			Err:  err,
		}
	}
	return &entity, nil
}

// Create implements domain.DiscountRepository.
func (m *mysqlDiscountRepository) Create(entity domain.Discount) (*domain.Discount, *resp.ErrorResp) {
	if err := m.db.Create(&entity).Error; err != nil {
		return nil, &resp.ErrorResp{
			Code: 500,
			Err:  err,
		}
	}
	return &entity, nil
}

// Delete implements domain.DiscountRepository.
func (m *mysqlDiscountRepository) Delete(entity domain.Discount) *resp.ErrorResp {
	if err := m.db.Delete(&entity).Error; err != nil {
		return &resp.ErrorResp{
			Code: 500,
			Err:  err,
		}
	}
	return nil
}

// FindAll implements domain.DiscountRepository.
func (m *mysqlDiscountRepository) FindAll(offset int, limit int) []domain.Discount {
	var entity []domain.Discount
	m.db.Scopes(utils.Paginate(offset, limit)).Find(&entity)
	return entity
}

// FindOneById implements domain.DiscountRepository.
func (m *mysqlDiscountRepository) FindOneById(id int) (*domain.Discount, *resp.ErrorResp) {
	var entity domain.Discount
	if err := m.db.Where("id = ?", id).First(&entity).Error; err != nil {
		return nil, &resp.ErrorResp{
			Code: 500,
			Err:  err,
		}
	}
	return &entity, nil
}

// TotalCountDiscount implements domain.DiscountRepository.
func (m *mysqlDiscountRepository) TotalCountDiscount() int64 {
	var Discount domain.Discount
	var totalDiscount int64
	m.db.Model(&Discount).Count(&totalDiscount)
	return totalDiscount
}

// Update implements domain.DiscountRepository.
func (m *mysqlDiscountRepository) Update(entity domain.Discount) (*domain.Discount, *resp.ErrorResp) {
	if err := m.db.Save(&entity).Error; err != nil {
		return nil, &resp.ErrorResp{
			Code: 500,
			Err:  err,
		}
	}
	return &entity, nil
}

func NewMysqlDiscountRepository(db *gorm.DB) domain.DiscountRepository {
	return &mysqlDiscountRepository{db: db}
}
