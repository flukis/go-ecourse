package domain

import (
	"e-course/pkg/resp"
	"time"

	"gorm.io/gorm"
)

type Cart struct {
	ID          int64          `json:"id"`
	User        *User          `json:"user" gorm:"foreignKey:UserID;references:ID"`
	UserID      *int64         `json:"user_id"`
	Product     *Product       `json:"product" gorm:"foreingKey:ProductID;references:ID"`
	ProductID   *int64         `json:"product_id"`
	Quantity    int64          `json:"quantity"`
	IsChecked   bool           `json:"is_checked"`
	CreatedBy   *User          `json:"-" gorm:"foreignKey:CreatedByID;references:ID"`
	CreatedByID *int64         `json:"created_by" gorm:"column:created_by"`
	UpdatedByID *int64         `json:"updated_by" gorm:"column:updated_by"`
	UpdatedBy   *User          `json:"-" gorm:"foreignKey:UpdatedByID;references:ID"`
	CreatedAt   *time.Time     `json:"created_at"`
	UpdatedAt   *time.Time     `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}

type CartRequestBody struct {
	ProductID int64 `json:"product_id" binding:"required,number"`
	UserID    int64 `json:"user_id"`
	CreatedBy int64
	UpdatedBy int64
}

type CartRequestUpdateBody struct {
	IsChecked bool   `json:"is_checked"`
	UserID    *int64 `json:"user_id"`
}

type CartRepository interface {
	FindAllByUserId(userId int, offset int, limit int) []Cart
	FindOneById(id int) (*Cart, *resp.ErrorResp)
	Create(entity Cart) (*Cart, *resp.ErrorResp)
	Update(entity Cart) (*Cart, *resp.ErrorResp)
	Delete(entity Cart) *resp.ErrorResp
	DeleteByUserId(userId int) *resp.ErrorResp
}

type CartUsecase interface {
	FindAllByUserId(userId int, offset int, limit int) []Cart
	FindOneById(id int) (*Cart, *resp.ErrorResp)
	Create(data CartRequestBody) (*Cart, *resp.ErrorResp)
	Delete(id int, userId int) *resp.ErrorResp
	DeleteByUserId(userId int) *resp.ErrorResp
	Update(id int, data CartRequestUpdateBody) (*Cart, *resp.ErrorResp)
}
