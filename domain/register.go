package domain

import "e-course/pkg/resp"

type VerificationEmailRequestBody struct {
	CodeVerified string `json:"code_verified" binding:"required"`
}

type RegisterUsecase interface {
	Register(dto UserCreateRequestBody) *resp.ErrorResp
	VerificationCode(dto VerificationEmailRequestBody) *resp.ErrorResp
}
