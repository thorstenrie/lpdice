package lpdice

import (
	"fmt"
	"testing"
)

func TestD6(t *testing.T) {
	fmt.Println(D6())
	fmt.Println(D6())
	fmt.Println(D6())
	fmt.Println(D6())
	fmt.Println(D6())
}

func TestD6s(t *testing.T) {
	fmt.Println(D6())
	Seed(1)
	fmt.Println(D6())
	fmt.Println(D6())
	Seed(1)
	fmt.Println(D6())
	fmt.Println(D6())
	Seed(1)
	fmt.Println(D6())
	fmt.Println(D6())
	Seed(2)
	fmt.Println(D6())
	fmt.Println(D6())
}
