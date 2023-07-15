package classroom

import (
	"e-course/domain"
	"e-course/pkg/resp"
	"errors"

	"gorm.io/gorm"
)

type classroomUsecase struct {
	repo domain.ClassRoomRepository
}

// Create implements domain.ClassRoomUsecase.
func (uc *classroomUsecase) Create(dto domain.ClassRoomRequestBody) (*domain.ClassRoom, *resp.ErrorResp) {
	existedClassroom, err := uc.repo.FindOneByUserIdAndProductId(int(dto.UserID), int(dto.ProductID))

	if err != nil && !errors.Is(err.Err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if existedClassroom != nil {
		return nil, &resp.ErrorResp{
			Code: 404,
			Err:  errors.New("already buy this class"),
		}
	}

	classroom := domain.ClassRoom{
		UserID:      dto.UserID,
		ProductID:   &dto.ProductID,
		CreatedByID: &dto.UserID,
	}

	data, err := uc.repo.Create(classroom)
	if err != nil {
		return nil, err
	}

	return data, err
}

// FindOneByUserIdAndProductId implements domain.ClassRoomUsecase.
func (uc *classroomUsecase) FindOneByUserIdAndProductId(userId int, productId int) (*domain.ClassRoomResponseBody, *resp.ErrorResp) {
	classroom, _ := uc.repo.FindOneByUserIdAndProductId(userId, productId)
	classroomResp := domain.CreateClassRoomResponse(*classroom)
	return &classroomResp, nil
}

// FindallByUserId implements domain.ClassRoomUsecase.
func (uc *classroomUsecase) FindallByUserId(userId int, offset int, limit int) domain.ClassRoomListResponse {
	classroom := uc.repo.FindAllByUserId(userId, offset, limit)
	classroomResp := domain.CreateClassRoomListResponse(classroom)
	return classroomResp
}

func NewClassroomUsecase(repo domain.ClassRoomRepository) domain.ClassRoomUsecase {
	return &classroomUsecase{repo}
}
