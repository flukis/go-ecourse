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

type DiscountHandler struct {
	uc domain.DiscountUsecase
}

func NewDiscountHandler(uc domain.DiscountUsecase) DiscountHandler {
	return DiscountHandler{uc}
}

func (h DiscountHandler) Route(r *gin.RouterGroup) {
	v1 := r.Group("/api/v1")
	v1.Use(middleware.AuthJwt, middleware.AuthAdmin)
	{
		v1.POST("/discounts", h.Create)
		v1.GET("/discounts", h.FindAll)
		v1.GET("/discounts/:id", h.FindByID)
		v1.PATCH("/discounts/:id", h.Update)
		v1.DELETE("/discounts/:id", h.Delete)
	}
}

func (h DiscountHandler) FindAll(ctx *gin.Context) {
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

func (h DiscountHandler) FindByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	data, err := h.uc.FindOneById(id)
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

func (h DiscountHandler) FindByCode(ctx *gin.Context) {
	id := ctx.Param("code")

	data, err := h.uc.FindOneByCode(id)
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

func (h DiscountHandler) Create(ctx *gin.Context) {
	var input domain.DiscountRequestBody
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
	input.CreatedBy = &user.ID

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

func (h DiscountHandler) Update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var input domain.DiscountRequestBody
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
	input.UpdatedBy = &user.ID

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

func (h DiscountHandler) Delete(ctx *gin.Context) {
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
