package domain

import (
	"e-course/pkg/resp"
	"time"

	"gorm.io/gorm"
)

type ClassRoom struct {
	ID          int64          `json:"id"`
	User        *User          `json:"user" gorm:"foreignKey:UserID;references:ID"`
	UserID      int64          `json:"user_id"`
	Product     *Product       `json:"product" gorm:"foreignKey:ProductID;references:ID"`
	ProductID   *int64         `json:"product_id"`
	CreatedBy   *User          `json:"-" gorm:"foreignKey:CreatedByID;references:ID"`
	CreatedByID *int64         `json:"created_by" gorm:"column:created_by"`
	UpdateBy    *User          `json:"-" gorm:"foreignKey:CreatedByID;references:ID"`
	UpdatedByID *int64         `json:"updated_by" gorm:"column:updated_by"`
	CreatedAt   *time.Time     `json:"created_at"`
	UpdatedAt   *time.Time     `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}

type ClassRoomRepository interface {
	FindAllByUserId(userId int, offset int, limit int) []ClassRoom
	FindOneByUserIdAndProductId(userId int, productId int) (*ClassRoom, *resp.ErrorResp)
	Create(entity ClassRoom) (*ClassRoom, *resp.ErrorResp)
}

type ClassRoomUsecase interface {
	FindallByUserId(userId int, offset int, limit int) ClassRoomListResponse
	FindOneByUserIdAndProductId(userId int, productId int) (*ClassRoomResponseBody, *resp.ErrorResp)
	Create(dto ClassRoomRequestBody) (*ClassRoom, *resp.ErrorResp)
}

type ClassRoomRequestBody struct {
	UserID    int64 `json:"user_id"`
	ProductID int64 `json:"product_id"`
}

type ClassRoomResponseBody struct {
	ID        int64          `json:"id"`
	User      *User          `json:"user"`
	Product   *Product       `json:"product"`
	CreatedBy *User          `json:"created_by"`
	UpdatedBy *User          `json:"updated_by"`
	CreatedAt *time.Time     `json:"created_at"`
	UpdatedAt *time.Time     `json:"updated_at"`
	DeleteAt  gorm.DeletedAt `json:"deleted_at"`
}

func CreateClassRoomResponse(entity ClassRoom) ClassRoomResponseBody {
	return ClassRoomResponseBody{
		ID:        entity.ID,
		User:      entity.User,
		Product:   entity.Product,
		CreatedBy: entity.CreatedBy,
		UpdatedBy: entity.UpdateBy,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
		DeleteAt:  entity.DeletedAt,
	}
}

type ClassRoomListResponse []ClassRoomResponseBody

func CreateClassRoomListResponse(entity []ClassRoom) ClassRoomListResponse {
	classRoomResp := ClassRoomListResponse{}

	for _, classRoom := range entity {
		classRoom.Product.VideoLink = classRoom.Product.Video

		classRoomResponse := CreateClassRoomResponse(classRoom)
		classRoomResp = append(classRoomResp, classRoomResponse)
	}

	return classRoomResp
}
