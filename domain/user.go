package domain

import (
	"e-course/pkg/resp"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID              int64          `json:"id"`
	Name            string         `json:"name"`
	Password        string         `json:"-"`
	Email           string         `json:"email"`
	CodeVerified    string         `json:"-"`
	EmailVerifiedAt *time.Time     `json:"email_verified_at"`
	CreatedAt       *time.Time     `json:"created_at"`
	UpdatedAt       *time.Time     `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at"`

	//CreatedByID *int64 `json:"created_by" gorm:"column:created_by"`
	//CreatedBy   *Admin `json:"-" gorm:"foreignKey:CreatedByID;references:ID"`

	//UpdatedByID *int64 `json:"updated_by" gorm:"column:updated_by"`
	//UpdatedBy   *Admin `json:"-" gorm:"foreignKey:UpdatedByID;references:ID"`
}

type UserRepository interface {
	FindAll(offset int, limit int) []User
	FindOneByID(id int) (*User, *resp.ErrorResp)
	FindByEmail(email string) (*User, *resp.ErrorResp)
	FindOneByVerificationCode(email string) (*User, *resp.ErrorResp)
	Create(u User) (*User, *resp.ErrorResp)
	Update(u User) (*User, *resp.ErrorResp)
	Delete(u User) (*User, *resp.ErrorResp)
	TotalCountUser() int64
}

type UserCreateRequestBody struct {
	Name      string `json:"name" binding:"required"`
	Email     string `json:"email" binding:"email"`
	Password  string `json:"password" binding:"required"`
	CreatedBy *int64 `json:"created_by"`
}

type UserUpdateRequestBody struct {
	Name            string     `json:"name" `
	Email           string     `json:"email"`
	Password        *string    `json:"password"`
	EmailVerifiedAt *time.Time `json:"email_verified_at"`
	CreatedBy       *int64     `json:"created_by"`
}

type UserUsecase interface {
	Create(data UserCreateRequestBody) (*User, *resp.ErrorResp)
	FindByEmail(email string) (*User, *resp.ErrorResp)
	FindOneByID(id int) (*User, *resp.ErrorResp)
	UpdatePassword(id int, data UserUpdateRequestBody) (*User, *resp.ErrorResp)
}
