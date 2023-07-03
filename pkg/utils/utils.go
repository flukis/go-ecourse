package utils

import (
	"e-course/pkg/resp"
	"errors"
	"math/rand"

	"gorm.io/gorm"
)

func RandString(length int) string {
	var letterRune = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, length)
	for i := range b {
		b[i] = letterRune[rand.Intn(len(letterRune))]
	}
	return string(b)
}

func IsErrorNot404(err *resp.ErrorResp) bool {
	return err != nil && !errors.Is(err.Err, gorm.ErrRecordNotFound)
}
