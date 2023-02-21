package lpdice

import (
	"testing"

	"github.com/thorstenrie/tserr"
)

var (
	tc = []struct {
		n int
		f func() (*Die, error)
	}{
		{4, NewD4},
		{6, NewD6},
		{8, NewD8},
		{10, NewD10},
		{12, NewD12},
		{20, NewD20},
	}
)

func TestSeed(t *testing.T) {
	var (
		a, b int
		s    int64 = 1
		itr  int   = 100
		f          = func(d *Die) (int, error) { d.Seed(s); return d.Roll() }
	)
	d, e := NewD6()
	if e != nil {
		t.Fatal(tserr.NotAvailable(&tserr.NotAvailableArgs{S: "D6", Err: e}))
	}
	if a, e = f(d); e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "f()", Fn: "a", Err: e}))
	}
	for i := 0; i < itr; i++ {
		if b, e = f(d); e != nil {
			t.Error(tserr.Op(&tserr.OpArgs{Op: "f()", Fn: "b", Err: e}))
		}
		if a != b {
			t.Error(tserr.Equal(&tserr.EqualArgs{Var: "D6", Actual: int64(b), Want: int64(a)}))
		}
	}
	d.NoSeed()
	testD(t, 6, d)
}

func TestDice(t *testing.T) {
	for _, c := range tc {
		d, e := c.f()
		if e != nil {
			t.Error(tserr.NotAvailable(&tserr.NotAvailableArgs{S: "Die", Err: e}))
		} else {
			testD(t, c.n, d)
		}
	}
}

func TestNil(t *testing.T) {
	var d *Die = nil
	if _, err := d.init(); err == nil {
		t.Error(tserr.NilFailed("init"))
	}
	testE(t, d)
}

func TestNotSet(t *testing.T) {
	var d Die
	testE(t, &d)
}
