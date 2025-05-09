package main

import (
	"context"
	"fmt"
)

func handlerFollowing(s *state, cmd command) error {
	if len(cmd.arguments) != 0 {
		return fmt.Errorf("command following does not require aditional arguments")
	}

	currentUser := s.config.CurrentUserName
	user, err := s.db.GetUser(context.Background(), currentUser)
	if err != nil {
		return fmt.Errorf("error getting user: %w", err)
	}

	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("error getting feed follows: %w", err)
	}

	fmt.Printf("User %s is following:\n", currentUser)
	for _, follow := range feedFollows {
		fmt.Println("* ", follow.FeedName)
	}

	return nil

}
