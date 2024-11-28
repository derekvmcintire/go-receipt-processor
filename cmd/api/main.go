package main

import (
	"github.com/gin-gonic/gin"
	"go-receipt-processor/internal/container"
)

func main() {
	// Create the container that manages all dependencies
	appContainer := container.NewContainer()

	// Set up Gin router
	r := gin.Default()

	// Register routes and pass the handlers from the container
	r.POST("/receipt/process", appContainer.NewReceiptProcessHandler().ProcessReceipt)

	// Run the Gin server
	r.Run(":8080")
}
