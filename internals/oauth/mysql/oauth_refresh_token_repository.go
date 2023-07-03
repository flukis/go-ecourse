package repository

import (
	"e-course/domain"
	"e-course/pkg/resp"

	"gorm.io/gorm"
)

type mysqlOauthRefreshTokenRepository struct {
	db *gorm.DB
}

// Delete implements domain.OauthRefreshTokenRepository.
func (*mysqlOauthRefreshTokenRepository) Delete(domain.OauthRefreshToken) *resp.ErrorResp {
	panic("unimplemented")
}

// FindOneByToken implements domain.OauthRefreshTokenRepository.
func (*mysqlOauthRefreshTokenRepository) FindOneByToken(token string) (*domain.OauthRefreshToken, *resp.ErrorResp) {
	panic("unimplemented")
}

// Create implements domain.OauthRefreshTokenRepository.
func (m *mysqlOauthRefreshTokenRepository) Create(entity domain.OauthRefreshToken) (*domain.OauthRefreshToken, *resp.ErrorResp) {
	if err := m.db.Create(&entity).Error; err != nil {
		return nil, &resp.ErrorResp{
			Code: 500,
			Err:  err,
		}
	}
	return &entity, nil
}

// FindOneByRefreshToken implements domain.OauthRefreshTokenRepository.
func (*mysqlOauthRefreshTokenRepository) FindOneByRefreshToken(RefreshToken string) (*domain.OauthRefreshToken, *resp.ErrorResp) {
	panic("unimplemented")
}

func NewOauthRefreshTokenRepository(db *gorm.DB) domain.OauthRefreshTokenRepository {
	return &mysqlOauthRefreshTokenRepository{db: db}
}
