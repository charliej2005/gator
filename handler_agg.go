package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/charliej2005/gator/internal/database"
	"github.com/google/uuid"
	"github.com/lib/pq"
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
	if err == sql.ErrNoRows {
		fmt.Println("No feeds to fetch.")
		return nil
	}
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
		fmt.Printf("itemname: %s\n", item.Title)
		nullTime := sql.NullTime{Time: time.Now(), Valid: true}

		pubTime := parsePubDate(item.PubDate)

		params := database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   nullTime,
			UpdatedAt:   nullTime,
			Title:       sql.NullString{String: item.Title, Valid: item.Title != ""},
			Url:         item.Link,
			Description: sql.NullString{String: item.Description, Valid: item.Description != ""},
			PublishedAt: sql.NullTime{Time: pubTime, Valid: !pubTime.IsZero()},
			FeedID:      dbFeed.ID,
		}
		_, err = s.db.CreatePost(context.Background(), params)
		if err != nil {
			if pqErr, ok := err.(*pq.Error); !ok || pqErr.Code != "23505" {
				return err
			}
		}
	}

	return nil
}

func parsePubDate(pubDate string) time.Time {
	layouts := []string{
		time.RFC1123Z,
		time.RFC1123,
		time.RFC822Z,
		time.RFC822,
	}
	for _, layout := range layouts {
		if t, err := time.Parse(layout, pubDate); err == nil {
			return t
		}
	}
	return time.Time{}
}
