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

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.args) < 2 {
		return errors.New("missing name and/or url")
	}

	user, err := s.db.GetUser(context.Background(), s.config.CurrentUserName)
	if err != nil {
		return err
	}

	name := cmd.args[0]
	nullTime := sql.NullTime{Time: time.Now(), Valid: true}
	url := cmd.args[1]
	userID := user.ID

	params := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: nullTime,
		UpdatedAt: nullTime,
		Name:      name,
		Url:       url,
		UserID:    userID,
	}

	_, err = s.db.CreateFeed(context.Background(), params)
	if err != nil {
		return err
	}

	fmt.Println("feed added successfully")
	return nil
}
