package gorand

import (
	"time"
	"math/rand"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}
