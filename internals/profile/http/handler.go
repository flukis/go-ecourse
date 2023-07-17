package handler

import (
	"e-course/domain"
	"e-course/internals/middleware"
	"e-course/pkg/resp"
	"e-course/pkg/utils"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type ProfileHandler struct {
	uc domain.ProfileUsecase
}

func NewProfileHandler(uc domain.ProfileUsecase) *ProfileHandler {
	return &ProfileHandler{uc}
}

func (h *ProfileHandler) Route(r *gin.RouterGroup) {
	v1 := r.Group("/api/v1")

	v1.Use(middleware.AuthJwt)
	{
		v1.GET("/profile/:id", h.Profile)
		v1.POST("/profile/logout", h.Logout)
		v1.PATCH("/profile/:id", h.Update)
		v1.DELETE("/profile/:id", h.Deactive)
	}
}

func (h *ProfileHandler) Profile(ctx *gin.Context) {
	user := utils.GetCurrentUser(ctx)

	data, err := h.uc.FindProfile(int(user.ID))
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

func (h *ProfileHandler) Logout(ctx *gin.Context) {
	var header domain.ProfileRequestLogoutBody
	ctx.ShouldBindHeader(&header)

	reqToken := header.Authorization
	splittedToken := strings.Split(reqToken, "Bearer ")

	err := h.uc.Logout(splittedToken[1])
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

func (h *ProfileHandler) Deactive(ctx *gin.Context) {
	user := utils.GetCurrentUser(ctx)

	err := h.uc.Deactive(int(user.ID))
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

func (h *ProfileHandler) Update(ctx *gin.Context) {
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
