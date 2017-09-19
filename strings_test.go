package rand

import (
	"testing"
	"regexp"
	"github.com/cheebo/rand"
)

const length = 1024

func TestRandomAlpha(t *testing.T) {
	alpha := rand.RandomAlpha(length)
	var valid = regexp.MustCompile(`[a-zA-Z]+`)
	if len(alpha) != length {
		t.Fail()
	}
	if !valid.MatchString(alpha) {
		t.Fail()
	}
}


func TestRandomAlphaNum(t *testing.T) {
	alphaNum := rand.RandomAlpha(length)
	var valid = regexp.MustCompile(`[a-zA-Z0-9]+`)
	if len(alphaNum) != length {
		t.Fail()
	}
	if !valid.MatchString(alphaNum) {
		t.Fail()
	}
}

func TestRandomString(t *testing.T) {
	str := rand.RandomString(length)
	var valid = regexp.MustCompile(`[a-zA-Z0-9\_\-\=\+\.\,\:\?\!\~\%\$\#\@\&\*]+`)
	if len(str) != length {
		t.Fail()
	}
	if !valid.MatchString(str) {
		t.Fail()
	}
}