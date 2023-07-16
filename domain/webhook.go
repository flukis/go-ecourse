package domain

import "e-course/pkg/resp"

type WebhookRequestBody struct {
	ID string `json:"id"`
}

type WebhookUsecase interface {
	UpdatePayment(id string) *resp.ErrorResp
}
