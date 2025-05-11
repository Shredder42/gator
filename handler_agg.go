package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Shredder42/gator/internal/database"
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

	fmt.Println(rssFeed.Channel.Title)
	for _, item := range rssFeed.Channel.Item {
		fmt.Println(item.Title)
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
