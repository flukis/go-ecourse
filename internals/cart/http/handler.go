package product

import (
	"e-course/domain"
	"e-course/internals/middleware"
	"e-course/pkg/resp"
	"e-course/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	uc domain.CartUsecase
}

func NewCartHandler(uc domain.CartUsecase) *CartHandler {
	return &CartHandler{uc}
}

func (h *CartHandler) Route(r *gin.RouterGroup) {
	v1 := r.Group("/api/v1")
	v1.Use(middleware.AuthJwt)
	{
		v1.POST("/carts", h.Create)
		v1.GET("/carts", h.FindByUserID)
		v1.PATCH("/carts/:id", h.Update)
		v1.DELETE("/carts/:id", h.Delete)
	}
}

func (h *CartHandler) FindByUserID(ctx *gin.Context) {
	offset, _ := strconv.Atoi(ctx.Query("offset"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))

	user := utils.GetCurrentUser(ctx)

	data := h.uc.FindAllByUserId(int(user.ID), offset, limit)

	ctx.JSON(
		http.StatusOK,
		resp.Response(
			http.StatusOK,
			http.StatusText(http.StatusOK),
			data,
		),
	)
}

func (h *CartHandler) Create(ctx *gin.Context) {
	var input domain.CartRequestBody
	if err := ctx.ShouldBind(&input); err != nil {
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
	input.UserID = user.ID
	input.CreatedBy = user.ID

	res, err := h.uc.Create(input)
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
			res,
		),
	)
}

func (h *CartHandler) Update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var input domain.CartRequestUpdateBody
	if err := ctx.ShouldBind(&input); err != nil {
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
	input.UserID = &user.ID

	res, err := h.uc.Update(id, input)
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
			res,
		),
	)
}

func (h *CartHandler) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	err := h.uc.DeleteByUserId(id)
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
