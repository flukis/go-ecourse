package domain

import (
	"e-course/pkg/resp"
	"time"

	"gorm.io/gorm"
)

type Discount struct {
	ID                int64          `json:"id"`
	Name              string         `json:"name"`
	Code              string         `json:"code"`
	Quantity          int64          `json:"quantity"`
	RemainingQuantity int64          `json:"remaining_quantity"`
	Type              string         `json:"type"`
	Value             int64          `json:"value"`
	StartDate         *time.Time     `json:"start_date"`
	EndDate           *time.Time     `json:"end_date"`
	CreatedByID       *int64         `json:"created_by" gorm:"column:created_by"`
	CreatedBy         *Admin         `json:"-" gorm:"foreignKey:CreatedByID;references:ID"`
	UpdatedByID       *int64         `json:"updated_by" gorm:"column:updated_by"`
	UpdatedBy         *Admin         `json:"-" gorm:"foreignKey:UpdatedByID;references:ID"`
	CreatedAt         *time.Time     `json:"created_at"`
	UpdatedAt         *time.Time     `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `json:"deleted_at"`
}

type DiscountRequestBody struct {
	Name      string     `json:"name" binding:"required"`
	Code      string     `json:"code" binding:"required"`
	Quantity  int64      `json:"quantity" binding:"required"`
	Type      string     `json:"type" binding:"required"`
	Value     int64      `json:"value" binding:"required"`
	StartDate *time.Time `json:"start_date"`
	EndDate   *time.Time `json:"end_date"`
	CreatedBy *int64
	UpdatedBy *int64
}

type DiscountRepository interface {
	FindAll(offset int, limit int) []Discount
	FindOneById(id int) (*Discount, *resp.ErrorResp)
	FindOneByCode(code string) (*Discount, *resp.ErrorResp)
	Create(Discount) (*Discount, *resp.ErrorResp)
	Update(Discount) (*Discount, *resp.ErrorResp)
	Delete(Discount) *resp.ErrorResp
}

type DiscountUsecase interface {
	FindAll(offset int, limit int) []Discount
	FindOneById(id int) (*Discount, *resp.ErrorResp)
	FindOneByCode(code string) (*Discount, *resp.ErrorResp)
	Create(data DiscountRequestBody) (*Discount, *resp.ErrorResp)
	Update(id int, data DiscountRequestBody) (*Discount, *resp.ErrorResp)
	Delete(id int) *resp.ErrorResp
	UpdateRemainingQuantity(id int, quantity int, operator string) (*Discount, *resp.ErrorResp)
}
