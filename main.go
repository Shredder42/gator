package main

import (
	"fmt"
	"log"

	"github.com/Shredder42/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	err = cfg.SetUser("shredder")

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	fmt.Printf("Output: %+v\n", cfg)
}
