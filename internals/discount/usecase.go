package discount

import (
	"e-course/domain"
	media "e-course/pkg/media/cloudinary"
	"e-course/pkg/resp"
)

type discountUsecase struct {
	repo  domain.DiscountRepository
	media media.Media
}

// UpdateRemainingQuantity implements domain.DiscountUsecase.
func (uc *discountUsecase) UpdateRemainingQuantity(id int, quantity int, operator string) (*domain.Discount, *resp.ErrorResp) {
	existed, err := uc.repo.FindOneById(id)
	if err != nil {
		return nil, err
	}

	if operator == "+" {
		existed.RemainingQuantity = existed.Quantity + int64(quantity)
	} else if operator == "-" {
		existed.RemainingQuantity = existed.Quantity - int64(quantity)
	}

	res, err := uc.repo.Update(*existed)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Create implements domain.discountUsecase.
func (uc *discountUsecase) Create(data domain.DiscountRequestBody) (*domain.Discount, *resp.ErrorResp) {
	discount := domain.Discount{
		Name:              data.Name,
		Code:              data.Code,
		Quantity:          data.Quantity,
		RemainingQuantity: data.Quantity,
		Type:              data.Type,
		Value:             data.Value,
		StartDate:         data.StartDate,
		EndDate:           data.EndDate,
		CreatedByID:       data.CreatedBy,
	}

	res, err := uc.repo.Create(discount)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Delete implements domain.discountUsecase.
func (uc *discountUsecase) Delete(id int) *resp.ErrorResp {
	discount, err := uc.repo.FindOneById(id)
	if err != nil {
		return err
	}
	return uc.repo.Delete(*discount)
}

// FindAll implements domain.discountUsecase.
func (uc *discountUsecase) FindAll(offset int, limit int) []domain.Discount {
	return uc.repo.FindAll(offset, limit)
}

// FindOneById implements domain.discountUsecase.
func (uc *discountUsecase) FindOneById(id int) (*domain.Discount, *resp.ErrorResp) {
	return uc.repo.FindOneById(id)
}

// FindOneByCode implements domain.discountUsecase.
func (uc *discountUsecase) FindOneByCode(code string) (*domain.Discount, *resp.ErrorResp) {
	return uc.repo.FindOneByCode(code)
}

// Update implements domain.discountUsecase.
func (uc *discountUsecase) Update(id int, data domain.DiscountRequestBody) (*domain.Discount, *resp.ErrorResp) {
	existed, err := uc.repo.FindOneById(id)
	if err != nil {
		return nil, err
	}

	existed.Name = data.Name
	existed.Code = data.Code
	existed.Quantity = data.Quantity
	existed.RemainingQuantity = data.Quantity
	existed.Type = data.Type
	existed.Value = data.Value
	existed.StartDate = data.StartDate
	existed.EndDate = data.EndDate
	existed.UpdatedByID = data.UpdatedBy

	res, err := uc.repo.Update(*existed)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func NewDiscountUsecase(
	repo domain.DiscountRepository,
	media media.Media) domain.DiscountUsecase {
	return &discountUsecase{repo, media}
}
