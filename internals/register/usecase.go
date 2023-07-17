package register

import (
	"e-course/domain"
	email "e-course/pkg/mail/sendgrid"
	"e-course/pkg/resp"
	"time"
)

type registerUsecase struct {
	userUsecase domain.UserUsecase
	mail        email.Mail
}

// VerificationCode implements domain.RegisterUsecase.
func (u *registerUsecase) VerificationCode(dto domain.VerificationEmailRequestBody) *resp.ErrorResp {
	user, err := u.userUsecase.FindOneByCodeVerified(dto.CodeVerified)
	if err != nil {
		return err
	}
	timeNow := time.Now()
	dataUpdateUser := domain.UserUpdateRequestBody{
		EmailVerifiedAt: &timeNow,
	}
	_, err = u.userUsecase.UpdatePassword(int(user.ID), dataUpdateUser)
	if err != nil {
		return err
	}
	return nil
}

func (u *registerUsecase) Register(dto domain.UserCreateRequestBody) *resp.ErrorResp {
	user, err := u.userUsecase.Create(dto)
	if err != nil {
		return err
	}

	data := email.EmailVerificationBodyRequest{
		SUBJECT:           "Verification Account",
		EMAIL:             dto.Email,
		NAME:              dto.Name,
		VERIFICATION_CODE: user.CodeVerified,
	}

	go u.mail.SendVerificationCode(
		dto.Email,
		data,
	)

	return nil
}

func NewRegisterUsecase(userUsecase domain.UserUsecase, mail email.Mail) domain.RegisterUsecase {
	return &registerUsecase{userUsecase, mail}
}
