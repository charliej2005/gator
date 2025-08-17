package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/charliej2005/gator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
	userID := user.ID

	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), userID)
	if err != nil {
		return err
	}

	if len(feedFollows) == 0 {
		return errors.New("no followed feeds found")
	}

	for i := range feedFollows {
		fmt.Printf("* %v\n", feedFollows[i].FeedName)
	}

	return nil
}
