package payment

import (
	"e-course/domain"
	"e-course/pkg/resp"
	"os"

	xendit "github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/invoice"
)

type paymentUsecase struct{}

// Create implements domain.PaymentUsecase.
func (uc *paymentUsecase) Create(dto domain.PaymentRequestBody) (*xendit.Invoice, *resp.ErrorResp) {
	data := invoice.CreateParams{
		ExternalID:  dto.ExternalID,
		Amount:      float64(dto.Amount),
		Description: dto.Description,
		PayerEmail:  dto.PayerEmail,
		Customer: xendit.InvoiceCustomer{
			Email: dto.PayerEmail,
		},
		CustomerNotificationPreference: xendit.InvoiceCustomerNotificationPreference{
			InvoiceCreated:  []string{"email"},
			InvoiceReminder: []string{"email"},
			InvoicePaid:     []string{"email"},
			InvoiceExpired:  []string{"email"},
		},
		InvoiceDuration:    86400,
		SuccessRedirectURL: os.Getenv("XENDIT_SUCCESS_URL"),
	}

	res, err := invoice.Create(&data)
	if err != nil {
		return nil, &resp.ErrorResp{
			Code: 500,
			Err:  err,
		}
	}

	return res, nil
}

func NewPaymentUsecase() domain.PaymentUsecase {
	return &paymentUsecase{}
}
