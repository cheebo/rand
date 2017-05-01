package gorand

import (
	"testing"
	"regexp"
	"fmt"
)

const length = 1024

func TestRandomAlpha(t *testing.T) {
	alpha := RandomAlpha(length)
	var valid = regexp.MustCompile(`[a-zA-Z]+`)
	if len(alpha) != length {
		t.Fail()
	}
	if !valid.MatchString(alpha) {
		t.Fail()
	}
}


func TestRandomAlphaNum(t *testing.T) {
	alphaNum := RandomAlpha(length)
	var valid = regexp.MustCompile(`[a-zA-Z0-9]+`)
	if len(alphaNum) != length {
		t.Fail()
	}
	if !valid.MatchString(alphaNum) {
		t.Fail()
	}
}

func TestRandomString(t *testing.T) {
	str := RandomString(length)
	var valid = regexp.MustCompile(`[a-zA-Z0-9\_\-\=\+\.\,\:\?\!\~\%\$\#\@\&\*]+`)
	if len(str) != length {
		t.Fail()
	}
	if !valid.MatchString(str) {
		t.Fail()
	}
}


func ExampleRandomAlpha() {
	alpha := RandomAlpha(8)
	fmt.Println(alpha)
	// Output: alpha random string
}

func ExampleRandomAlphaNum() {
	alphaNum := RandomAlpha(8)
	fmt.Println(alphaNum)
	// Output: alpha-numerical random string
}

func ExampleRandomString() {
	str := RandomAlpha(8)
	fmt.Println(str)
	// Output: random string with special characters
}