package rand

import (
	"math/rand"
)

var stringSet = []rune("0123456789ABCDEFGHIJKLMNOPQRSTUFVXYZabcdefghijklmnopqrstuvwxyz_-=+.,:?!~%$#@&*")
var alphaNumSet = []rune("0123456789ABCDEFGHIJKLMNOPQRSTUFVXYZabcdefghijklmnopqrstuvwxyz")
var alphaSet = []rune("ABCDEFGHIJKLMNOPQRSTUFVXYZabcdefghijklmnopqrstuvwxyz")

func RandomAlphaNum(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = alphaNumSet[rand.Intn(len(alphaNumSet))]
	}
	return string(b)
}

func RandomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = stringSet[rand.Intn(len(stringSet))]
	}
	return string(b)
}

func RandomAlpha(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = alphaSet[rand.Intn(len(alphaSet))]
	}
	return string(b)
}

