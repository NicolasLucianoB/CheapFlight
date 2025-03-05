package main

import (
	"CheapFlight/models/database"
	"CheapFlight/models/migrations"
)

func main() {
	database.Connect()
	migrations.RunMigrations()
}
