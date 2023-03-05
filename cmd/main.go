package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/thorstenrie/lpdice"
)

func main() {

	var d *lpdice.Die

	d, _ = lpdice.NewD6()

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		switch s.Text() {
		case "roll":
			r, _ := d.Roll()
			fmt.Println(r)
		case "stop":
			os.Exit(0)
		default:
			fmt.Println("unknown")
		}
	}

	if err := s.Err(); err != nil {
		log.Println(err)
	}

}
