package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
)

var (
	cmds  map[string]*command = make(map[string]*command)
	exitc *command
)

func register(key string, f CommandFunc) error {
	if _, e := find(key); e == nil {
		return errors.New("command already exists")
	}
	cmd, e := newCommand(key, f)
	if e != nil {
		return e
	}
	cmds[key] = cmd
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

func run(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	ch := make(chan string)
	go input(ctx, ch)
	for {
		fmt.Printf("Input: ")
		select {
		case i := <-ch:
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
			fmt.Printf("Output: ")
			if err := c.function(ctx, args); err != nil {
				fmt.Printf("\nError: %s\n", err)
			}
			if c == exitc {
				return
			}
		case <-ctx.Done():
			fmt.Printf("\nOutput: ")
			if err := Exit(ctx); err != nil {
				fmt.Printf("\nError: %s\n", err)
			}
			return
		}
	}
}

func find(cmd string) (*command, error) {
	if f, ok := cmds[cmd]; ok {
		return f, nil
	}
	return nil, errors.New("command does not exist")
}

func setExit(cmd string) error {
	c, e := find(cmd)
	if e != nil {
		return e
	}
	exitc = c
	return nil
}

func Exit(ctx context.Context) error {
	if exitc == nil {
		return errors.New("no exit function")
	}
	return exitc.function(ctx, nil)
}
