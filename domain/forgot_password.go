package domain

import (
	"e-course/pkg/resp"
	"time"

	"gorm.io/gorm"
)

type ForgotPassword struct {
	ID        int64          `json:"id"`
	User      *User          `json:"user" gorm:"foreignKey:UserID;reference:ID"`
	UserID    *int64         `json:"user_id"`
	Valid     bool           `json:"valid"`
	Code      string         `json:"code"`
	ExpiredAt *time.Time     `json:"expired_at"`
	CreatedAt *time.Time     `json:"created_at"`
	UpdatedAt *time.Time     `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type ForgotPasswordRequestBody struct {
	Email string `json:"email" binding:"email"`
}

type ForgotPasswordUpdateRequestBody struct {
	Code     string `json:"code" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ForgotPasswordEmailRequestBody struct {
	UserID    string     `json:"user_id"`
	Valid     bool       `json:"valid"`
	Code      string     `json:"code"`
	ExpiredAt *time.Time `json:"expired_at"`
}

type ForgotPasswordRepository interface {
	Create(ForgotPassword) (*ForgotPassword, *resp.ErrorResp)
	Update(ForgotPassword) (*ForgotPassword, *resp.ErrorResp)
	FindOneByCode(code string) (*ForgotPassword, *resp.ErrorResp)
}

type ForgotPasswordUsecase interface {
	Create(ForgotPasswordRequestBody) (*ForgotPassword, *resp.ErrorResp)
	Update(ForgotPasswordUpdateRequestBody) (*ForgotPassword, *resp.ErrorResp)
}
