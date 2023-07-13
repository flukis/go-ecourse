package domain

import (
	"e-course/pkg/resp"

	"github.com/xendit/xendit-go"
)

type PaymentUsecase interface {
	Create(data PaymentRequestBody) (*xendit.Invoice, *resp.ErrorResp)
}

type PaymentRequestBody struct {
	ExternalID  string `json:"external_id"`
	Amount      int64  `json:"amount"`
	PayerEmail  string `json:"payer_email"`
	Description string `json:"description"`
}
