package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz	"

func init() {
	rand.Seed(time.Now().UnixNano())
}

//random integer
func RandomInt(min, max int32) int32 {
	return min + rand.Int31n(max-min+1)
}

//rand string
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

//rand name
func RandomName() string {
	return RandomString(6)
}

//rand int

func RandomNum() int32 {
	return RandomInt(0, 10000)
}

func RandomCategory() string {
	category := []string{"backend", "frontend", "databse", "mobile"}
	n := len(category)
	return category[rand.Intn(n)]
}
