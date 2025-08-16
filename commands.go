package main

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
