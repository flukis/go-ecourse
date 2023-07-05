package admin

import (
	"e-course/domain"
	"e-course/pkg/resp"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type adminUsecase struct {
	repo domain.AdminRepository
}

// Create implements domain.AdminUsecase.
func (u *adminUsecase) Create(data domain.AdminRequestBody) (*domain.Admin, *resp.ErrorResp) {
	hashedPassword, errBcrypt := bcrypt.GenerateFromPassword([]byte(*data.Password), bcrypt.DefaultCost)
	if errBcrypt != nil {
		return nil, &resp.ErrorResp{
			Code: 500,
			Err:  errBcrypt,
		}
	}

	dataAdmin := domain.Admin{
		Name:     data.Name,
		Password: string(hashedPassword),
		Email:    data.Email,
	}

	admin, errCreateAdmin := u.repo.Create(dataAdmin)
	if errCreateAdmin != nil {
		return nil, errCreateAdmin
	}

	return admin, nil
}

// Delete implements domain.AdminUsecase.
func (u *adminUsecase) Delete(id int) *resp.ErrorResp {
	// check if account is there
	existed, err := u.repo.FindOneByID(id)
	if err != nil {
		return err
	}

	if err := u.repo.Delete(*existed); err != nil {
		return err
	}

	return nil
}

// FindAll implements domain.AdminUsecase.
func (u *adminUsecase) FindAll(offset int, limit int) []domain.Admin {
	return u.repo.FindAll(offset, limit)
}

// FindOneByEmail implements domain.AdminUsecase.
func (u *adminUsecase) FindOneByEmail(email string) (*domain.Admin, *resp.ErrorResp) {
	return u.repo.FindOneByEmail(email)
}

// FindOneByID implements domain.AdminUsecase.
func (u *adminUsecase) FindOneByID(id int) (*domain.Admin, *resp.ErrorResp) {
	return u.repo.FindOneByID(id)
}

// TotalCountAdmin implements domain.AdminUsecase.
func (*adminUsecase) TotalCountAdmin() int64 {
	panic("unimplemented")
}

// Update implements domain.AdminUsecase.
func (u *adminUsecase) Update(id int, data domain.AdminRequestBody) (*domain.Admin, *resp.ErrorResp) {
	// check if account is there
	existed, err := u.repo.FindOneByID(id)
	if err != nil {
		return nil, err
	}

	existed.Name = data.Name
	existed.Email = data.Email
	if data.Password == nil {
		return nil, &resp.ErrorResp{
			Code: 500,
			Err:  errors.New("password missing"),
		}
	}
	hashedPwd, bcryptErr := bcrypt.GenerateFromPassword([]byte(*data.Password), bcrypt.DefaultCost)
	if bcryptErr != nil {
		return nil, &resp.ErrorResp{
			Code: 500,
			Err:  bcryptErr,
		}
	}
	existed.Password = string(hashedPwd)

	updatedAdmin, err := u.repo.Update(*existed)
	if err != nil {
		return nil, err
	}

	return updatedAdmin, nil
}

func NewAdminUsecase(repo domain.AdminRepository) domain.AdminUsecase {
	return &adminUsecase{repo: repo}
}
