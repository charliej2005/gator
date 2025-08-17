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

func handlerFollow(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("missing url")
	}
	url := cmd.args[0]

	nullTime := sql.NullTime{Time: time.Now(), Valid: true}
	feed, err := s.db.GetFeed(context.Background(), url)
	if err != nil {
		return err
	}
	feedID := feed.ID

	user, err := s.db.GetUser(context.Background(), s.config.CurrentUserName)
	if err != nil {
		return err
	}
	userID := user.ID

	params := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: nullTime,
		UpdatedAt: nullTime,
		UserID:    userID,
		FeedID:    feedID,
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), params)
	if err != nil {
		return err
	}

	fmt.Printf("user '%v' now following '%v'\n", feedFollow.UserName, feedFollow.FeedName)

	return nil
}
