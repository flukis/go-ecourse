package register

import (
	"e-course/domain"
	email "e-course/pkg/mail/sendgrid"
	"e-course/pkg/resp"
)

type registerUsecase struct {
	userUsecase domain.UserUsecase
	mail        email.Mail
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
