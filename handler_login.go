package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

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
