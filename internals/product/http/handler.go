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

type ProductHandler struct {
	uc domain.ProductUsecase
}

func NewProductHandler(uc domain.ProductUsecase) *ProductHandler {
	return &ProductHandler{uc}
}

func (h *ProductHandler) Route(r *gin.RouterGroup) {
	v1 := r.Group("/api/v1")

	v1.GET("/products", h.FindAll)
	v1.GET("/products/:id", h.FindByID)
	v1.Use(middleware.AuthJwt, middleware.AuthAdmin)
	{
		v1.POST("/products", h.Create)
		v1.PUT("/products/:id", h.Update)
		v1.DELETE("/products/:id", h.Delete)
	}
}

func (h *ProductHandler) FindAll(ctx *gin.Context) {
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

func (h *ProductHandler) FindByID(ctx *gin.Context) {
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

func (h *ProductHandler) Create(ctx *gin.Context) {
	var input domain.ProductRequestBody
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

func (h *ProductHandler) Update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var input domain.ProductRequestBody
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

func (h *ProductHandler) Delete(ctx *gin.Context) {
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
