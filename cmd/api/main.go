package main

import (
	"go-receipt-processor/cmd/container"

	"github.com/gin-gonic/gin"
)

// main is the entry point of the application.
func main() {
	// Initialize the dependency container, which manages all application services and handlers.
	c := container.NewContainer()

	// Create a new Gin router instance for handling HTTP requests.
	g := gin.Default()

	// Register the routes
	g.POST("/receipt/process", c.NewReceiptProcessHandler().ProcessReceipt)
	g.GET("/receipt/:id/points", c.NewGetReceiptPointsHandler().GetPoints)

	// Start the Gin HTTP server on port 8080.
	g.Run(":8080")
}
