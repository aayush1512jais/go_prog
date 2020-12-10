package main

import (
	"database/sql"

	"github.com/aayush1512jais/go_prog/MockingTestify/domain"
	"github.com/aayush1512jais/go_prog/MockingTestify/repository"
)

func main() {
	var db *sql.DB
	appRepo := repository.NewDatabase(db)
	domain.NewDomain(appRepo)
}
