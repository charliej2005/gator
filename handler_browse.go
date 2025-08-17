package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/charliej2005/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	var limit int32
	if len(cmd.args) == 0 {
		limit = 2
	} else {
		parsedLimit, err := strconv.Atoi(cmd.args[0])
		if err != nil {
			return err
		}
		limit = int32(parsedLimit)
	}

	params := database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  limit,
	}

	posts, err := s.db.GetPostsForUser(context.Background(), params)
	if err != nil {
		return err
	}

	for i, post := range posts {
		if i == 0 {
			fmt.Println()
		}
		var description string
		if !post.Description.Valid {
			description = "[empty description]"
		} else {
			description = post.Description.String
		}
		var publishedAt string
		if post.PublishedAt.Valid {
			publishedAt = post.PublishedAt.Time.Format("2006-01-02 15:04:05")
		} else {
			publishedAt = "[unknown publish time]"
		}
		fmt.Printf("%v:\n[%v]\n%v\n", post.Title.String, publishedAt, description)
		fmt.Println()
	}
	return nil
}
