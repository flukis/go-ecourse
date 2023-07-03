package repository

import (
	"e-course/domain"
	"e-course/pkg/resp"

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
	if err := m.db.Where("id = ?", id).Error; err != nil {
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
	panic("not implemented")
}

func (m mysqlUserRepository) FindOneByVerificationCode(email string) (*domain.User, *resp.ErrorResp) {
	panic("not implemented")
}

func (m mysqlUserRepository) Delete(u domain.User) (*domain.User, *resp.ErrorResp) {
	panic("not implemented")
}

func (m mysqlUserRepository) TotalCountUser() int64 {
	panic("not implemented")
}

func NewMysqlUserRepository(db *gorm.DB) domain.UserRepository {
	return &mysqlUserRepository{db: db}
}
