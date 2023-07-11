package product

import (
	"e-course/domain"
	media "e-course/pkg/media/cloudinary"
	"e-course/pkg/resp"
)

type productUsecase struct {
	repo  domain.ProductRepository
	media media.Media
}

// Create implements domain.ProductUsecase.
func (uc *productUsecase) Create(data domain.ProductRequestBody) (*domain.Product, *resp.ErrorResp) {
	product := domain.Product{
		ProductCategoryID: &data.ProductCategoryID,
		Title:             data.Title,
		Description:       data.Description,
		IsHighlighted:     data.IsHighlighted,
		Price:             int64(data.Price),
		CreatedByID:       data.CreatedBy,
	}

	// check if image exists, if exists then upload to cld
	if data.Image != nil {
		img, err := uc.media.Upload(*data.Image)
		if err != nil {
			return nil, err
		}
		if img != nil {
			product.Image = img
		}
	}

	// check if video exists, if exists then upload to cld
	if data.Video != nil {
		vid, err := uc.media.Upload(*data.Video)
		if err != nil {
			return nil, err
		}
		if vid != nil {
			product.Video = vid
		}
	}

	res, err := uc.repo.Create(product)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Delete implements domain.ProductUsecase.
func (uc *productUsecase) Delete(id int) *resp.ErrorResp {
	product, err := uc.repo.FindOneById(id)
	if err != nil {
		return err
	}
	return uc.repo.Delete(*product)
}

// FindAll implements domain.ProductUsecase.
func (uc *productUsecase) FindAll(offset int, limit int) []domain.Product {
	return uc.repo.FindAll(offset, limit)
}

// FindOneById implements domain.ProductUsecase.
func (uc *productUsecase) FindOneById(id int) (*domain.Product, *resp.ErrorResp) {
	return uc.repo.FindOneById(id)
}

// TotalCountProduct implements domain.ProductUsecase.
func (uc *productUsecase) TotalCountProduct() int64 {
	return uc.repo.TotalCountProduct()
}

// Update implements domain.ProductUsecase.
func (uc *productUsecase) Update(id int, data domain.ProductRequestBody) (*domain.Product, *resp.ErrorResp) {
	product, err := uc.repo.FindOneById(id)
	if err != nil {
		return nil, err
	}

	product.ProductCategoryID = &data.ProductCategoryID
	product.Title = data.Title
	product.Description = data.Description
	product.IsHighlighted = data.IsHighlighted
	product.Price = int64(data.Price)
	product.UpdatedByID = data.UpdatedBy

	// check if image exists, if exists then upload to cld
	if data.Image != nil {
		img, err := uc.media.Upload(*data.Image)
		if err != nil {
			return nil, err
		}
		if product.Image != nil {
			// delete
			_, err := uc.media.Delete(*product.Image)
			if err != nil {
				return nil, err
			}
		}
		if img != nil {
			product.Image = img
		}
	}

	// check if video exists, if exists then upload to cld
	if data.Video != nil {
		vid, err := uc.media.Upload(*data.Video)
		if err != nil {
			return nil, err
		}
		if product.Video != nil {
			// delete
			_, err := uc.media.Delete(*product.Video)
			if err != nil {
				return nil, err
			}
		}
		if vid != nil {
			product.Video = vid
		}
	}

	res, err := uc.repo.Update(*product)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func NewProductUsecase(
	repo domain.ProductRepository,
	media media.Media) domain.ProductUsecase {
	return &productUsecase{repo, media}
}
