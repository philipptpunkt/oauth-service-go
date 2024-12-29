package main

import (
	"flag"
	"log"
	"strconv"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	// Flags for command-line options
	migrationDir := flag.String("path", "./migrations", "Path to the migrations folder")
	databaseURL := flag.String("db", "postgres://user:password@localhost:5432/oauthservice?sslmode=disable", "Database connection URL")
	action := flag.String("action", "up", "Migration action: up, down, or drop")
	stepsStr := flag.String("steps", "", "Number of steps to migrate (only for 'up' or 'down')")

	flag.Parse()

	// Initialize migration driver
	m, err := migrate.New(
		"file://"+*migrationDir,
		*databaseURL,
	)
	if err != nil {
		log.Fatalf("Failed to create migration instance: %v", err)
	}

	// Perform the requested action
	switch *action {
	case "up":
		if *stepsStr != "" {
			steps, err := strconv.Atoi(*stepsStr)
			if err != nil {
				log.Fatalf("Invalid steps argument: %v", err)
			}
			if err := m.Steps(steps); err != nil && err != migrate.ErrNoChange {
				log.Fatalf("Migration up %d steps failed: %v", steps, err)
			}
			log.Printf("Migration up %d steps completed", steps)
		} else {
			if err := m.Up(); err != nil && err != migrate.ErrNoChange {
				log.Fatalf("Migration up failed: %v", err)
			}
			log.Println("Migration up completed")
		}
	case "down":
		if *stepsStr != "" {
			steps, err := strconv.Atoi(*stepsStr)
			if err != nil {
				log.Fatalf("Invalid steps argument: %v", err)
			}
			if err := m.Steps(-steps); err != nil {
				log.Fatalf("Migration down %d steps failed: %v", steps, err)
			}
			log.Printf("Migration down %d steps completed", steps)
		} else {
			if err := m.Down(); err != nil {
				log.Fatalf("Migration down failed: %v", err)
			}
			log.Println("Migration down completed")
		}
	case "drop":
		if err := m.Drop(); err != nil {
			log.Fatalf("Migration drop failed: %v", err)
		}
		log.Println("Database dropped successfully")
	default:
		log.Fatalf("Unknown action: %s", *action)
	}
}
