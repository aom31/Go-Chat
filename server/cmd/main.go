package main

import (
	"log"
	"server/db"
	"server/internal/config"
)

func main() {
	container, err := config.NewConfig()
	if err != nil {
		log.Fatalf("error loading environment variables: %v", err)

	}

	if _, err := db.NewDatabase(container.DBPostgres); err != nil {
		log.Fatalf("could not initialize database connection: %s", err)
	}
}
