// Copyright (c) 2023 thorstenrie
// All rights reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package lpdice

// Import package testing as well as tserr
import (
	"testing" // testing

	"github.com/thorstenrie/lpstats"
	"github.com/thorstenrie/tserr" // tserr
)

// tc contains dice to be tested. It is a slice of structs.
// Each struct holds the number of n sides of the die and
// the function to create the corresponding n-sided Die.
var (
	tc = []struct {
		n int
		f func() (*Die, error)
	}{
		{4, NewD4},   // four-sided die
		{6, NewD6},   // six-sided die
		{8, NewD8},   // eight-sided die
		{10, NewD10}, // ten-sided die
		{12, NewD12}, // twelve-sided die
		{20, NewD20}, // 20-sided die
	}
)

// TestSeed rolls two six-sided seeded dice with the same seed. The test fails
// if the results of the two six-sided dice differ. Afterwards, it sets one of
// both dice to a non-seeded dice. It compares newly generated results from
// rolling the non-seeded die to the seeded die. The test fails, if the results
// of the seeded die to the non-seeded die are the same.
func TestSeed(t *testing.T) {
	var (
		// Roll dice itr times
		itr int = 1000000
		// y1 and y2 hold a slice of results from rolling a die itr times
		y1, y2 []int = make([]int, itr), make([]int, itr)
		// f creates a six-sided die d, sets the same seed for the die and returns the die
		f = func() *Die {
			// Retrieve a six-sided die
			d, e := NewD6()
			// The test fails if NewD6 returns an error
			if e != nil {
				t.Fatal(tserr.NotAvailable(&tserr.NotAvailableArgs{S: "D6", Err: e}))
			}
			// Seed the six-sided die with 1
			if e = d.Seed(1); e != nil {
				// The test fails if Seed returns an error
				t.Fatal(tserr.Op(&tserr.OpArgs{Op: "Seed", Fn: "d", Err: e}))
			}
			// Return the new six-sided die
			return d
		}
		// d1 and d2 hold a six-sided die with the same seed
		d1, d2 = f(), f()
	)
	// Roll die d1 itr times and retrieve results in y1
	itrD(t, d1, y1, itr)
	// Roll die d2 itr times and retrieve results in y2
	itrD(t, d2, y2, itr)
	// Compare results in y1 and y2
	if e := lpstats.EqualS(y1, y2); e != nil {
		// The test fails if EqualS returns an error
		t.Error(tserr.Op(&tserr.OpArgs{Op: "EqualS", Fn: "y1 and y2", Err: e}))
	}
	// Use non-seeded random number generator for d2
	d2.NoSeed()
	// Roll die d2 itr times and retrieve results in y2
	itrD(t, d2, y2, itr)
	// Compare results in y1 and y2
	if e := lpstats.EqualS(y1, y2); e == nil {
		// The test fails if EqualS returns nil
		t.Error(tserr.NilFailed("EqualS"))
	}
}

// TestDice creates dice with 4, 6, 8, 10, 12, 20 sides and rolls them.
// For the results, the arithmetic mean and variance are calculated. The test
// fails if the arithmetic mean and the variance of the results are not
// near equal to expected values.
func TestDice(t *testing.T) {
	// Iterate all testcases in tc
	for _, c := range tc {
		// Create the die to be tested
		d, e := c.f()
		// The test fails, if creating the die returned an error
		if e != nil {
			t.Error(tserr.NotAvailable(&tserr.NotAvailableArgs{S: "Die", Err: e}))
		} else {
			// Roll the die d for a lot of times and calculate the arithmetic mean and variance of
			// the results. If the arithmetic mean or variance do not near equal the exepcted value or
			// expected variance the test fails.
			testD(t, c.n, d)
		}
	}
}

// TestNil checks if all functions of a Die return an error
// if the functions are called for a nil Die. The test fails
// if any function returns nil instead of an error.
func TestNil(t *testing.T) {
	// Define nil Die d
	var d *Die = nil
	// The test fails if init returns nil
	if _, err := d.init(); err == nil {
		t.Error(tserr.NilFailed("init"))
	}
	// The test fails if notSet returns nil
	if err := d.notSet(); err == nil {
		t.Error(tserr.NilFailed("lateInit"))
	}
	// The test fails if Roll returns nil
	if _, err := d.Roll(); err == nil {
		t.Error(tserr.NilFailed("Roll"))
	}
	// The test fails if Seed returns nil
	if err := d.Seed(0); err == nil {
		t.Error(tserr.NilFailed("Seed"))
	}
	// The test fails if NoSeed returns nil
	if err := d.NoSeed(); err == nil {
		t.Error(tserr.NilFailed("NoSeed"))
	}
}

// TestNotSet rolls dice which were not initialized. Therefore it
// rolls standard Dies and stores the results. If the arithmetic mean
// and variance of the results does not near equal the expected value
// and expected variance of a six-sided die, the test fails.
func TestNotSet(t *testing.T) {
	// Define four standard Dies d1, d2, d3, d4
	var d1, d2, d3, d4 Die
	// Test d1
	testD(t, defaultN, &d1)
	// Set d2 to a non-seeded die
	d2.NoSeed()
	// Test d2
	testD(t, defaultN, &d2)
	// Seed d3
	d3.Seed(1)
	// Test d3
	testD(t, defaultN, &d3)
	// Roll d4
	d4.Roll()
	// Test d4
	testD(t, defaultN, &d3)
}
