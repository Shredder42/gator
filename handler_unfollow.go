package main

import (
	"context"
	"fmt"

	"github.com/Shredder42/gator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("command unfollow requires url only")
	}

	url := cmd.arguments[0]

	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("error getting feed: %w", err)
	}

	err = s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("error deleting feed, feed not deleted: %w", err)
	}

	fmt.Printf("Feed: %s at %s unfollowed", feed.Name, feed.Url)
	return nil
}
