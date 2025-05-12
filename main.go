package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/Shredder42/gator/internal/config"
	"github.com/Shredder42/gator/internal/database"

	_ "github.com/lib/pq"
)

func main() {

	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DbUrl)
	if err != nil {
		log.Fatalf("error connecting to db: %v", err)
	}
	defer db.Close()
	dbQueries := database.New(db)

	programState := &state{
		db:     dbQueries,
		config: &cfg,
	}

	allCommands := commands{
		cmdMap: make(map[string]func(*state, command) error),
	}

	allCommands.register("login", handlerLogin)
	allCommands.register("register", handlerRegister)
	allCommands.register("reset", handlerReset)
	allCommands.register("users", handlerUsers)
	allCommands.register("agg", handlerAgg)
	allCommands.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	allCommands.register("feeds", handlerFeeds)
	allCommands.register("follow", middlewareLoggedIn(handlerFollow))
	allCommands.register("following", middlewareLoggedIn(handlerFollowing))
	allCommands.register("unfollow", middlewareLoggedIn(handlerUnfollow))
	allCommands.register("browse", middlewareLoggedIn(handlerBrowse))

	args := os.Args
	if len(args) < 2 {
		log.Fatalf("require command name")
	}

	commandName := args[1]
	commandArguments := args[2:]

	err = allCommands.run(programState, command{
		name:      commandName,
		arguments: commandArguments,
	})
	if err != nil {
		fmt.Println(err)
	}

}

type state struct {
	db     *database.Queries
	config *config.Config
}

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		currentUserName := s.config.CurrentUserName
		user, err := s.db.GetUser(context.Background(), currentUserName)
		if err != nil {
			log.Fatalf("user not found %s", currentUserName)
		}

		handler(s, cmd, user)

		return nil

	}

}
