package repository

import (
	"e-course/domain"
	"e-course/pkg/resp"

	"gorm.io/gorm"
)

type mysqlOAuthClientRepository struct {
	db *gorm.DB
}

// FindByClientIDnClientSecret implements domain.OAuthClientRepository.
func (m *mysqlOAuthClientRepository) FindByClientIDnClientSecret(clientId string, clientSecret string) (*domain.OauthClient, *resp.ErrorResp) {
	var oAuthClient domain.OauthClient
	if err := m.db.
		Where("client_id = ?", clientId).
		Where("client_secret = ?", clientSecret).
		First(&oAuthClient).
		Error; err != nil {
		return nil, &resp.ErrorResp{
			Code: 500,
			Err:  err,
		}
	}
	return &oAuthClient, nil
}

func NewOAuthClientRepository(db *gorm.DB) domain.OauthClientRepository {
	return &mysqlOAuthClientRepository{db: db}
}
