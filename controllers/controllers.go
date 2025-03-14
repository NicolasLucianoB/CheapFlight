package controllers

import (
	"CheapFlight/database"
	"CheapFlight/models"

	"github.com/gin-gonic/gin"
)

// GetFlights retorna a lista de todos os voos
// @Summary Lista todos os voos
// @Description Retorna todos os voos disponíveis
// @Tags flights
// @Produce json
// @Success 200 {array} models.Flight
// @Router /api/flights [get]
func GetFlights(c *gin.Context) {
	var flights []models.Flight
	database.DB.Find(&flights)
	c.JSON(200, flights)
}

// GetFlight retorna um voo pelo ID
// @Summary Obtém um voo específico
// @Description Retorna os detalhes de um voo pelo ID
// @Tags flights
// @Produce json
// @Param id path int true "ID do voo"
// @Success 200 {object} models.Flight
// @Failure 404 {object} map[string]string
// @Router /api/flights/{id} [get]
func GetFlight(c *gin.Context) {
	id := c.Param("id")
	var flight models.Flight
	result := database.DB.First(&flight, id)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Voo não encontrado"})
		return
	}
	c.JSON(200, flight)
}

// CreateFlight cria um novo voo
// @Summary Cria um novo voo
// @Description Adiciona um novo voo à base de dados
// @Tags flights
// @Accept json
// @Produce json
// @Param flight body models.Flight true "Dados do voo"
// @Success 201 {object} models.Flight
// @Failure 400 {object} map[string]string
// @Router /api/flights [post]
func CreateFlight(c *gin.Context) {
	var flight models.Flight
	if err := c.ShouldBindJSON(&flight); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&flight)
	c.JSON(201, flight)
}

// UpdateFlight atualiza um voo existente
// @Summary Atualiza um voo
// @Description Modifica os dados de um voo pelo ID
// @Tags flights
// @Accept json
// @Produce json
// @Param id path int true "ID do voo"
// @Param flight body models.Flight true "Novos dados do voo"
// @Success 200 {object} models.Flight
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/flights/{id} [put]
func UpdateFlight(c *gin.Context) {
	id := c.Param("id")
	var flight models.Flight
	database.DB.First(&flight, id)
	if flight.ID == 0 {
		c.JSON(404, gin.H{"error": "Voo não encontrado"})
		return
	}

	if err := c.ShouldBindJSON(&flight); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	database.DB.Save(&flight)
	c.JSON(200, flight)
}

// DeleteFlight deleta um voo
// @Summary Deleta um voo
// @Description Remove um voo da base de dados pelo ID
// @Tags flights
// @Param id path int true "ID do voo"
// @Success 200 {object} map[string]string
// @Router /api/flights/{id} [delete]
func DeleteFlight(c *gin.Context) {
	id := c.Param("id")
	database.DB.Delete(&models.Flight{}, id)
	c.JSON(200, gin.H{"message": "Voo deletado com sucesso"})
}
