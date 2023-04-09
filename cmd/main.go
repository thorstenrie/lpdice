package main

import (
	"context"

	"github.com/thorstenrie/lpdice"
)

func main() {

	d, _ = lpdice.NewD6()

	AppName("dice")
	HelpText("Throw a die")
	Version("0.1.0")
	HelpCommand("help")
	Add(&Command{Key: "roll", Function: roll, Help: "Roll the die"})
	Add(&Command{Key: "sides", Function: sides, Help: "New die with {4, 6, 8, 10, 12, 20} sides and no seed"})
	Add(&Command{Key: "seed", Function: seed, Help: "Set seed"})
	Add(&Command{Key: "stop", Function: stop, Help: "Exit application"})
	SetExit("stop")

	ctx := context.Background()
	//ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	Run(ctx)
	//cancel()
}
