package repository

import (
	"e-course/domain"
	"e-course/pkg/resp"
	"e-course/pkg/utils"

	"gorm.io/gorm"
)

type mysqlProductCategoryRepository struct {
	db *gorm.DB
}

// Create implements domain.ProductCategoryRepository.
func (m *mysqlProductCategoryRepository) Create(data domain.ProductCategory) (*domain.ProductCategory, *resp.ErrorResp) {
	if err := m.db.Create(&data).Error; err != nil {
		return nil, &resp.ErrorResp{
			Code: 500,
			Err:  err,
		}
	}
	return &data, nil
}

// Delete implements domain.ProductCategoryRepository.
func (m *mysqlProductCategoryRepository) Delete(data domain.ProductCategory) *resp.ErrorResp {
	if err := m.db.Delete(&data).Error; err != nil {
		return &resp.ErrorResp{
			Code: 500,
			Err:  err,
		}
	}
	return nil
}

// FindAll implements domain.ProductCategoryRepository.
func (m *mysqlProductCategoryRepository) FindAll(offset int, limit int) []domain.ProductCategory {
	var res []domain.ProductCategory
	m.db.Scopes(utils.Paginate(offset, limit)).Find(&res)
	return res
}

// FindOneByID implements domain.ProductCategoryRepository.
func (m *mysqlProductCategoryRepository) FindOneByID(id int) (*domain.ProductCategory, *resp.ErrorResp) {
	var res domain.ProductCategory
	if err := m.db.Where("id = ?", id).First(&res).Error; err != nil {
		return nil, &resp.ErrorResp{
			Code: 500,
			Err:  err,
		}
	}
	return &res, nil
}

// Update implements domain.ProductCategoryRepository.
func (m *mysqlProductCategoryRepository) Update(data domain.ProductCategory) (*domain.ProductCategory, *resp.ErrorResp) {
	if err := m.db.Save(&data).Error; err != nil {
		return nil, &resp.ErrorResp{
			Code: 500,
			Err:  err,
		}
	}
	return &data, nil
}

func NewMysqlProductCategoryRepository(db *gorm.DB) domain.ProductCategoryRepository {
	return &mysqlProductCategoryRepository{db: db}
}
