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
func (*mysqlOauthAccessTokenRepository) Delete(domain.OauthAccessToken) *resp.ErrorResp {
	panic("unimplemented")
}

// FindOneByAccessToken implements domain.OauthAccessTokenRepository.
func (*mysqlOauthAccessTokenRepository) FindOneByAccessToken(accessToken string) (*domain.OauthAccessToken, *resp.ErrorResp) {
	panic("unimplemented")
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
