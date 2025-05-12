package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Shredder42/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) > 1 {
		return fmt.Errorf("command agg requires nothing or limit only")
	}

	postLimit := 2
	if len(cmd.arguments) == 1 {
		if specifiedLimit, err := strconv.Atoi(cmd.arguments[0]); err == nil {
			postLimit = specifiedLimit
		}
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(postLimit),
	})
	if err != nil {
		return fmt.Errorf("error getting posts: %w", err)
	}

	for _, post := range posts {
		fmt.Println(post.Title)
		fmt.Println(post.Url)
		fmt.Println(post.Description)
		fmt.Println(post.PublishedAt)
		fmt.Println("==================================")
	}
	return nil

}
