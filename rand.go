package lpdice

import (
	"math/rand"
)

var (
	rnd *rand.Rand
	cs  csource
	ms  msource
)

func init() {
	if cs.available() {
		rnd = rand.New(cs)
	} else {
		rnd = rand.New(ms)
	}
}

func Seed(s int64) {
	ms.Seed(s)
	rnd = rand.New(ms)
}

func d(sides int) int {
	return rnd.Intn(sides) + 1
}
