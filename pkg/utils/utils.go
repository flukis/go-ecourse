package utils

import (
	"e-course/pkg/resp"
	"errors"
	"math/rand"

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
