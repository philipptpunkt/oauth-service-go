#!/bin/bash

# Function to print errors in red
print_error() {
  echo -e "\033[31mERROR: $1\033[0m"
}

if [ -z "$1" ]; then
  print_error "Database URL is missing. Usage: ./run-migration.sh <DATABASE_URL> <ACTION>"
  echo "Example: ./migrate.sh postgres://user:password@localhost:5432/database?sslmode=disable up"
  exit 1
fi

if [ -z "$2" ]; then
  print_error "Migration action is missing. Usage: ./run-migration.sh <DATABASE_URL> <ACTION>"
  echo "Example: ./migrate.sh postgres://user:password@localhost:5432/database?sslmode=disable up"
  exit 1
fi

go run ./main.go --action $2 --path ./migrations --db $1 --steps $3
