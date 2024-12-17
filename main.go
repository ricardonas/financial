package main

import (
	"context"
	"financial/config/database"
	"financial/controller"
	"financial/repository"
	"financial/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func main() {
	// Create a context and PostgreSQL connection string
	ctx := context.Background()
	connString := "postgres://admin:admin@localhost:5432/financial" // Correct connection string format

	// Initialize PostgreSQL database connection
	_, err := database.NewPG(ctx, connString)
	if err != nil {
		log.Fatalf("Error initializing PostgreSQL: %v", err)
	}

	// Create the repository
	repo, err := repository.NewFinancialRepository(ctx, connString)
	if err != nil {
		log.Fatalf("Error initializing repository: %v", err)
	}

	// Create the service
	financialService := service.NewFinancialService(repo)

	// Create the controller
	financialController := controller.NewFinancialController(financialService)

	// Initialize Gin router
	r := gin.Default()

	// Define the route
	r.GET("/financial/:id", func(c *gin.Context) {
		// Retrieve the 'id' parameter from the URL
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr) // Convert string to integer
		if err != nil {
			// If the id is not a valid integer, return an error
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		// Call the controller method to get financial details
		financial, err := financialController.GetFinancialById(c, id) // Pass c to GetFinancialById
		if err != nil {
			// Handle errors (e.g., if the record is not found)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Return the financial details as JSON
		c.JSON(http.StatusOK, financial)
	})

	// Start the server
	r.Run(":8080")
}
