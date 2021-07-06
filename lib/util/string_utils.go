package util

import (
	"github.com/Authing/authing-go-sdk/lib/constant"
	"math/rand"
)

var letters = []rune("abcdefhijkmnprstwxyz2345678")

func RandomString(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func GetValidValue(value ...string) string {
	if value == nil || len(value) == 0 {
		return constant.StringEmpty
	}
	for _, val := range value {
		if val != "" {
			return val
		}
	}
	return constant.StringEmpty
}
