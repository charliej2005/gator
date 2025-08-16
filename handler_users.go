package main

import (
	"context"
	"errors"
	"fmt"
)

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}
	if len(users) == 0 {
		return errors.New("no users found")
	}
	for _, user := range users {
		if user.Name == s.config.CurrentUserName {
			fmt.Printf("* %v (current)\n", user.Name)
		} else {
			fmt.Printf("* %v\n", user.Name)
		}
	}
	return nil
}
