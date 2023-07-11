package product_category

import (
	"e-course/domain"
	media "e-course/pkg/media/cloudinary"
	"e-course/pkg/resp"
)

type productCategoryUsecase struct {
	repo  domain.ProductCategoryRepository
	media media.Media
}

// Create implements domain.ProductCategoryUsecase.
func (uc *productCategoryUsecase) Create(data domain.ProductCategoryRequestBody) (*domain.ProductCategory, *resp.ErrorResp) {
	prdCat := domain.ProductCategory{
		Name:        data.Name,
		CreatedByID: data.CreatedBy,
	}

	if data.Image != nil {
		image, err := uc.media.Upload(*data.Image)
		if err != nil {
			return nil, err
		}

		if image != nil {
			prdCat.Image = image
		}
	}

	res, err := uc.repo.Create(prdCat)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Delete implements domain.ProductCategoryUsecase.
func (uc *productCategoryUsecase) Delete(id int) *resp.ErrorResp {
	existed, err := uc.repo.FindOneByID(id)
	if err != nil {
		return err
	}

	if err := uc.repo.Delete(*existed); err != nil {
		return err
	}

	return nil
}

// FindAll implements domain.ProductCategoryUsecase.
func (uc *productCategoryUsecase) FindAll(offset int, limit int) []domain.ProductCategory {
	return uc.repo.FindAll(offset, limit)
}

// FindOneByID implements domain.ProductCategoryUsecase.
func (uc *productCategoryUsecase) FindOneByID(id int) (*domain.ProductCategory, *resp.ErrorResp) {
	return uc.repo.FindOneByID(id)
}

// Update implements domain.ProductCategoryUsecase.
func (uc *productCategoryUsecase) Update(id int, data domain.ProductCategoryRequestBody) (*domain.ProductCategory, *resp.ErrorResp) {
	existed, err := uc.repo.FindOneByID(id)
	if err != nil {
		return nil, err
	}

	existed.Name = data.Name
	existed.UpdatedByID = data.UpdatedBy

	if data.Image != nil {
		image, err := uc.media.Upload(*data.Image)
		if err != nil {
			return nil, err
		}

		if existed.Image != nil {
			// delete image lama
			_, err := uc.media.Delete(*existed.Image)
			if err != nil {
				return nil, err
			}
		}

		if image != nil {
			existed.Image = image
		}
	}

	res, err := uc.repo.Update(*existed)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func NewProductCategoryUsecase(
	repo domain.ProductCategoryRepository,
	media media.Media) domain.ProductCategoryUsecase {
	return &productCategoryUsecase{repo, media}
}
