package main

import (
	"context"

	"github.com/thorstenrie/lpdice"
)

var (
	d *lpdice.Die
)

func main() {

	d, _ = lpdice.NewD6()

	register("roll", roll)
	register("sides", sides)
	register("seed", seed)
	register("stop", stop)
	setExit("stop")

	ctx := context.Background()
	//ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	run(ctx)
	//cancel()
}
