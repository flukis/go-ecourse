package forgotpassword

import (
	"e-course/domain"
	email "e-course/pkg/mail/sendgrid"
	"e-course/pkg/resp"
	"e-course/pkg/utils"
	"errors"
	"time"
)

type forgotPasswordUsecase struct {
	repo   domain.ForgotPasswordRepository
	userUC domain.UserUsecase
	mail   email.Mail
}

// Create implements domain.ForgotPasswordUsecase.
func (u *forgotPasswordUsecase) Create(data domain.ForgotPasswordRequestBody) (*domain.ForgotPassword, *resp.ErrorResp) {
	// check email is exist
	user, err := u.userUC.FindByEmail(data.Email)
	if err != nil {
		return nil, &resp.ErrorResp{
			Code: 404,
			Err:  errors.New("user not found"),
		}
	}

	if user == nil {
		return nil, &resp.ErrorResp{
			Code: 200,
			Err:  errors.New("success, please check your email"),
		}
	}

	expDate := time.Now().Add(3 * 1 * time.Hour)
	forgotPwd := domain.ForgotPassword{
		UserID:    &user.ID,
		Valid:     true,
		Code:      utils.RandString(12),
		ExpiredAt: &expDate,
	}

	resForgotPwd, err := u.repo.Create(forgotPwd)

	// send email
	dataEmailForgotPwd := email.EmailForgotPasswordBodyRequest{
		SUBJECT: "Code Forgot Password",
		EMAIL:   user.Email,
		NAME:    user.Name,
		CODE:    forgotPwd.Code,
	}

	go u.mail.SendForgotPassword(user.Email, dataEmailForgotPwd)

	if err != nil {
		return nil, err
	}

	return resForgotPwd, nil
}

// Update implements domain.ForgotPasswordUsecase.
func (u *forgotPasswordUsecase) Update(data domain.ForgotPasswordUpdateRequestBody) (*domain.ForgotPassword, *resp.ErrorResp) {
	// check code
	code, err := u.repo.FindOneByCode(data.Code)
	if err != nil || !code.Valid {
		return nil, &resp.ErrorResp{
			Code: 400,
			Err:  errors.New("code is invalid"),
		}
	}

	//search user
	existingUser, err := u.userUC.FindOneByID(int(*code.UserID))
	if err != nil {
		return nil, err
	}

	dataUser := domain.UserUpdateRequestBody{
		Password: &data.Password,
	}

	_, err = u.userUC.UpdatePassword(int(existingUser.ID), dataUser)
	if err != nil {
		return nil, err
	}

	code.Valid = false
	u.repo.Update(*code)
	return code, err
}

func NewForgotPasswordUsecase(repo domain.ForgotPasswordRepository, uuc domain.UserUsecase, mail email.Mail) domain.ForgotPasswordUsecase {
	return &forgotPasswordUsecase{repo: repo, userUC: uuc, mail: mail}
}
