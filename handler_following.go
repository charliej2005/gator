package main

import (
	"context"
	"errors"
	"fmt"
)

func handlerFollowing(s *state, cmd command) error {

	user, err := s.db.GetUser(context.Background(), s.config.CurrentUserName)
	if err != nil {
		return err
	}
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
