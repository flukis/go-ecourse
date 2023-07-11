package mysql

import (
	"e-course/domain"
	"e-course/pkg/resp"
	"e-course/pkg/utils"

	"gorm.io/gorm"
)

type mysqlProductRepository struct {
	db *gorm.DB
}

// Create implements domain.ProductRepository.
func (m *mysqlProductRepository) Create(entity domain.Product) (*domain.Product, *resp.ErrorResp) {
	if err := m.db.Create(&entity).Error; err != nil {
		return nil, &resp.ErrorResp{
			Code: 500,
			Err:  err,
		}
	}
	return &entity, nil
}

// Delete implements domain.ProductRepository.
func (m *mysqlProductRepository) Delete(entity domain.Product) *resp.ErrorResp {
	if err := m.db.Delete(&entity).Error; err != nil {
		return &resp.ErrorResp{
			Code: 500,
			Err:  err,
		}
	}
	return nil
}

// FindAll implements domain.ProductRepository.
func (m *mysqlProductRepository) FindAll(offset int, limit int) []domain.Product {
	var entity []domain.Product
	m.db.Scopes(utils.Paginate(offset, limit)).Find(&entity)
	return entity
}

// FindOneById implements domain.ProductRepository.
func (m *mysqlProductRepository) FindOneById(id int) (*domain.Product, *resp.ErrorResp) {
	var entity domain.Product
	if err := m.db.Where("id = ?", id).First(&entity).Error; err != nil {
		return nil, &resp.ErrorResp{
			Code: 500,
			Err:  err,
		}
	}
	return &entity, nil
}

// TotalCountProduct implements domain.ProductRepository.
func (m *mysqlProductRepository) TotalCountProduct() int64 {
	var product domain.Product
	var totalProduct int64
	m.db.Model(&product).Count(&totalProduct)
	return totalProduct
}

// Update implements domain.ProductRepository.
func (m *mysqlProductRepository) Update(entity domain.Product) (*domain.Product, *resp.ErrorResp) {
	if err := m.db.Save(&entity).Error; err != nil {
		return nil, &resp.ErrorResp{
			Code: 500,
			Err:  err,
		}
	}
	return &entity, nil
}

func NewMysqlProductRepository(db *gorm.DB) domain.ProductRepository {
	return &mysqlProductRepository{db: db}
}
