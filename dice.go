package lpdice

import (
	"math/rand"
	"sync"

	"github.com/thorstenrie/tserr"
	"github.com/thorstenrie/tsrand"
)

type Dice struct {
	prnd *rand.Rand
	drnd *rand.Rand
	grnd *rand.Rand
	mu   sync.Mutex
}

func New() (*Dice, error) {
	var err error
	d := &Dice{}
	if d.prnd, err = tsrand.NewCryptoRand(); err != nil {
		if d.prnd, err = tsrand.NewPseudoRandomRand(); err != nil {
			return nil, tserr.NotAvailable(&tserr.NotAvailableArgs{S: "tsrand.NewPseudoRandomRand", Err: err})
		}
	}
	if d.drnd, err = tsrand.NewDeterministicRand(); err != nil {
		return nil, tserr.NotAvailable(&tserr.NotAvailableArgs{S: "tsrand.NewDeterministicRand", Err: err})
	}
	d.grnd = d.prnd
	return d, nil
}

var (
	globalDice *Dice
)

func init() {
	var err error
	if globalDice, err = New(); err != nil {
		panic(tserr.NotAvailable(&tserr.NotAvailableArgs{S: "Dice", Err: err}))
	}
}

func (d *Dice) Seed(s int64) error {
	d.mu.Lock()
	if d.drnd == nil {
		d.mu.Unlock()
		return tserr.NilPtr()
	}
	d.drnd.Seed(s)
	d.grnd = d.drnd
	d.mu.Unlock()
	return nil
}

func Seed(s int64) error {
	if globalDice == nil {
		return tserr.NilPtr()
	}
	return globalDice.Seed(s)
}

func (d *Dice) NoSeed() error {
	d.mu.Lock()
	if d.prnd == nil {
		d.mu.Unlock()
		return tserr.NilPtr()
	}
	d.grnd = d.prnd
	d.mu.Unlock()
	return nil
}

func NoSeed() error {
	if globalDice == nil {
		return tserr.NilPtr()
	}
	return globalDice.NoSeed()
}

func (d *Dice) d(sides int) (int, error) {
	d.mu.Lock()
	if d.grnd == nil {
		d.mu.Unlock()
		return 0, tserr.NilPtr()
	}
	r := d.grnd.Intn(sides) + 1
	d.mu.Unlock()
	return r, nil
}

func D4() (int, error) {
	if globalDice == nil {
		return 0, tserr.NilPtr()
	}
	return globalDice.d(4)
}

func D6() (int, error) {
	if globalDice == nil {
		return 0, tserr.NilPtr()
	}
	return globalDice.d(6)
}

func D8() (int, error) {
	if globalDice == nil {
		return 0, tserr.NilPtr()
	}
	return globalDice.d(8)
}

func D10() (int, error) {
	if globalDice == nil {
		return 0, tserr.NilPtr()
	}
	return globalDice.d(10)
}

func D12() (int, error) {
	if globalDice == nil {
		return 0, tserr.NilPtr()
	}
	return globalDice.d(12)
}

func D20() (int, error) {
	if globalDice == nil {
		return 0, tserr.NilPtr()
	}
	return globalDice.d(20)
}
