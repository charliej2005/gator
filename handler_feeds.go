package main

import (
	"context"
	"errors"
	"fmt"
)

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	if len(feeds) == 0 {
		return errors.New("no feeds found")
	}

	for _, feed := range feeds {
		name := feed.Name
		url := feed.Url

		userID := feed.UserID
		user, err := s.db.GetUserFromID(context.Background(), userID)
		if err != nil {
			return err
		}

		username := user.Name

		fmt.Printf("name: '%v' | url: '%v' | creator: '%v'\n", name, url, username)
	}

	return nil
}
