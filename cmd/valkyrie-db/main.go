package main

import (
	"fmt"
	"log"

	"github.com/valkyriedb/valkyrie/config"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Port: %s\n", cfg.Port)
}
