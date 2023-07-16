package user

import (
	"e-course/domain"
	"e-course/pkg/resp"
	"e-course/pkg/utils"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userUsecase struct {
	userRepo domain.UserRepository
}

// Delete implements domain.UserUsecase.
func (u *userUsecase) Delete(id int) *resp.ErrorResp {
	user, err := u.userRepo.FindOneByID(id)
	if err != nil {
		return err
	}
	return u.userRepo.Delete(*user)
}

// FindAll implements domain.UserUsecase.
func (u *userUsecase) FindAll(offset int, limit int) []domain.User {
	return u.userRepo.FindAll(offset, limit)
}

// FindOneByCodeVerified implements domain.UserUsecase.
func (u *userUsecase) FindOneByCodeVerified(codeVerified string) (*domain.User, *resp.ErrorResp) {
	return u.userRepo.FindOneByVerificationCode(codeVerified)
}

// FindOneById implements domain.UserUsecase.
func (u *userUsecase) FindOneById(id int) (*domain.User, *resp.ErrorResp) {
	return u.userRepo.FindOneByID(id)
}

// TotalCountUser implements domain.UserUsecase.
func (u *userUsecase) TotalCountUser() int64 {
	return u.userRepo.TotalCountUser()
}

// UpdatePassword implements domain.UserUsecase.
func (u *userUsecase) UpdatePassword(id int, data domain.UserUpdateRequestBody) (*domain.User, *resp.ErrorResp) {
	user, err := u.userRepo.FindOneByID(id)
	if err != nil {
		return nil, err
	}
	if data.Email != nil {
		if user.Email != *data.Email {
			user.Email = *data.Email
		}
	}
	if data.Name != nil {
		if user.Name != *data.Name {
			user.Name = *data.Name
		}
	}
	user.Name = *data.Name
	if data.Password != nil {
		hashedPwd, errBcrypt := bcrypt.GenerateFromPassword([]byte(*data.Password), bcrypt.DefaultCost)
		if errBcrypt != nil {
			return nil, &resp.ErrorResp{
				Code: 500,
				Err:  errBcrypt,
			}
		}
		user.Password = string(hashedPwd)
	}
	if data.UpdatedBy == nil {
		user.UpdatedByID = data.UpdatedBy
	}
	updatedUser, err := u.userRepo.Update(*user)
	if err != nil {
		return nil, err
	}
	return updatedUser, nil
}

// FindOneByID implements domain.UserUsecase.
func (u *userUsecase) FindOneByID(id int) (*domain.User, *resp.ErrorResp) {
	user, err := u.userRepo.FindOneByID(id)
	if err != nil {
		if errors.Is(err.Err, gorm.ErrRecordNotFound) {
			return nil, &resp.ErrorResp{
				Code: 404,
				Err:  nil,
			}
		}
		return nil, err
	}
	return user, nil
}

// FindByEmail implements domain.UserUsecase.
func (u *userUsecase) FindByEmail(email string) (*domain.User, *resp.ErrorResp) {
	user, err := u.userRepo.FindByEmail(email)
	if err != nil {
		if errors.Is(err.Err, gorm.ErrRecordNotFound) {
			return nil, &resp.ErrorResp{
				Code: 404,
				Err:  nil,
			}
		}
		return nil, err
	}
	return user, nil
}

func (u *userUsecase) Create(data domain.UserCreateRequestBody) (*domain.User, *resp.ErrorResp) {
	existedUser, err := u.userRepo.FindByEmail(data.Email)
	if utils.IsErrorNot404(err) {
		return nil, err
	}

	if existedUser != nil {
		return nil, &resp.ErrorResp{
			Code: 409,
			Err:  errors.New("email is already registered"),
		}
	}

	hashedPwd, errHashedPwd := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if errHashedPwd != nil {
		return nil, &resp.ErrorResp{
			Code: 500,
			Err:  errHashedPwd,
		}
	}

	user := domain.User{
		Name:         data.Email,
		Email:        data.Email,
		Password:     string(hashedPwd),
		CodeVerified: utils.RandString(6),
	}

	dataUser, err := u.userRepo.Create(user)
	if err != nil {
		return nil, &resp.ErrorResp{
			Code: 500,
			Err:  errHashedPwd,
		}
	}

	return dataUser, nil
}

func NewUserUsacase(userUC domain.UserRepository) domain.UserUsecase {
	return &userUsecase{userRepo: userUC}
}
