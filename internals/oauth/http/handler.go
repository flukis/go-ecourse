package oauth

import (
	"e-course/domain"
	"e-course/pkg/resp"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OAuthHandler struct {
	uc domain.OauthUsecase
}

func NewOAuthHandler(uc domain.OauthUsecase) *OAuthHandler {
	return &OAuthHandler{uc: uc}
}

func (h *OAuthHandler) Login(ctx *gin.Context) {
	var input domain.LoginRequestBody
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

	// Panggil fungsi dari Login
	data, err := h.uc.Login(input)
	if err != nil {
		ctx.JSON(int(err.Code), resp.Response(
			int(err.Code),
			http.StatusText(int(err.Code)),
			err.Err.Error(),
		))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, resp.Response(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		data,
	))
}

func (h *OAuthHandler) Refresh(ctx *gin.Context) {
	var input domain.RefreshTokenRequestBody
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

	// Panggil fungsi dari Login
	data, err := h.uc.Refresh(input)
	if err != nil {
		ctx.JSON(int(err.Code), resp.Response(
			int(err.Code),
			http.StatusText(int(err.Code)),
			err.Err.Error(),
		))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, resp.Response(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		data,
	))
}

func (h *OAuthHandler) Route(r *gin.RouterGroup) {
	v1 := r.Group("/api/v1")

	v1.POST("/login", h.Login)
	v1.POST("/refresh_token", h.Refresh)
}
