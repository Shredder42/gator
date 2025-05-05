package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Shredder42/gator/internal/config"
)

func main() {

	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	programState := &state{
		config: &cfg,
	}

	allCommands := commands{
		cmdMap: make(map[string]func(*state, command) error),
	}

	allCommands.register("login", handlerLogin)

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
	config *config.Config
}
