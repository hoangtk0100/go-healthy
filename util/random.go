package util

import (
	"math/rand"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandomString(n int) string {
	var sb strings.Builder
	length := len(alphabet)

	for index := 0; index < n; index++ {
		c := alphabet[rand.Intn(length)]
		sb.WriteByte(c)
	}

	return sb.String()
}
