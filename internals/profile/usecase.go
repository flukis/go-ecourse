package repository

import (
	"e-course/domain"
	"e-course/pkg/resp"
)

type profileUsecase struct {
	user  domain.UserUsecase
	oauth domain.OauthUsecase
}

// Deactive implements domain.ProfileUsecase.
func (uc *profileUsecase) Deactive(id int) *resp.ErrorResp {
	user, err := uc.user.FindOneByID(id)
	if err != nil {
		return err
	}

	return uc.user.Delete(int(user.ID))
}

// FindProfile implements domain.ProfileUsecase.
func (uc *profileUsecase) FindProfile(id int) (*domain.ProfileResponseBody, *resp.ErrorResp) {
	user, err := uc.user.FindOneByID(id)
	if err != nil {
		return nil, err
	}
	userResp := domain.CreateProfileResponse(*user)
	return &userResp, err
}

// Logout implements domain.ProfileUsecase.
func (uc *profileUsecase) Logout(accessToken string) *resp.ErrorResp {
	return uc.oauth.Logout(accessToken)
}

// Update implements domain.ProfileUsecase.
func (uc *profileUsecase) Update(id int, dto domain.UserUpdateRequestBody) (*domain.User, *resp.ErrorResp) {
	user, err := uc.user.FindOneByID(id)
	if err != nil {
		return nil, err
	}

	updateUser, err := uc.user.UpdatePassword(int(user.ID), dto)
	if err != nil {
		return nil, err
	}

	return updateUser, nil
}

func NewProfileUsecase(
	user domain.UserUsecase,
	oauth domain.OauthUsecase) domain.ProfileUsecase {
	return &profileUsecase{
		user,
		oauth,
	}
}
