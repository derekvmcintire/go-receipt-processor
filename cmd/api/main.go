// Package main serves as the entry point of the application.
// It sets up the necessary infrastructure, such as dependency injection and HTTP routing.
package main

import (
	"go-receipt-processor/internal/container"

	"github.com/gin-gonic/gin"
)

// main is the entry point of the application.
// It initializes the dependency container, sets up routes, and starts the HTTP server.
func main() {
	// Initialize the dependency container, which manages all application services and handlers.
	c := container.NewContainer()

	// Create a new Gin router instance for handling HTTP requests.
	g := gin.Default()

	// Register the route for processing receipts.
	// The handler is retrieved from the container to ensure proper dependency injection.
	g.POST("/receipt/process", c.NewReceiptProcessHandler().ProcessReceipt)

	// Register the route to get points by receipt id.
	// The handler is retrieved from the container to ensure proper dependency injection.
	g.GET("/receipt/:id/points", c.NewGetReceiptPointsHandler().GetPoints)

	// Start the Gin HTTP server on port 8080.
	// This will block until the application is terminated or the server fails to start.
	g.Run(":8080")
}
