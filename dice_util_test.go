// Copyright (c) 2023 thorstenrie
// All rights reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package lpdice

// Import testing as well as lpstats and tserr
import (
	"testing" // testing

	"github.com/thorstenrie/lpstats" // lpstats
	"github.com/thorstenrie/tserr"   // tserr
)

// itrD rolls die d itr times and retrieves the result in y.
func itrD(t *testing.T, d *Die, y []int, itr int) {
	var e error
	// Roll the die for itr times
	for i := 0; i < itr; i++ {
		// Retrieve result of rolling the die in slice y
		if y[i], e = d.Roll(); e != nil {
			// The test fails if rolling the die returns an error
			t.Error(tserr.Op(&tserr.OpArgs{Op: "Roll", Fn: "y", Err: e}))
		}
	}
}

// testD rolls n-sided Die d for a lot of times and calculates the arithmetic mean and variance of
// the results. If the arithmetic mean or variance do not near equal the exepcted value or
// expected variance the test fails.
func testD(t *testing.T, n int, d *Die) {
	var (
		// e holds the error of rolling the die if any
		e error
		// The n-sided Die d is rolled itr times
		itr int = 1000000
		// maxDiff holds the maximum difference of near equal floats
		maxDiff float64 = 0.1
		// y holds the slice of results from rolling the die itr times
		y []int = make([]int, itr)
	)
	// Roll the die d for itr times and retrieve the results in y
	itrD(t, d, y, itr)
	// Calculate the arithmetic mean of the results from rolling the die
	mean, e := lpstats.ArithmeticMean(y)
	// The test fails if arithmetic mean has an error
	if e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "ArithmeticMean", Fn: "y", Err: e}))
	}
	// Calculate the expected value for the n-sided die
	meane := lpstats.ExpectedValueU(1, n)
	// The test fails if the arithmetic mean does not equal the expected value with a maximum difference of maxDiff
	if !lpstats.NearEqual(mean, meane, maxDiff) {
		t.Error(tserr.Equalf(&tserr.EqualfArgs{Var: "arithmetic mean of y", Actual: mean, Want: meane}))
	}
	// Calculate the variance of the results from rolling the die
	vari, e := lpstats.Variance(y)
	// The test fails if variance returns an error
	if e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "variance", Fn: "y", Err: e}))
	}
	// Calculate the expected variance
	varie := lpstats.VarianceN(uint(n))
	// The test fails if the variance does not equal the expected variance with a maximum difference of maxDiff
	if !lpstats.NearEqual(vari, varie, maxDiff) {
		t.Error(tserr.Equalf(&tserr.EqualfArgs{Var: "variance of y", Actual: vari, Want: varie}))
	}
}
