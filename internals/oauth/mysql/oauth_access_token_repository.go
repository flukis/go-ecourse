package repository

import (
	"e-course/domain"
	"e-course/pkg/resp"

	"gorm.io/gorm"
)

type mysqlOauthAccessTokenRepository struct {
	db *gorm.DB
}

// Delete implements domain.OauthAccessTokenRepository.
func (m *mysqlOauthAccessTokenRepository) Delete(data domain.OauthAccessToken) *resp.ErrorResp {
	if err := m.db.Delete(&data).Error; err != nil {
		return &resp.ErrorResp{
			Code: 500,
			Err:  err,
		}
	}

	return nil
}

// FindOneByAccessToken implements domain.OauthAccessTokenRepository.
func (m *mysqlOauthAccessTokenRepository) FindOneByAccessToken(accessToken string) (*domain.OauthAccessToken, *resp.ErrorResp) {
	var oauthAccessToken domain.OauthAccessToken
	if err := m.db.
		Where("token = ?", accessToken).
		First(&oauthAccessToken).
		Error; err != nil {
		return nil, &resp.ErrorResp{
			Code: 500,
			Err:  err,
		}
	}

	return &oauthAccessToken, nil
}

// FindOneByOauthAccessTokenID implements domain.OauthAccessTokenRepository.
func (*mysqlOauthAccessTokenRepository) FindOneByOauthAccessTokenID(id int) (*domain.OauthAccessToken, *resp.ErrorResp) {
	panic("unimplemented")
}

// Create implements domain.OauthAccessTokenRepository.
func (m *mysqlOauthAccessTokenRepository) Create(entity domain.OauthAccessToken) (*domain.OauthAccessToken, *resp.ErrorResp) {
	if err := m.db.Create(&entity).Error; err != nil {
		return nil, &resp.ErrorResp{
			Code: 500,
			Err:  err,
		}
	}
	return &entity, nil
}

func NewOauthAccessTokenRepository(db *gorm.DB) domain.OauthAccessTokenRepository {
	return &mysqlOauthAccessTokenRepository{db: db}
}
