package main

import (
	"context"
	"financial/config/database"
	"financial/controller"
	"financial/repository"
	"financial/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	// Load environment variables
	loadEnv()

	// Initialize application context and dependencies
	ctx := context.Background()
	financialController := initializeDependencies(ctx)

	// Initialize and start the router
	r := setupRouter(financialController)
	log.Println("Starting server on port 8080...")
	r.Run(":8080")
}

// loadEnv loads environment variables from .env file
func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

// initializeDependencies sets up the database, repository, service, and controller
func initializeDependencies(ctx context.Context) *controller.FinancialController {
	// Initialize database connection
	connString := os.Getenv("DATABASE_CONNECTION")
	_, err := database.NewPG(ctx, connString)
	if err != nil {
		log.Fatalf("Error initializing PostgreSQL: %v", err)
	}

	// Initialize repository
	repo, err := repository.NewFinancialRepository(ctx, connString)
	if err != nil {
		log.Fatalf("Error initializing repository: %v", err)
	}

	// Initialize service
	financialService := service.NewFinancialService(repo)

	// Initialize controller
	return controller.NewFinancialController(financialService)
}

// setupRouter sets up the Gin routes
func setupRouter(financialController *controller.FinancialController) *gin.Engine {
	r := gin.Default()

	// Define routes
	r.GET("/financial/:id", func(c *gin.Context) {
		// Retrieve the 'id' parameter from the URL
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		// Call the controller method to get financial details
		financial, err := financialController.GetFinancialById(c, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Return the financial details as JSON
		c.JSON(http.StatusOK, financial)
	})

	return r
}
