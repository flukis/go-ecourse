package http

import (
	"e-course/domain"
	"e-course/internals/middleware"
	"e-course/pkg/resp"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DashboardHandler struct {
	uc domain.DashboardUsecase
}

func NewDashboardHandler(uc domain.DashboardUsecase) *DashboardHandler {
	return &DashboardHandler{uc}
}

func (h *DashboardHandler) Route(r *gin.RouterGroup) {
	v1 := r.Group("/api/v1")
	v1.Use(middleware.AuthJwt, middleware.AuthAdmin)
	{
		v1.GET("/dashboards", h.Get)
	}
}

func (h *DashboardHandler) Get(ctx *gin.Context) {
	data := h.uc.GetDashboard()

	ctx.JSON(http.StatusOK, resp.Response(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		data,
	))
}
