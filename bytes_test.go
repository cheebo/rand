package gorand

import "testing"

const length = 10

func TestGenerateRandomBytes(t *testing.T) {
	b, err := GenerateRandomBytes(length)
	if err != nil {
		t.Fail()
	}
	if len(b) != length {
		t.Fail()
	}
}
