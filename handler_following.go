package main

import (
	"context"
	"fmt"

	"github.com/Shredder42/gator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) != 0 {
		return fmt.Errorf("command following does not require aditional arguments")
	}

	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("error getting feed follows: %w", err)
	}

	if len(feedFollows) == 0 {
		fmt.Printf("No feed follows found for user %s\n", user.Name)
	} else {
		fmt.Printf("User %s is following:\n", user.Name)
		for _, follow := range feedFollows {
			fmt.Println("* ", follow.FeedName)
		}
	}

	return nil

}
