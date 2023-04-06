package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/thorstenrie/lpstats"
)

type CommandFunc func(context.Context, []string) error

type Command struct {
	Key      string
	Help     string
	Function CommandFunc
}

type runner struct {
	help string
	cmds map[string]*Command
	exit *Command
}

var (
	run = runner{cmds: make(map[string]*Command)}
)

func HelpText(text string) error {
	if text != printable(text) {
		return errors.New("only printable characters allowed in help text")
	}
	run.help = text
	return nil
}

func HelpCommand(c string) error {
	return Add(&Command{Key: c, Function: printHelp, Help: "Print usage statement"})
}

func printHelp(ctx context.Context, args []string) error {
	text := ""
	if run.help != "" {
		text += fmt.Sprintf("%s\n\n", run.help)
	}
	if len(run.cmds) == 0 {
		fmt.Println(text)
		return nil
	}
	text += fmt.Sprintf("  Usage:\n    [command] [arguments]")
	text += "\n\n  Available commands:"
	m := 0
	for c := range run.cmds {
		m = lpstats.Max(m, len(c))
	}
	for c := range run.cmds {
		text += "\n    " +
			c +
			strings.Repeat(" ", m+1-len(c)) +
			run.cmds[c].Help
	}
	fmt.Println(text)
	return nil
}

func Add(cmd *Command) error {
	if cmd.Function == nil {
		return errors.New("function cannot be nil")
	}
	if cmd.Key == "" {
		return errors.New("command cannot be empty")
	}
	if cmd.Key != printable(cmd.Key) {
		return errors.New("only printable characters allowed in key")
	}
	if _, e := find(cmd.Key); e == nil {
		return errors.New("command already exists")
	}
	run.cmds[cmd.Key] = cmd
	return nil
}

func split(l string) (string, []string, error) {
	a := strings.Split(printable(l), " ")
	if len(a) == 0 {
		return "", nil, errors.New("empty line")
	}
	if len(a) == 1 {
		return a[0], nil, nil
	}
	return a[0], a[1:], nil
}

func input(ctx context.Context, ch chan string) {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		ch <- s.Text()
		select {
		case <-ctx.Done():
			return
		default:
		}
	}
	if err := s.Err(); err != nil {
		ch <- err.Error()
	}
}

func Run(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	ch := make(chan string)
	go input(ctx, ch)
	for {
		fmt.Printf("< ")
		select {
		case i := <-ch:
			fmt.Printf("> ")
			cmd, args, e := split(i)
			if e != nil {
				fmt.Printf("Error: %s\n", e)
				continue
			}
			c, e := find(cmd)
			if e != nil {
				fmt.Printf("Error: %s\n", e)
				continue
			}
			if err := c.Function(ctx, args); err != nil {
				fmt.Printf("Error: %s\n", err)
			}
			if c == run.exit {
				return
			}
		case <-ctx.Done():
			fmt.Printf("\n> ")
			if err := exit(ctx); err != nil {
				fmt.Printf("Error: %s\n", err)
			}
			return
		}
	}
}

func find(cmd string) (*Command, error) {
	if f, ok := run.cmds[cmd]; ok {
		return f, nil
	}
	return nil, errors.New("command does not exist")
}

func SetExit(cmd string) error {
	c, e := find(cmd)
	if e != nil {
		return e
	}
	run.exit = c
	return nil
}

func exit(ctx context.Context) error {
	if run.exit == nil {
		return errors.New("no exit function")
	}
	return run.exit.Function(ctx, nil)
}
