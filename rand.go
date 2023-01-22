package lpdice

import (
	"math/rand"
)

func CryptoRand() (*rand.Rand, error) {
	return New(cryptoSource())
}

func NewDeterministicRand() (*rand.Rand, error) {
	return New(newDeterministicSource())
}

// Seed and Read not secure for sources from NewSource()!
func New(src Source) (*rand.Rand, error) {
	if src.assert(); src.err() != nil {
		// TODO: add context to error
		return nil, src.err()
	}
	return rand.New(src), nil
}
