package lpdice

import (
	"testing"

	"github.com/thorstenrie/tserr"
)

var (
	tc = []struct {
		n int
		f func() (int, error)
	}{
		{4, D4},
		{6, D6},
		{8, D8},
		{10, D10},
		{12, D12},
		{20, D20},
	}
)

func TestSeed(t *testing.T) {
	var (
		a, b int
		e    error
		s    int64               = 1
		itr  int                 = 100
		f    func() (int, error) = func() (int, error) { Seed(s); return D6() }
	)
	if a, e = f(); e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "f()", Fn: "a", Err: e}))
	}
	for i := 0; i < itr; i++ {
		if b, e = f(); e != nil {
			t.Error(tserr.Op(&tserr.OpArgs{Op: "f()", Fn: "b", Err: e}))
		}
		if a != b {
			t.Error(tserr.Equal(&tserr.EqualArgs{Var: "D6", Actual: int64(b), Want: int64(a)}))
		}
	}
	NoSeed()
}

func TestDices(t *testing.T) {
	for _, c := range tc {
		testD(t, c)
	}
}
