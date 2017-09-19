package rand

import (
	"math/rand"
)

var pinSet = []rune("0123456789")

func RandomPin(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = pinSet[rand.Intn(len(pinSet))]
	}
	return string(b)
}
