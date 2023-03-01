// Copyright (c) 2023 thorstenrie
// All rights reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package lpdice

// Import standard library package math/rand as well as tserr and tsrand
import (
	"math/rand" // rand

	"github.com/thorstenrie/tserr"  // tserr
	"github.com/thorstenrie/tsrand" // tsrand
)

// defaultN defines the default no. sides of a directly instantiated Die
const (
	defaultN int = 6
)

// A Die holds random number generators to roll a die with s sides. It contains a pointer to a pseudo-random number generator in prnd
// and a pointer to a deterministic random number generator in drnd. The currently used random number generator is stored in grnd.
type Die struct {
	prnd *rand.Rand // Pseudo-random number generator
	drnd *rand.Rand // Deterministic pseudo-random number generator
	grnd *rand.Rand // Currently used random number generator, either prnd or drnd
	s    int        // number of sides
}

// init initializes a Die. It returns a pointer to the die and an error, if any.
// A pointer to a new cryptographically secure random number generator will be stored in prnd. If it
// is not available on the platform, a pseudo-random number generator will be used. A pointer to a
// new deterministic pseudo-random number generator will be stored in drnd. The currently used
// random number generator grnd is set to prnd.
func (d *Die) init() (*Die, error) {
	// Return an error if d is nil
	if d == nil {
		return nil, tserr.NilPtr()
	}
	// err holds the error of creating a random number generator, if any
	var err error
	// Retrieve a new cryptographically secure random number generator in prnd
	if d.prnd, err = tsrand.NewCryptoRand(); err != nil {
		// Retrieve a new pseudo-random number generator in prnd, if the cryptographically secure random number generator is not available on the platform
		if d.prnd, err = tsrand.NewPseudoRandomRand(); err != nil {
			// Return nil and an error if any
			return nil, tserr.NotAvailable(&tserr.NotAvailableArgs{S: "tsrand.NewPseudoRandomRand", Err: err})
		}
	}
	// Retrieve a new deterministic random number generator in drnd
	if d.drnd, err = tsrand.NewDeterministicRand(); err != nil {
		// Return nil and an error if any
		return nil, tserr.NotAvailable(&tserr.NotAvailableArgs{S: "tsrand.NewDeterministicRand", Err: err})
	}
	// Set the currently used random number generator to prnd
	d.grnd = d.prnd
	// Return d and nil
	return d, nil
}

// notSet sets the number of sides of Die d to the default number of sides defaultN
// if the Die has not been initialized. In this case, it also initializes the Die d.
func (d *Die) notSet() error {
	// Return an error if d is nil
	if d == nil {
		return tserr.NilPtr()
	}
	// If s is zero the Die d is not initialized yet
	if d.s == 0 {
		// Set the number of sides s to the default number of sides defaultN
		d.s = defaultN
		// Initialize the die
		_, e := d.init()
		// Return an error if any
		return e
	}
	// Return nil if the die is already initialized
	return nil
}

// Roll returns the result of rolling the die. It returns zero and an error, if any.
// For rolling the dice, the currently set random number generator in grnd is used.
func (d *Die) Roll() (int, error) {
	// Return an error if d is nil
	if d == nil {
		return 0, tserr.NilPtr()
	}
	// Initialize the die if not initialized yet
	if e := d.notSet(); e != nil {
		// Return zero and an error if the initialization fails
		return 0, e
	}
	// Return zero and an error if grnd is nil
	if d.grnd == nil {
		return 0, tserr.NilPtr()
	}
	// Return the result of rolling the die through grnd
	return d.grnd.Intn(d.s) + 1, nil
}

// Seed seeds the die with seed s. Therefore it sets the currently used random number generator grnd
// to the deterministic random number generator drnd, which will be seeded with s.
func (d *Die) Seed(s int64) error {
	// Return an error if d is nil
	if d == nil {
		return tserr.NilPtr()
	}
	// Initialize the die if not initialized yet
	if e := d.notSet(); e != nil {
		// Return an error if the initialization fails
		return e
	}
	// Return an error if drnd is nil
	if d.drnd == nil {
		return tserr.NilPtr()
	}
	// Seed the deterministic random number generator drnd
	d.drnd.Seed(s)
	// Set the currently used random number generator grnd to drnd
	d.grnd = d.drnd
	// Return nil
	return nil
}

// NoSeed sets the die to the cryptographically secure random number generator or pseudo-random number generator,
// if the cryptographically secure random number generator is not available on the platform. NoSeed is used
// if a seeded die should be changed back to a non-seeded die.
func (d *Die) NoSeed() error {
	// Return an error if d is nil
	if d == nil {
		return tserr.NilPtr()
	}
	// Initialize the die if not initialized yet
	if e := d.notSet(); e != nil {
		// Return an error if the initialization fails
		return e
	}
	// Return an error if prnd is nil
	if d.prnd == nil {
		return tserr.NilPtr()
	}
	// Set the currently used random number generator grnd to prnd
	d.grnd = d.prnd
	// Return nil
	return nil
}
