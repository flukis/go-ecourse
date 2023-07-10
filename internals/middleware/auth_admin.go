package middleware

import (
	"e-course/pkg/resp"
	"e-course/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthAdmin(ctx *gin.Context) {
	admin := utils.GetCurrentUser(ctx)

	if !admin.IsAdmin {
		ctx.JSON(http.StatusUnauthorized, resp.Response(
			http.StatusUnauthorized,
			http.StatusText(http.StatusUnauthorized), "unauthorized",
		))

		ctx.Abort()
		return
	}

	ctx.Next()
}
