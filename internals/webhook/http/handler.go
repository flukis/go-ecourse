package webhook

import (
	"e-course/domain"
	"e-course/pkg/resp"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WebhookHandler struct {
	uc domain.WebhookUsecase
}

func NewWebhookHandler(uc domain.WebhookUsecase) *WebhookHandler {
	return &WebhookHandler{uc}
}

func (h *WebhookHandler) Route(r *gin.RouterGroup) {
	r.POST("api/v1/webhooks/xendit", h.Xendit)
}

func (h *WebhookHandler) Xendit(ctx *gin.Context) {
	var input domain.WebhookRequestBody
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

	err := h.uc.UpdatePayment(input.ID)
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
			http.StatusText(http.StatusOK),
		),
	)
}
