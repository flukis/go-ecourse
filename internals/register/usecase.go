package register

import (
	"e-course/domain"
	"e-course/pkg/resp"
)

type registerUsecase struct {
	userUsecase domain.UserUsecase
}

func (u *registerUsecase) Register(dto domain.UserCreateRequestBody) *resp.ErrorResp {
	_, err := u.userUsecase.Create(dto)
	if err != nil {
		return err
	}

	return nil
}

func NewRegisterUsecase(userUsecase domain.UserUsecase) domain.RegisterUsecase {
	return &registerUsecase{userUsecase}
}
