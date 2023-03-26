package main

import (
	"strings"
	"unicode"
)

func printable(a string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsPrint(r) {
			return r
		}
		return -1
	}, a)
}
