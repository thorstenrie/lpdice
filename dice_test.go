package lpdice

import (
	"fmt"
	"testing"
)

func TestD6(t *testing.T) {
	r := D6()
	fmt.Println(r)
}

func TestD6s(t *testing.T) {
	Seed(2)
	r := D6()
	fmt.Println(r)
}
