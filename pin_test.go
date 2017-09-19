package rand_test

import (
	"testing"
	"regexp"
	"github.com/cheebo/rand"
)

const length = 8

func TestRandomPin(t *testing.T) {
	alpha := rand.RandomPin(length)
	var valid = regexp.MustCompile(`[0-9]+`)
	if len(alpha) != length {
		t.Fail()
	}
	if !valid.MatchString(alpha) {
		t.Fail()
	}
}
