package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/charliej2005/gator/internal/config"
	"github.com/charliej2005/gator/internal/database"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

func main() {
	cfg, _ := config.Read()

	db, err := sql.Open("postgres", cfg.DbURL)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	dbQueries := database.New(db)

	s := state{db: dbQueries, config: &cfg}

	cmds := commands{
		handler: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)

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
	err = cmds.run(&s, userCmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

type state struct {
	db     *database.Queries
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

	_, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("no user with the given name")
		}
		return err
	}

	err = s.config.SetUser(username)
	if err != nil {
		return err
	}
	fmt.Printf("The user has been set to '%s'\n", username)
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("username not provided")
	}
	username := cmd.args[0]
	nullTime := sql.NullTime{Time: time.Now(), Valid: true}
	userParams := database.CreateUserParams{
		ID:        uuid.NullUUID{UUID: uuid.New(), Valid: true},
		CreatedAt: nullTime,
		UpdatedAt: nullTime,
		Name:      username,
	}

	_, err := s.db.GetUser(context.Background(), username)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	if err == nil {
		return errors.New("user with that name already exists")
	}

	_, err = s.db.CreateUser(context.Background(), userParams)
	if err != nil {
		return err
	}

	err = s.config.SetUser(username)
	if err != nil {
		return err
	}

	fmt.Printf("user has been created:\nid: %s\ncreated at: %s\nupdated at: %s\nname: %s\n",
		userParams.ID.UUID, userParams.CreatedAt.Time, userParams.UpdatedAt.Time, userParams.Name)

	return nil
}
