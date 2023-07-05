package forgotpassword

import (
	"e-course/domain"
	"e-course/pkg/resp"

	"gorm.io/gorm"
)

type mysqlForgotPasswordRepository struct {
	db *gorm.DB
}

// Create implements domain.ForgotPasswordRepository.
func (m *mysqlForgotPasswordRepository) Create(data domain.ForgotPassword) (*domain.ForgotPassword, *resp.ErrorResp) {
	if err := m.db.Create(&data).Error; err != nil {
		return nil, &resp.ErrorResp{
			Code: 500,
			Err:  err,
		}
	}

	return &data, nil
}

// FindOneByID implements domain.ForgotPasswordRepository.
func (m *mysqlForgotPasswordRepository) FindOneByCode(code string) (*domain.ForgotPassword, *resp.ErrorResp) {
	var forgotPwd domain.ForgotPassword
	if err := m.db.Where("code = ?", code).First(&forgotPwd).Error; err != nil {
		return nil, &resp.ErrorResp{
			Code: 500,
			Err:  err,
		}
	}

	return &forgotPwd, nil
}

// Update implements domain.ForgotPasswordRepository.
func (m *mysqlForgotPasswordRepository) Update(data domain.ForgotPassword) (*domain.ForgotPassword, *resp.ErrorResp) {
	if err := m.db.Save(&data).Error; err != nil {
		return nil, &resp.ErrorResp{
			Code: 500,
			Err:  err,
		}
	}

	return &data, nil
}

func NewMysqlForgotPasswordRepository(db *gorm.DB) domain.ForgotPasswordRepository {
	return &mysqlForgotPasswordRepository{db: db}
}
