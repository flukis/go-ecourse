package repository

import (
	"e-course/domain"
	"e-course/pkg/resp"
	"e-course/pkg/utils"

	"gorm.io/gorm"
)

type mysqlUserRepository struct {
	db *gorm.DB
}

func (m mysqlUserRepository) Create(u domain.User) (*domain.User, *resp.ErrorResp) {
	if err := m.db.Create(&u).Error; err != nil {
		return nil, &resp.ErrorResp{
			Code: 500,
			Err:  err,
		}
	}

	return &u, nil
}

func (m mysqlUserRepository) FindByEmail(email string) (*domain.User, *resp.ErrorResp) {
	var user domain.User
	if err := m.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, &resp.ErrorResp{
			Code: 500,
			Err:  err,
		}
	}
	return &user, nil
}

func (m mysqlUserRepository) FindOneByID(id int) (*domain.User, *resp.ErrorResp) {
	var user domain.User
	if err := m.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, &resp.ErrorResp{
			Code: 500,
			Err:  err,
		}
	}
	return &user, nil
}

func (m mysqlUserRepository) Update(u domain.User) (*domain.User, *resp.ErrorResp) {
	if err := m.db.Save(&u).Error; err != nil {
		return nil, &resp.ErrorResp{
			Code: 500,
			Err:  err,
		}
	}
	return &u, nil
}

func (m mysqlUserRepository) FindAll(offset int, limit int) []domain.User {
	var users []domain.User
	m.db.Scopes(utils.Paginate(offset, limit)).Find(&users)
	return users
}

func (m mysqlUserRepository) FindOneByVerificationCode(code string) (*domain.User, *resp.ErrorResp) {
	var user domain.User
	if err := m.db.Where("code_verified = ?", code).First(&user).Error; err != nil {
		return nil, &resp.ErrorResp{
			Code: 500,
			Err:  err,
		}
	}
	return &user, nil
}

func (m mysqlUserRepository) Delete(u domain.User) *resp.ErrorResp {
	if err := m.db.Delete(&u).Error; err != nil {
		return &resp.ErrorResp{
			Code: 500,
			Err:  err,
		}
	}
	return nil
}

func (m mysqlUserRepository) TotalCountUser() int64 {
	var user domain.User
	var totalUser int64
	m.db.Model(&user).Count(&totalUser)
	return totalUser
}

func NewMysqlUserRepository(db *gorm.DB) domain.UserRepository {
	return &mysqlUserRepository{db: db}
}
