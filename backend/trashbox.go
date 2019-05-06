package main

import (
	"log"

	"github.com/Mushus/trashbox/backend/server"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	s, cleanup, err := server.InitializeServer()
	if err != nil {
		log.Fatalf("failed to initialize app: %v", err)
	}
	defer cleanup()

	s.Start()
}
