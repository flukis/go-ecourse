package utils

import (
	"e-course/domain"
	"e-course/pkg/resp"
	"errors"
	"math/rand"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

func RandString(length int) string {
	var letterRune = []rune("abcdefghijklmnopqrstABCDEFGHIJKLMNOPQRSTUVWXYZ123456789")
	b := make([]rune, length)
	for i := range b {
		b[i] = letterRune[rand.Intn(len(letterRune))]
	}
	return string(b)
}

func IsErrorNot404(err *resp.ErrorResp) bool {
	return err != nil && !errors.Is(err.Err, gorm.ErrRecordNotFound)
}

func GenerateRefreshToken() (ulid.ULID, error) {
	entropy := rand.New(rand.NewSource(time.Now().UnixNano()))
	ms := ulid.Timestamp(time.Now())
	return ulid.New(ms, entropy)
}

func Paginate(offset, limit int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page := offset
		if page <= 0 {
			page = 1
		}
		pageSize := limit
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset = (page - 1) * limit
		return db.Offset(offset).Limit(pageSize)
	}
}

func GetCurrentUser(ctx *gin.Context) *domain.MapClaimResponse {
	user, _ := ctx.Get("user")
	return user.(*domain.MapClaimResponse)
}

func GetFileName(filename string) string {
	file := filepath.Base(filename)

	return file[:len(file)-len(filepath.Ext(file))]
}
