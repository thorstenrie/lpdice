package lpdice

import (
	"math/rand"
	"time"
)

type msource struct{}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func (m msource) Seed(s int64) {
	rand.Seed(s)
}

func (m msource) Int63() int64 {
	return rand.Int63()
}
