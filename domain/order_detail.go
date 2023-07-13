package domain

import (
	"e-course/pkg/resp"
	"time"

	"gorm.io/gorm"
)

type OrderDetail struct {
	ID          int64          `json:"id"`
	OrderID     int64          `json:"order_id"`
	Product     *Product       `json:"product" gorm:"foreignKey:ProductID;references:ID"`
	ProductID   *int64         `json:"product_id"`
	Price       int64          `json:"price"`
	CreatedBy   *User          `json:"-" gorm:"foreignKey:CreatedByID;references:ID"`
	CreatedByID *int64         `json:"created_by" gorm:"column:created_by"`
	UpdatedByID *int64         `json:"updated_by" gorm:"column:updated_by"`
	UpdatedBy   *User          `json:"-" gorm:"foreignKey:UpdatedByID;references:ID"`
	CreatedAt   *time.Time     `json:"created_at"`
	UpdatedAt   *time.Time     `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}

type OrderDetailRepository interface {
	Create(entity OrderDetail) (*OrderDetail, *resp.ErrorResp)
}

type OrderDetailUsecase interface {
	Create(entity OrderDetail)
}
