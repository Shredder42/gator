package main

import (
	"context"
	"fmt"
	"log"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.arguments) != 0 {
		return fmt.Errorf("command agg does not require username")
	}

	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		log.Fatalf("error fetching feed")
	}

	fmt.Printf("%v\n", feed)
	return nil

}
