// Package lpdice provides a simple API for dice with 4, 6, 8, 10, 12 and 20 sides. A new die is retrieved with
// NewD{4, 6, 8, 10, 12, 20}. A directly instantiated Die has 6 sides. The underlying random number generator is a
// cryptographically secure random number generator based on crypto/rand. If the cryptographically secure random
// number generator source is not available on the system,
// a pseudo-random number generator based on math/rand is used. A die can be seeded to generate deterministic random
// numbers based on a seed.
//
// Copyright (c) 2023 thorstenrie
// All rights reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package lpdice

// NewD4 returns a pointer to a four-sided die. It returns nil and an error, if any.
func NewD4() (*Die, error) {
	return newDie(4)
}

// NewD6 returns a pointer to a six-sided die. It returns nil and an error, if any.
func NewD6() (*Die, error) {
	return newDie(6)
}

// NewD8 returns a pointer to an eight-sided die. It returns nil and an error, if any.
func NewD8() (*Die, error) {
	return newDie(8)
}

// NewD10 returns a pointer to a ten-sided die. It returns nil and an error, if any.
func NewD10() (*Die, error) {
	return newDie(10)
}

// NewD12 returns a pointer to a twelve-sided die. It returns nil and an error, if any.
func NewD12() (*Die, error) {
	return newDie(12)
}

// NewD20 returns a pointer to a 20-sided die. It returns nil and an error, if any.
func NewD20() (*Die, error) {
	return newDie(20)
}

// newDie returns a pointer to an n-sided die. It returns nil and an error, if any.
func newDie(n int) (*Die, error) {
	// Create an instance of Die with n sides and return result of function init.
	return (&Die{s: n}).init()
}
