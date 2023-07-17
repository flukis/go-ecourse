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
func (m *mysqlOauthRefreshTokenRepository) Delete(data domain.OauthRefreshToken) *resp.ErrorResp {
	if err := m.db.Delete(&data).Error; err != nil {
		return &resp.ErrorResp{
			Code: 500,
			Err:  err,
		}
	}

	return nil
}

// FindOneByToken implements domain.OauthRefreshTokenRepository.
func (m *mysqlOauthRefreshTokenRepository) FindOneByToken(token string) (*domain.OauthRefreshToken, *resp.ErrorResp) {
	var refrehToken domain.OauthRefreshToken
	if err := m.db.
		Preload("OauthAccessToken").
		Where("token = ?", token).
		First(&refrehToken).
		Error; err != nil {
		return nil, &resp.ErrorResp{
			Code: 500,
			Err:  err,
		}
	}

	return &refrehToken, nil
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
func (m *mysqlOauthRefreshTokenRepository) FindOneByAccessTokenID(accessToken int) (*domain.OauthRefreshToken, *resp.ErrorResp) {
	var refrehToken domain.OauthRefreshToken
	if err := m.db.
		Where("oauth_access_token_id = ?", accessToken).
		First(&refrehToken).
		Error; err != nil {
		return nil, &resp.ErrorResp{
			Code: 500,
			Err:  err,
		}
	}

	return &refrehToken, nil
}

func NewOauthRefreshTokenRepository(db *gorm.DB) domain.OauthRefreshTokenRepository {
	return &mysqlOauthRefreshTokenRepository{db: db}
}
