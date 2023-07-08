package middleware

import (
	"e-course/domain"
	"e-course/pkg/resp"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Header struct {
	Authorization string `header:"authorization" binding:"required"`
}

func AuthJwt(ctx *gin.Context) {
	var input Header

	if err := ctx.ShouldBindHeader(&input); err != nil {
		ctx.JSON(http.StatusUnauthorized, resp.Response(
			http.StatusUnauthorized,
			http.StatusText(http.StatusUnauthorized),
			"unauthorized",
		))
		ctx.Abort()
		return
	}

	reqToken := input.Authorization
	splittedToken := strings.Split(reqToken, "Bearer ")

	if len(splittedToken) != 2 {
		ctx.JSON(http.StatusUnauthorized, resp.Response(
			http.StatusUnauthorized,
			http.StatusText(http.StatusUnauthorized),
			"unauthorized",
		))
		ctx.Abort()
		return
	}

	reqToken = splittedToken[1]
	claims := &domain.MapClaimResponse{}

	token, err := jwt.ParseWithClaims(reqToken, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, resp.Response(
			http.StatusUnauthorized,
			http.StatusText(http.StatusUnauthorized),
			"unauthorized",
		))
		ctx.Abort()
		return
	}

	if !token.Valid {
		ctx.JSON(http.StatusUnauthorized, resp.Response(
			http.StatusUnauthorized,
			http.StatusText(http.StatusUnauthorized),
			"unauthorized",
		))
		ctx.Abort()
		return
	}

	claims = token.Claims.(*domain.MapClaimResponse)
	ctx.Set("user", claims)
	ctx.Next()
}
