package main

import "errors"

type command struct {
	name      string
	arguments []string
}

type commands struct {
	cmdMap map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {

	_, ok := c.cmdMap[cmd.name]
	if !ok {
		return errors.New("command not found")
	}

	err := c.cmdMap[cmd.name](s, cmd)

	return err
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.cmdMap[name] = f
}
