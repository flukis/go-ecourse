package classroom

import (
	"e-course/domain"
	"e-course/internals/middleware"
	"e-course/pkg/resp"
	"e-course/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ClassRoomHandler struct {
	uc domain.ClassRoomUsecase
}

func NewClassroomHandler(uc domain.ClassRoomUsecase) *ClassRoomHandler {
	return &ClassRoomHandler{uc}
}

func (h *ClassRoomHandler) Route(r *gin.RouterGroup) {
	v1 := r.Group("/api/v1")
	v1.Use(middleware.AuthJwt)
	{
		v1.GET("/classrooms", h.FindByUserID)
	}
}

func (h *ClassRoomHandler) FindByUserID(ctx *gin.Context) {
	offset, _ := strconv.Atoi(ctx.Query("offset"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))

	user := utils.GetCurrentUser(ctx)

	data := h.uc.FindallByUserId(int(user.ID), offset, limit)

	ctx.JSON(
		http.StatusOK,
		resp.Response(
			http.StatusOK,
			http.StatusText(http.StatusOK),
			data,
		),
	)
}
