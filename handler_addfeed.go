package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Shredder42/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.arguments) != 2 {
		log.Fatalf("addfeed command requires 2 arguments: feed name and url")
	}

	feedName := cmd.arguments[0]
	feedUrl := cmd.arguments[1]
	currentUserName := s.config.CurrentUserName

	user, err := s.db.GetUser(context.Background(), currentUserName)
	if err != nil {
		log.Fatalf("user not found %s", currentUserName)
	}

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      feedName,
		Url:       feedUrl,
		UserID:    user.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't add feed: %w", err)
	}

	fmt.Println("Feed created successfully:")
	printFeed(feed, user)
	return nil
}

func printFeed(feed database.Feed, user database.User) {
	fmt.Printf("* ID:            %s\n", feed.ID)
	fmt.Printf("* Created:       %v\n", feed.CreatedAt)
	fmt.Printf("* Updated:       %v\n", feed.UpdatedAt)
	fmt.Printf("* Name:          %s\n", feed.Name)
	fmt.Printf("* URL:           %s\n", feed.Url)
	fmt.Printf("* UserID:        %s\n", feed.UserID)
	fmt.Printf("* User:			 %s\n", user.Name)
}
