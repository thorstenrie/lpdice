package main

import (
	"context"
	"errors"
)

type CommandFunc func(context.Context, []string) error

type command struct {
	name     string
	function CommandFunc
}

func newCommand(name string, function CommandFunc) (*command, error) {
	if name != printable(name) {
		return nil, errors.New("invalid characters in name")
	}
	cmd := &command{name: name, function: function}
	return cmd, nil
}
