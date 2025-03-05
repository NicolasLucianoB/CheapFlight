package migrations

import (
	"CheapFlight/models/database"
	"CheapFlight/models/models"
	"fmt"
	"log"
)

func RunMigrations() {
	fmt.Println("Rodando migrações...")

	err := database.DB.AutoMigrate(&models.Flight{})
	if err != nil {
		log.Fatalf("Erro ao rodar migrações: %v", err)
	}

	fmt.Println("Migrações concluídas!")
}
