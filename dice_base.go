package lpdice

import (
	"math/rand"

	"github.com/thorstenrie/tserr"
	"github.com/thorstenrie/tsrand"
)

type Die struct {
	prnd  *rand.Rand
	drnd  *rand.Rand
	grnd  *rand.Rand
	sides int
}

func (d *Die) init() (*Die, error) {
	if d == nil {
		return nil, tserr.NilPtr()
	}
	var err error
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

func (d *Die) notSet() error {
	if d == nil {
		return tserr.NilPtr()
	}
	if d.sides <= 0 {
		return tserr.NotSet("Sides of die")
	} else if (d.prnd == nil) || (d.drnd == nil) || (d.grnd == nil) {
		return tserr.NilPtr()
	}
	return nil
}

func (d *Die) Roll() (int, error) {
	if d == nil {
		return 0, tserr.NilPtr()
	}
	if e := d.notSet(); e != nil {
		return 0, e
	}
	return d.grnd.Intn(d.sides) + 1, nil
}

func (d *Die) Seed(s int64) error {
	if d == nil {
		return tserr.NilPtr()
	}
	if e := d.notSet(); e != nil {
		return e
	}
	d.drnd.Seed(s)
	d.grnd = d.drnd
	return nil
}

func (d *Die) NoSeed() error {
	if d == nil {
		return tserr.NilPtr()
	}
	if e := d.notSet(); e != nil {
		return e
	}
	d.grnd = d.prnd
	return nil
}
