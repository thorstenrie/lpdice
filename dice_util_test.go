package lpdice

import (
	"testing"

	"github.com/thorstenrie/lpstats"
	"github.com/thorstenrie/tserr"
)

func testD(
	t *testing.T,
	tc struct {
		n int
		f func() (int, error)
	},
) {
	var (
		e       error
		itr     int     = 1000000
		maxDiff float64 = 0.1
		y       []int   = make([]int, itr)
	)
	for i := 0; i < itr; i++ {
		if y[i], e = tc.f(); e != nil {
			t.Error(tserr.Op(&tserr.OpArgs{Op: "tc.f()", Fn: "y", Err: e}))
		}
	}
	// Calculate the arithmetic mean of the random integers
	mean, e := lpstats.ArithmeticMean(y)
	// The test fails if arithmetic mean has an error
	if e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "ArithmeticMean", Fn: "y", Err: e}))
	}
	// Calculate the expected value for the interval [0, testIntn-1]
	meane := lpstats.ExpectedValueU(1, tc.n)
	// The test fails if the arithmetic mean does not equal the expected value with a maximum difference of maxDiff
	if !lpstats.NearEqual(mean, meane, maxDiff) {
		t.Error(tserr.Equalf(&tserr.EqualfArgs{Var: "arithmetic mean of y", Actual: mean, Want: meane}))
	}
	// Calculate the variance of the random integers
	vari, e := lpstats.Variance(y)
	// The test fails if variance returns an error
	if e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "variance", Fn: "y", Err: e}))
	}
	// Calculate the expected variance
	varie := lpstats.VarianceN(uint(tc.n))
	// The test fails if the variance does not equal the expected variance with a maximum difference of maxDiff
	if !lpstats.NearEqual(vari, varie, maxDiff) {
		t.Error(tserr.Equalf(&tserr.EqualfArgs{Var: "variance of y", Actual: vari, Want: varie}))
	}
}
