package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Shredder42/gator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		log.Fatalf("register handler expects username")
	}

	registerName := cmd.arguments[0]

	userName, err := s.db.GetUser(context.Background(), registerName)
	if err == nil {
		if userName != "" {
			log.Fatalf("username %s already exists", registerName)
		}
	}

	userInfo, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      registerName,
	})
	if err != nil {
		return err
	}

	err = s.config.SetUser(registerName)
	if err != nil {
		return err
	}

	fmt.Printf("user '%s' was created\n", registerName)
	fmt.Printf("User info: %+v\n", userInfo)

	return nil

}
