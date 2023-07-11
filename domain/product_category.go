package domain

import (
	"e-course/pkg/resp"
	"mime/multipart"
	"time"

	"gorm.io/gorm"
)

type ProductCategory struct {
	ID          int64          `json:"id"`
	Name        string         `json:"name"`
	Image       *string        `json:"image"`
	CreatedByID *int64         `json:"created_by" gorm:"column:created_by"`
	CreatedBy   *Admin         `json:"-" gorm:"foreignKey:CreatedByID;references:ID"`
	UpdatedByID *int64         `json:"updated_by" gorm:"column:updated_by"`
	UpdatedBy   *Admin         `json:"-" gorm:"foreignKey:UpdatedByID;references:ID"`
	CreatedAt   *time.Time     `json:"created_at"`
	UpdatedAt   *time.Time     `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}

type ProductCategoryRepository interface {
	FindAll(offset, limit int) []ProductCategory
	FindOneByID(id int) (*ProductCategory, *resp.ErrorResp)
	Create(data ProductCategory) (*ProductCategory, *resp.ErrorResp)
	Update(data ProductCategory) (*ProductCategory, *resp.ErrorResp)
	Delete(data ProductCategory) *resp.ErrorResp
}

type ProductCategoryUsecase interface {
	FindAll(offset, limit int) []ProductCategory
	FindOneByID(id int) (*ProductCategory, *resp.ErrorResp)
	Create(data ProductCategoryRequestBody) (*ProductCategory, *resp.ErrorResp)
	Update(id int, data ProductCategoryRequestBody) (*ProductCategory, *resp.ErrorResp)
	Delete(id int) *resp.ErrorResp
}

type ProductCategoryRequestBody struct {
	Name      string                `form:"name" binding:"required"`
	Image     *multipart.FileHeader `form:"image"`
	CreatedBy *int64
	UpdatedBy *int64
}
