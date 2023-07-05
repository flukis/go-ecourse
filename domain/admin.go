package domain

import (
	"e-course/pkg/resp"
	"time"

	"gorm.io/gorm"
)

type Admin struct {
	ID          int64          `json:"id"`
	Name        string         `json:"name"`
	Password    string         `json:"-"`
	Email       string         `json:"email"`
	CreatedAt   *time.Time     `json:"created_at"`
	UpdatedAt   *time.Time     `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
	CreatedByID *int64         `json:"created_by" gorm:"column:created_by"`
	CreatedBy   *Admin         `json:"-" gorm:"foreignKey:CreatedByID;references:ID"`
	UpdatedByID *int64         `json:"updated_by" gorm:"column:updated_by"`
	UpdatedBy   *Admin         `json:"-" gorm:"foreignKey:UpdatedByID;references:ID"`
}

type AdminRequestBody struct {
	Email     string  `json:"email" binding:"email"`
	Name      string  `json:"name" binding:"required"`
	Password  *string `json:"password"`
	CreatedBy *int64  `json:"created_by"`
	UpdatedBy *int64  `json:"updated_by"`
}

type AdminRepository interface {
	FindAll(offset, limit int) []Admin
	FindOneByID(id int) (*Admin, *resp.ErrorResp)
	FindOneByEmail(email string) (*Admin, *resp.ErrorResp)
	Create(Admin) (*Admin, *resp.ErrorResp)
	Update(Admin) (*Admin, *resp.ErrorResp)
	Delete(Admin) *resp.ErrorResp
	TotalCountAdmin() int64
}
