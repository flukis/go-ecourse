package domain

import "e-course/pkg/resp"

type RegisterUsecase interface {
	Register(dto UserCreateRequestBody) *resp.ErrorResp
}
