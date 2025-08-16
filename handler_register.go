package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/charliej2005/gator/internal/database"
	"github.com/google/uuid"
)

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
