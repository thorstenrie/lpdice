package main

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/thorstenrie/lpdice"
	"github.com/thorstenrie/lpstats"
)

var history []int

func roll(ctx context.Context, args []string) /*(bool,*/ error /*)*/ {
	if len(args) != 0 {
		return /*true,*/ errors.New("Unexpected argument")
	}
	r, e := d.Roll()
	if e != nil {
		return /*true,*/ e
	}
	history = append(history, r)
	fmt.Printf("%d\n", r)
	return /*true,*/ nil
}

func stop(ctx context.Context, args []string) error {
	if len(args) != 0 {
		return errors.New("Unexpected argument")
	}
	m, _ := lpstats.ArithmeticMean(history)
	fmt.Printf("average = %f\n", m)
	return nil
}

func sides(ctx context.Context, args []string) /*(bool,*/ error /*)*/ {
	if len(args) != 1 {
		return /*true,*/ errors.New("Expected one argument")
	}
	i, e := strconv.ParseInt(args[0], 10, 0)
	if e != nil {
		return /*true,*/ errors.New("Argument must be an integer")
	}
	switch i {
	case 4:
		d, _ = lpdice.NewD4()
	case 6:
		d, _ = lpdice.NewD6()
	case 8:
		d, _ = lpdice.NewD8()
	case 10:
		d, _ = lpdice.NewD10()
	case 12:
		d, _ = lpdice.NewD12()
	case 20:
		d, _ = lpdice.NewD20()
	default:
		return /*true,*/ errors.New("Die has 4, 6, 8, 10, 12 or 20 sides")
	}
	fmt.Printf("new die with %d sides\n", i)
	return /*true,*/ nil
}

func seed(ctx context.Context, args []string) /*(bool,*/ error /*)*/ {
	if len(args) != 1 {
		return /*true,*/ errors.New("Expected one argument")
	}
	i, e := strconv.ParseInt(args[0], 10, 0)
	if e != nil {
		return /*true,*/ errors.New("Argument must be an integer")
	}
	e = d.Seed(i)
	if e != nil {
		return /*true,*/ e
	}
	fmt.Printf("die seeded with %d\n", i)
	return /*true,*/ nil
}
