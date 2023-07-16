package webhook

import (
	"e-course/domain"
	"e-course/pkg/resp"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/invoice"
)

type webhookUsecase struct {
	classroomUC domain.ClassRoomUsecase
	orderUC     domain.OrderUsecase
}

// UpdatePayment implements domain.WebhookUsecase.
func (uc *webhookUsecase) UpdatePayment(id string) *resp.ErrorResp {
	params := invoice.GetParams{
		ID: id,
	}

	dataXendit, err := invoice.Get(&params)
	if err != nil {
		return &resp.ErrorResp{
			Code: 500,
			Err:  err,
		}
	}

	if dataXendit == nil {
		return &resp.ErrorResp{
			Code: 404,
			Err:  errors.New("order is not found"),
		}
	}

	dataOrder, errOrd := uc.orderUC.FindOneByExternalId(dataXendit.ExternalID)

	if errOrd != nil {
		return errOrd
	}

	if dataOrder == nil {
		return &resp.ErrorResp{
			Code: 404,
			Err:  errors.New("order is not found"),
		}
	}

	if dataOrder.Status == "settled" {
		return &resp.ErrorResp{
			Code: 400,
			Err:  errors.New("payment is already processed"),
		}
	}

	if dataOrder.Status != "paid" {
		if dataXendit.Status == "PAID" || dataXendit.Status == "SETTLED" {
			for _, orderDetail := range dataOrder.OrderDetails {
				dataClassRoom := domain.ClassRoomRequestBody{
					UserID:    *dataOrder.UserID,
					ProductID: *orderDetail.ProductID,
				}

				_, err := uc.classroomUC.Create(dataClassRoom)
				if err != nil {
					fmt.Println(err)
				}
			}

			// trigger seperti notifikasi email dll
		}
	}

	order := domain.OrderRequestBody{
		Status: strings.ToLower(dataXendit.Status),
	}

	uc.orderUC.Update(int(dataOrder.ID), order)
	return nil
}

func NewWebhookUsecase(
	classroomUC domain.ClassRoomUsecase,
	orderUC domain.OrderUsecase) domain.WebhookUsecase {
	xendit.Opt.SecretKey = os.Getenv("XENDIT_APIKEY")
	return &webhookUsecase{classroomUC, orderUC}
}
