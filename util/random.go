package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
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

func RandomEmail() string {
	return fmt.Sprintf("%s@%s.%s", RandomString(6), RandomString(8), RandomString(3))
}

func RandomPhoneNumber() string {
	rand.Seed(time.Now().UnixNano())
	areaCode := rand.Intn(999-1) + 1 // Random digits between 1 and 999
	numbers := rand.Intn(1000000000) // Random digits between 0 and 1000000000
	return fmt.Sprintf("%d-%d", areaCode, numbers)
}
