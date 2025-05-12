package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/Shredder42/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("command agg requires time_between_reqs only")
	}

	time_between_reqs := cmd.arguments[0]

	timeBetweenRequests, err := time.ParseDuration(time_between_reqs)
	if err != nil {
		return fmt.Errorf("error parsing time duration: %w", err)
	}

	fmt.Printf("Collecting feeds every %s\n", timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}

}

func scrapeFeeds(s *state) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("error getting next feed: %w", err)
	}

	err = s.db.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
		LastFetchedAt: getCurrentTimeToSQLNullTime(),
		UpdatedAt:     time.Now().UTC(),
		ID:            feed.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't mark feed fetched: %w", err)
	}

	rssFeed, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		return fmt.Errorf("error fetching feed: %w", err)
	}

	fmt.Printf("Adding posts for %v\n", rssFeed.Channel.Title)
	for _, item := range rssFeed.Channel.Item {

		pubTime := parseTimeString(item.PubDate)

		_, err := s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Title:       item.Title,
			Url:         item.Link,
			Description: item.Description,
			PublishedAt: pubTime,
			FeedID:      feed.ID,
		})
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				continue
			}
			log.Printf("error creating post: %v", err)
			continue
		}
		fmt.Printf("Added: %s\n", item.Title)
	}
	fmt.Println("==============================================================")

	return nil
}

func getCurrentTimeToSQLNullTime() sql.NullTime {
	currentTime := time.Now().UTC()
	return sql.NullTime{
		Time:  currentTime,
		Valid: true,
	}
}

func parseTimeString(str string) time.Time {
	timeString := str[5:]

	contains_plus := strings.Contains(timeString, "+")
	if contains_plus {
		timeValue, err := time.Parse("02 Jan 2006 15:04:05 Z0700", timeString)
		if err != nil {
			log.Fatalf("error parsing time value")
		}
		return timeValue
	}

	contains_gmt := strings.Contains(timeString, "GMT")
	if contains_gmt {
		timeValue, err := time.Parse("02 Jan 2006 15:04:05 MST", timeString)
		if err != nil {
			log.Fatalf("error parsing time value")
		}
		return timeValue
	}

	return time.Time{}
}
