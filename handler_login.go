package main

import (
	"fmt"
	"log"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		log.Fatalf("login handler expects username")
	}

	err := s.config.SetUser(cmd.arguments[0])
	if err != nil {
		return err
	}

	fmt.Printf("User %s has been set\n", cmd.arguments[0])

	return nil
}
