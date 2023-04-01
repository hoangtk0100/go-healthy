package util

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func init() {
	// seed the random number generator
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int32) int32 {
	return min + rand.Int31n(max-min+1)
}

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
	// For unique email in current time
	timestamp := strconv.FormatInt(time.Now().UnixNano(), 10)
	return fmt.Sprintf("%s+%s@email.com", RandomString(6), timestamp)
}

func RandomPhoneNumber() string {
	rand.Seed(time.Now().UnixNano())
	areaCode := rand.Intn(999-1) + 1 // Random digits between 1 and 999
	numbers := rand.Intn(1000000000) // Random digits between 0 and 1000000000
	return fmt.Sprintf("%d-%d", areaCode, numbers)
}

func RandomAmount() int32 {
	return RandomInt(1, 400)
}
