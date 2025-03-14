package routes

import (
	"CheapFlight/controllers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files" // Usando o alias swaggerFiles
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title CheapFlight API
// @version 1.0
// @description API para busca e gerenciamento de voos.

// SetupRoutes define as rotas da API
func SetupRoutes(router *gin.Engine) {
	// Criando um grupo de rotas para organização
	api := router.Group("/api")
	{
		// Rota para listar todos os voos
		api.GET("/flights", controllers.GetFlights)

		// Rota para criar um voo
		api.POST("/flights", controllers.CreateFlight)

		// Rota para obter um voo específico
		api.GET("/flights/:id", controllers.GetFlight)

		// Rota para atualizar um voo
		api.PUT("/flights/:id", controllers.UpdateFlight)

		// Rota para deletar um voo
		api.DELETE("/flights/:id", controllers.DeleteFlight)
	}

	// Configuração do Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) // Usando o alias corretamente
}
