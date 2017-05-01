package gorand

import (
	"testing"
	"regexp"
)

const length = 8

func TestRandomPin(t *testing.T) {
	alpha := RandomPin(length)
	var valid = regexp.MustCompile(`[0-9]+`)
	if len(alpha) != length {
		t.Fail()
	}
	if !valid.MatchString(alpha) {
		t.Fail()
	}
}
