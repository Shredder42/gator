package main

import (
	"context"
	"fmt"
	"log"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		log.Fatalf("login handler expects username")
	}

	loginName := cmd.arguments[0]

	_, err := s.db.GetUser(context.Background(), loginName)
	if err != nil {
		log.Fatalf("can't login in to nonexistent account %s", loginName)
	}

	err = s.config.SetUser(loginName)
	if err != nil {
		return err
	}

	fmt.Printf("User %s has been logged in\n", loginName)

	return nil
}
