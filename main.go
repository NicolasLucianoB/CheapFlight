package main

import (
	"CheapFlight/database"
	_ "CheapFlight/docs"
	"CheapFlight/models"
	"CheapFlight/routes"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Iniciar conexão com banco de dados
	database.Connect()
	migrateDatabase()

	// Criar o roteador Gin
	router := gin.Default()

	// Definir as rotas da API
	routes.SetupRoutes(router)

	// Rodar servidor na porta 8080
	fmt.Println("Servidor rodando na porta 8080...")
	err := router.Run(":8080")
	if err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}

func migrateDatabase() {
	fmt.Println("Rodando migrações...")

	err := database.DB.AutoMigrate(&models.Flight{})
	if err != nil {
		log.Fatalf("Erro ao rodar migrações: %v", err)
	}

	fmt.Println("Migrações concluídas com sucesso!")
}
