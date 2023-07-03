package http_handler

import (
	"e-course/domain"
	"e-course/pkg/resp"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterHandler struct {
	registerUsecase domain.RegisterUsecase
}

func NewRegisterHandler(registerUsecase domain.RegisterUsecase) *RegisterHandler {
	return &RegisterHandler{registerUsecase}
}

func (h *RegisterHandler) Route(r *gin.RouterGroup) {
	r.POST("api/v1/register", h.Register)
}

func (h *RegisterHandler) Register(ctx *gin.Context) {
	var registerRequest domain.UserCreateRequestBody
	if err := ctx.ShouldBindJSON(&registerRequest); err != nil {
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

	err := h.registerUsecase.Register(registerRequest)
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
			"register success, please check your email for verification",
		),
	)
}
