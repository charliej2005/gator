package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/charliej2005/gator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) == 0 {
		return errors.New("missing url")
	}
	url := cmd.args[0]

	params := database.DeleteFeedFollowParams{
		UserID: user.ID,
		Url:    url,
	}

	err := s.db.DeleteFeedFollow(context.Background(), params)
	if err != nil {
		return err
	}

	fmt.Println("feed unfollowed successfully")
	return nil
}
