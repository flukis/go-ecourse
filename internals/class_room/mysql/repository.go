package repository

import (
	"e-course/domain"
	"e-course/pkg/resp"
	"e-course/pkg/utils"

	"gorm.io/gorm"
)

type classroomRepository struct {
	db *gorm.DB
}

// Create implements domain.ClassRoomRepository.
func (r *classroomRepository) Create(entity domain.ClassRoom) (*domain.ClassRoom, *resp.ErrorResp) {
	if err := r.db.Create(&entity).Error; err != nil {
		return nil, &resp.ErrorResp{
			Code: 500,
			Err:  err,
		}
	}

	return &entity, nil
}

// FindAllByUserId implements domain.ClassRoomRepository.
func (r *classroomRepository) FindAllByUserId(userId int, offset int, limit int) []domain.ClassRoom {
	var rooms []domain.ClassRoom

	r.db.Scopes(utils.Paginate(offset, limit)).
		Preload("Product.ProductCategory").
		Where("user_id = ?", userId).
		Find(&rooms)

	return rooms
}

// FindOneByUserIdAndProductId implements domain.ClassRoomRepository.
func (r *classroomRepository) FindOneByUserIdAndProductId(userId int, productId int) (*domain.ClassRoom, *resp.ErrorResp) {
	var room domain.ClassRoom

	if err := r.db.Preload("Product.ProductCategory").
		Where("user_id = ?", userId).
		Where("product_id = ?", userId).
		Find(&room); err != nil {
		return nil, &resp.ErrorResp{
			Code: 500,
			Err:  nil,
		}
	}

	return &room, nil
}

func NewClassroomRepository(db *gorm.DB) domain.ClassRoomRepository {
	return &classroomRepository{db}
}
