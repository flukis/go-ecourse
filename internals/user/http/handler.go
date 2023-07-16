package User

import (
	"e-course/domain"
	"e-course/internals/middleware"
	"e-course/pkg/resp"
	"e-course/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	uc domain.UserUsecase
}

func NewUserHandler(uc domain.UserUsecase) *UserHandler {
	return &UserHandler{uc}
}

func (h *UserHandler) Route(r *gin.RouterGroup) {
	v1 := r.Group("/api/v1")

	v1.Use(middleware.AuthJwt)
	{
		v1.GET("/users/:id", h.FindByID)
	}

	v1.Use(middleware.AuthJwt, middleware.AuthAdmin)
	{
		v1.POST("/users", h.Create)
		v1.GET("/users", h.FindAll)
		v1.PATCH("/users/:id", h.Update)
		v1.DELETE("/users/:id", h.Delete)
	}
}

func (h *UserHandler) Create(ctx *gin.Context) {
	var input domain.UserCreateRequestBody
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			resp.Response(
				http.StatusBadRequest,
				http.StatusText(http.StatusBadRequest),
				err.Error(),
			),
		)
		ctx.Abort()
		return
	}

	user := utils.GetCurrentUser(ctx)
	input.CreatedBy = &user.ID

	_, err := h.uc.Create(input)
	if err != nil {
		ctx.JSON(
			int(err.Code),
			resp.Response(
				int(err.Code),
				http.StatusText(int(err.Code)),
				err.Err.Error(),
			),
		)
		ctx.Abort()
		return
	}

	ctx.JSON(
		http.StatusCreated,
		resp.Response(
			http.StatusCreated,
			http.StatusText(http.StatusCreated),
			"register success",
		),
	)
}

func (h *UserHandler) Update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var input domain.UserUpdateRequestBody
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			resp.Response(
				http.StatusBadRequest,
				http.StatusText(http.StatusBadRequest),
				err.Error(),
			),
		)
		ctx.Abort()
		return
	}

	user := utils.GetCurrentUser(ctx)
	input.UpdatedBy = &user.ID

	_, err := h.uc.UpdatePassword(id, input)
	if err != nil {
		ctx.JSON(
			int(err.Code),
			resp.Response(
				int(err.Code),
				http.StatusText(int(err.Code)),
				err.Err.Error(),
			),
		)
		ctx.Abort()
		return
	}

	ctx.JSON(
		http.StatusOK,
		resp.Response(
			http.StatusOK,
			http.StatusText(http.StatusOK),
			"update success",
		),
	)
}

func (h *UserHandler) FindAll(ctx *gin.Context) {
	offset, _ := strconv.Atoi(ctx.Query("offset"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))

	data := h.uc.FindAll(offset, limit)

	ctx.JSON(
		http.StatusOK,
		resp.Response(
			http.StatusOK,
			http.StatusText(http.StatusOK),
			data,
		),
	)
}

func (h *UserHandler) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	err := h.uc.Delete(id)
	if err != nil {
		ctx.JSON(
			int(err.Code),
			resp.Response(
				int(err.Code),
				http.StatusText(int(err.Code)),
				err.Err.Error(),
			),
		)
		ctx.Abort()
		return
	}

	ctx.JSON(
		http.StatusOK,
		resp.Response(
			http.StatusOK,
			http.StatusText(http.StatusOK),
			"delete success",
		),
	)
}

func (h *UserHandler) FindByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	data, err := h.uc.FindOneByID(id)
	if err != nil {
		ctx.JSON(
			int(err.Code),
			resp.Response(
				int(err.Code),
				http.StatusText(int(err.Code)),
				err.Err.Error(),
			),
		)
		ctx.Abort()
		return
	}

	ctx.JSON(
		http.StatusOK,
		resp.Response(
			http.StatusOK,
			http.StatusText(http.StatusOK),
			data,
		),
	)
}
