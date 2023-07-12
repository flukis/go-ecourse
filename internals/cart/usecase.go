package cart

import (
	"e-course/domain"
	"e-course/pkg/resp"
	"errors"
	"net/http"
)

type cartUsecase struct {
	repo domain.CartRepository
}

// Create implements domain.CartUsecase.
func (uc *cartUsecase) Create(data domain.CartRequestBody) (*domain.Cart, *resp.ErrorResp) {
	cart := domain.Cart{
		UserID:      &data.UserID,
		ProductID:   &data.ProductID,
		Quantity:    1,
		IsChecked:   false,
		CreatedByID: &data.CreatedBy,
	}
	res, err := uc.repo.Create(cart)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Delete implements domain.CartUsecase.
func (uc *cartUsecase) Delete(id int, userId int) *resp.ErrorResp {
	cart, err := uc.repo.FindOneById(id)
	if err != nil {
		return err
	}

	if *(cart.UserID) != int64(id) {
		return &resp.ErrorResp{
			Code: http.StatusUnauthorized,
			Err:  errors.New("you unauthorized to make an action to this cart"),
		}
	}

	return uc.repo.Delete(*cart)
}

// DeleteByUserId implements domain.CartUsecase.
func (uc *cartUsecase) DeleteByUserId(userId int) *resp.ErrorResp {
	return uc.repo.DeleteByUserId(userId)
}

// FindAllByUserId implements domain.CartUsecase.
func (uc *cartUsecase) FindAllByUserId(userId int, offset int, limit int) []domain.Cart {
	return uc.repo.FindAllByUserId(userId, offset, limit)
}

// FindOneById implements domain.CartUsecase.
func (uc *cartUsecase) FindOneById(id int) (*domain.Cart, *resp.ErrorResp) {
	return uc.repo.FindOneById(id)
}

// Update implements domain.CartUsecase.
func (uc *cartUsecase) Update(id int, data domain.CartRequestUpdateBody) (*domain.Cart, *resp.ErrorResp) {
	cart, err := uc.repo.FindOneById(id)
	if err != nil {
		return nil, err
	}

	if *(cart.UserID) != *data.UserID {
		return nil, &resp.ErrorResp{
			Code: http.StatusUnauthorized,
			Err:  errors.New("you unauthorized to make an action to this cart"),
		}
	}

	cart.UserID = data.UserID
	cart.IsChecked = data.IsChecked

	res, err := uc.repo.Update(*cart)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func NewCartUsecase(repo domain.CartRepository) domain.CartUsecase {
	return &cartUsecase{repo}
}
