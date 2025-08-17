package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("time between requests not provided")
	}
	time_between_reqs, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Collecting feeds every %s\n", time_between_reqs)

	ticker := time.NewTicker(time_between_reqs)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *state) error {
	dbFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	err = s.db.MarkFeedFetched(context.Background(), dbFeed.ID)
	if err != nil {
		return err
	}

	feed, err := fetchFeed(context.Background(), dbFeed.Url)
	if err != nil {
		return err
	}

	fmt.Printf("%v:\n", feed.Channel.Title)

	for _, item := range feed.Channel.Item {
		fmt.Printf("* %v\n", item.Title)
	}

	return nil
}
