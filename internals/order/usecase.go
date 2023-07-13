package order

import (
	"e-course/domain"
	"e-course/pkg/resp"
	"e-course/pkg/utils"
	"errors"
	"fmt"
	"strconv"
	"time"
)

type orderUsecase struct {
	orderRepo          domain.OrderRepository
	cartUsecase        domain.CartUsecase
	discountUsecase    domain.DiscountUsecase
	orderDetailUsecase domain.OrderDetailUsecase
	paymentUsecase     domain.PaymentUsecase
	productUsecase     domain.ProductUsecase
}

// Create implements domain.OrderUsecase.
func (uc *orderUsecase) Create(dto domain.OrderRequestBody) (*domain.Order, *resp.ErrorResp) {
	price := 0
	totalPrice := 0
	desc := ""

	var products []domain.Product

	order := domain.Order{
		UserID: &dto.UserID,
		Status: "pending",
	}

	var discount *domain.Discount

	carts := uc.cartUsecase.FindAllByUserId(int(dto.UserID), 1, 9999)
	if len(carts) == 0 {
		// if empty, check user maybe he directly buy without cart
		if dto.ProductID == nil {
			return nil, &resp.ErrorResp{
				Code: 400,
				Err:  errors.New("carts is empty"),
			}
		}
	}

	if dto.DiscountCode != nil {
		dcn, err := uc.discountUsecase.FindOneByCode(*dto.DiscountCode)
		if err != nil {
			return nil, &resp.ErrorResp{
				Code: 400,
				Err:  errors.New("discount with that code is not found"),
			}
		}

		if dcn.RemainingQuantity == 0 {
			return nil, &resp.ErrorResp{
				Code: 400,
				Err:  errors.New("discount is expired"),
			}
		}

		if dcn.StartDate != nil && dcn.EndDate != nil {
			if dcn.StartDate.After(time.Now()) || dcn.EndDate.Before(time.Now()) {
				return nil, &resp.ErrorResp{
					Code: 400,
					Err:  errors.New("discount is expired"),
				}
			} else if dcn.StartDate.After(time.Now()) {
				return nil, &resp.ErrorResp{
					Code: 400,
					Err:  errors.New("discount is expired"),
				}
			} else if dcn.EndDate.Before(time.Now()) {
				return nil, &resp.ErrorResp{
					Code: 400,
					Err:  errors.New("discount is expired"),
				}
			}
		}

		discount = dcn
	}

	if len(carts) > 0 {
		// menggunakan data cart
		for _, cart := range carts {
			product, err := uc.productUsecase.FindOneById(int(*cart.ProductID))
			if err != nil {
				return nil, err
			}
			products = append(products, *product)
		}
	} else if dto.ProductID != nil {
		product, err := uc.productUsecase.FindOneById(int(*dto.ProductID))
		if err != nil {
			return nil, err
		}
		products = append(products, *product)
	}

	// calculate price
	for i, product := range products {
		price += int(product.Price)
		i := strconv.Itoa(i + 1)
		desc += fmt.Sprintf("%s. Product : %s<br />", i, product.Title)
	}

	totalPrice = price

	// check discount
	if discount != nil {
		if discount.Type == "rebate" {
			totalPrice = price - int(discount.Value)
		} else if discount.Type == "percent" {
			totalPrice = price - (price / 100 * int(discount.Value))
		}

		order.DiscountID = &discount.ID
	}

	order.Price = int64(price)
	order.TotalPrice = int64(totalPrice)
	order.CreatedByID = &dto.UserID

	extId, errGen := utils.GenerateRefreshToken()
	if errGen != nil {
		return nil, &resp.ErrorResp{
			Code: 500,
			Err:  errGen,
		}
	}
	externalId := extId.String()

	data, err := uc.orderRepo.Create(order)
	if err != nil {
		return nil, err
	}

	for _, product := range products {
		orderDetail := domain.OrderDetail{
			OrderID:     data.ID,
			ProductID:   &product.ID,
			Price:       product.Price,
			CreatedByID: order.UserID,
		}

		uc.orderDetailUsecase.Create(orderDetail)
	}

	// hit payment xendit
	dataPayment := domain.PaymentRequestBody{
		ExternalID:  externalId,
		Amount:      int64(totalPrice),
		PayerEmail:  dto.Email,
		Description: desc,
	}

	payment, err := uc.paymentUsecase.Create(dataPayment)
	if err != nil {
		return nil, err
	}

	data.CheckoutLink = payment.InvoiceURL

	// update qty
	if dto.DiscountCode != nil {
		_, err := uc.discountUsecase.UpdateRemainingQuantity(int(discount.ID), 1, "-")
		if err != nil {
			return nil, err
		}
	}

	// delete carts
	err = uc.cartUsecase.DeleteByUserId(int(dto.UserID))
	if err != nil {
		return nil, err
	}

	return data, err
}

// FindAllByUserId implements domain.OrderUsecase.
func (uc *orderUsecase) FindAllByUserId(userId int, offset int, limit int) []domain.Order {
	return uc.orderRepo.FindAllByUserId(userId, offset, limit)
}

// FindOneByExternalId implements domain.OrderUsecase.
func (uc *orderUsecase) FindOneByExternalId(externalId string) (*domain.Order, *resp.ErrorResp) {
	return uc.orderRepo.FindOneByExternalId(externalId)
}

// FindOneById implements domain.OrderUsecase.
func (uc *orderUsecase) FindOneById(id int) (*domain.Order, *resp.ErrorResp) {
	return uc.orderRepo.FindOneById(id)
}

// TotalCountOrder implements domain.OrderUsecase.
func (uc *orderUsecase) TotalCountOrder() int64 {
	return uc.orderRepo.TotalCountOrder()
}

// Update implements domain.OrderUsecase.
func (uc *orderUsecase) Update(id int, dto domain.OrderRequestBody) (*domain.Order, *resp.ErrorResp) {
	order, err := uc.orderRepo.FindOneById(id)
	if err != nil {
		return nil, err
	}

	order.Status = dto.Status

	updateOrder, err := uc.orderRepo.Update(*order)
	if err != nil {
		return nil, err
	}

	return updateOrder, nil
}

func NewOrderUsecase(
	orderRepo domain.OrderRepository,
	cartUsecase domain.CartUsecase,
	discountUsecase domain.DiscountUsecase,
	orderDetailUsecase domain.OrderDetailUsecase,
	paymentUsecase domain.PaymentUsecase,
	productUsecase domain.ProductUsecase) domain.OrderUsecase {
	return &orderUsecase{
		orderRepo:          orderRepo,
		cartUsecase:        cartUsecase,
		discountUsecase:    discountUsecase,
		orderDetailUsecase: orderDetailUsecase,
		paymentUsecase:     paymentUsecase,
		productUsecase:     productUsecase,
	}
}
