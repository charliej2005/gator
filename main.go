package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/charliej2005/gator/internal/config"
)

func main() {
	// placeholder main that accesses the config and updates it to my name :)
	cfg, _ := config.Read()
	s := state{config: &cfg}
	cmds := commands{
		handler: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	args := os.Args
	if len(args) < 2 {
		err := errors.New("command not provided")
		fmt.Println(err)
		os.Exit(1)
	}
	userCmdName := args[1]
	var userArgs []string
	if len(args) > 2 {
		userArgs = args[2:]
	}
	userCmd := command{name: userCmdName, args: userArgs}
	err := cmds.run(&s, userCmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

type state struct {
	config *config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	handler map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	err := c.handler[cmd.name](s, cmd)
	if err != nil {
		return err
	}
	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.handler[name] = f
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("username not provided")
	}
	username := cmd.args[0]
	err := s.config.SetUser(username)
	if err != nil {
		return err
	}
	fmt.Printf("The user has been set to '%s'\n", username)
	return nil
}
