package rand_test

import (
	"testing"
	"github.com/cheebo/rand"
)

const length = 10

func TestGenerateRandomBytes(t *testing.T) {
	b, err := rand.GenerateRandomBytes(length)
	if err != nil {
		t.Fail()
	}
	if len(b) != length {
		t.Fail()
	}
}
