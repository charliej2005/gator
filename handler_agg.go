package main

import (
	"context"
	"fmt"
)

func handlerAgg(s *state, cmd command) error {
	// currently only fetches from placeholder address
	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}
	fmt.Println(feed)
	return nil
}
