package lpdice

import (
	"math/rand"
	"time"
)

var (
	prnd *rand.Rand
	drnd *rand.Rand
	grnd *rand.Rand
)

func init() {
	var err error
	if prnd, err = CryptoRand(); err != nil {
		prnd, _ = NewDeterministicRand()
		prnd.Seed(time.Now().UnixNano())
	}
	drnd, _ = NewDeterministicRand()
	grnd = prnd
}

func Seed(s int64) {
	drnd.Seed(s)
	grnd = drnd
}

func d(sides int) int {
	return grnd.Intn(sides) + 1
}
