package main

import (
	"fmt"

	"github.com/charliej2005/gator/internal/config"
)

func main() {
	// placeholder main that accesses the config and updates it to my name :)
	cfg, _ := config.Read()
	fmt.Printf("dbUrl: %v, username: %v\n", cfg.DbURL, cfg.CurrentUserName)
	username := "charlie"
	fmt.Println("setting username")
	_ = cfg.SetUser(username)
	cfg, _ = config.Read()
	fmt.Printf("dbUrl: %v, username: %v\n", cfg.DbURL, cfg.CurrentUserName)
	// implement proper functionality later
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
