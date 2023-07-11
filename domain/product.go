package domain

import (
	"e-course/pkg/resp"
	"mime/multipart"
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID                int64            `json:"id"`
	ProductCategoryID *int64           `json:"product_category_id"`
	ProductCategory   *ProductCategory `json:"product_category" gorm:"foreingKey:ProductCategoryID;references:ID"`
	Title             string           `json:"title"`
	Image             *string          `json:"image"`
	Video             *string          `json:"-"`
	VideoLink         *string          `json:"video,omitempty" gorm:"-"`
	Description       string           `json:"description"`
	IsHighlighted     bool             `json:"is_highlighted"`
	Price             int64            `json:"price"`
	CreatedByID       *int64           `json:"created_by" gorm:"column:created_by"`
	CreatedBy         *Admin           `json:"-" gorm:"foreignKey:CreatedByID;references:ID"`
	UpdatedByID       *int64           `json:"updated_by" gorm:"column:updated_by"`
	UpdatedBy         *Admin           `json:"-" gorm:"foreignKey:UpdatedByID;references:ID"`
	CreatedAt         *time.Time       `json:"created_at"`
	UpdatedAt         *time.Time       `json:"updated_at"`
	DeletedAt         gorm.DeletedAt   `json:"deleted_at"`
}

type ProductRequestBody struct {
	ProductCategoryID int64                 `form:"product_category_id" binding:"required"`
	Title             string                `form:"title" binding:"required"`
	Image             *multipart.FileHeader `form:"image"`
	Video             *multipart.FileHeader `form:"video"`
	Description       string                `form:"description" binding:"required"`
	IsHighlighted     bool                  `form:"is_highlighted" default:"false"`
	Price             int                   `form:"price" binding:"required"`
	CreatedBy         *int64
	UpdatedBy         *int64
}

type ProductRepository interface {
	FindAll(offset int, limit int) []Product
	FindOneById(id int) (*Product, *resp.ErrorResp)
	Create(entity Product) (*Product, *resp.ErrorResp)
	Update(entity Product) (*Product, *resp.ErrorResp)
	Delete(entity Product) *resp.ErrorResp
	TotalCountProduct() int64
}

type ProductUsecase interface {
	FindAll(offset int, limit int) []Product
	FindOneById(id int) (*Product, *resp.ErrorResp)
	Create(dto ProductRequestBody) (*Product, *resp.ErrorResp)
	Update(id int, dto ProductRequestBody) (*Product, *resp.ErrorResp)
	Delete(id int) *resp.ErrorResp
	TotalCountProduct() int64
}
