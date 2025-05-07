package main

import (
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
