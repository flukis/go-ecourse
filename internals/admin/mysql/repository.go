package admin

import (
	"e-course/domain"
	"e-course/pkg/resp"
	"e-course/pkg/utils"

	"gorm.io/gorm"
)

type mysqlAdminRepository struct {
	db *gorm.DB
}

// Create implements domain.AdminRepository.
func (r *mysqlAdminRepository) Create(data domain.Admin) (*domain.Admin, *resp.ErrorResp) {
	if err := r.db.Create(&data).Error; err != nil {
		return nil, &resp.ErrorResp{
			Code: 500,
			Err:  err,
		}
	}
	return &data, nil
}

// Delete implements domain.AdminRepository.
func (r *mysqlAdminRepository) Delete(data domain.Admin) *resp.ErrorResp {
	if err := r.db.Delete(&data).Error; err != nil {
		return &resp.ErrorResp{
			Code: 500,
			Err:  err,
		}
	}

	return nil
}

// FindAll implements domain.AdminRepository.
func (r *mysqlAdminRepository) FindAll(offset int, limit int) []domain.Admin {
	var admins []domain.Admin

	r.db.Scopes(utils.Paginate(offset, limit)).Find(&admins)

	return admins
}

// FindOneByEmail implements domain.AdminRepository.
func (r *mysqlAdminRepository) FindOneByEmail(email string) (*domain.Admin, *resp.ErrorResp) {
	var admin domain.Admin
	if err := r.db.Where("email = ?", email).First(&admin).Error; err != nil {
		return nil, &resp.ErrorResp{
			Code: 500,
			Err:  err,
		}
	}
	return &admin, nil
}

// FindOneByID implements domain.AdminRepository.
func (r *mysqlAdminRepository) FindOneByID(id int) (*domain.Admin, *resp.ErrorResp) {
	var admin domain.Admin
	if err := r.db.Where("id = ?", id).First(&admin).Error; err != nil {
		return nil, &resp.ErrorResp{
			Code: 500,
			Err:  err,
		}
	}
	return &admin, nil
}

// TotalCountAdmin implements domain.AdminRepository.
func (*mysqlAdminRepository) TotalCountAdmin() int64 {
	panic("unimplemented")
}

// Update implements domain.AdminRepository.
func (r *mysqlAdminRepository) Update(data domain.Admin) (*domain.Admin, *resp.ErrorResp) {
	if err := r.db.Save(&data).Error; err != nil {
		return nil, &resp.ErrorResp{
			Code: 500,
			Err:  err,
		}
	}
	return &data, nil
}

func NewMysqlAdminRepository(db *gorm.DB) domain.AdminRepository {
	return &mysqlAdminRepository{db}
}
