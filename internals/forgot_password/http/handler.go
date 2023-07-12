package forgotpassword

import (
	"e-course/domain"
	"e-course/pkg/resp"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ForgotPasswordHandler struct {
	uc domain.ForgotPasswordUsecase
}

func (h *ForgotPasswordHandler) Route(r *gin.RouterGroup) {
	r.POST("api/v1/forgot_password", h.Create)
	r.PATCH("api/v1/forgot_password", h.Update)
}

func (h *ForgotPasswordHandler) Create(ctx *gin.Context) {
	var input domain.ForgotPasswordRequestBody
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

	_, err := h.uc.Create(input)
	if err != nil {
		ctx.JSON(int(err.Code), resp.Response(
			int(err.Code),
			http.StatusText(int(err.Code)),
			err.Err.Error(),
		))
		ctx.Abort()
		return
	}

	ctx.JSON(
		http.StatusOK,
		resp.Response(
			http.StatusOK,
			http.StatusText(http.StatusOK),
			"success, please check your email",
		),
	)
}

func (h *ForgotPasswordHandler) Update(ctx *gin.Context) {
	var input domain.ForgotPasswordUpdateRequestBody
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

	_, err := h.uc.Update(input)
	if err != nil {
		ctx.JSON(int(err.Code), resp.Response(
			int(err.Code),
			http.StatusText(int(err.Code)),
			err.Err.Error(),
		))
		ctx.Abort()
		return
	}

	ctx.JSON(
		http.StatusOK,
		resp.Response(
			http.StatusOK,
			http.StatusText(http.StatusOK),
			"success change your password",
		),
	)
}

func NewForgotPasswordhandler(uc domain.ForgotPasswordUsecase) *ForgotPasswordHandler {
	return &ForgotPasswordHandler{uc}
}
