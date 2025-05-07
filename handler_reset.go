package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	if len(cmd.arguments) != 0 {
		return fmt.Errorf("command reset does not require username")
	}

	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("database reset successful")

	return nil
}
