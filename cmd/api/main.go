package main

import (
	"go-receipt-processor/internal/container"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create the container that manages all dependencies
	c := container.NewContainer()

	// Set up Gin router
	g := gin.Default()

	// Register routes and pass the handlers from the container
	g.POST("/receipt/process", c.NewReceiptProcessHandler().ProcessReceipt)

	// Run the Gin server
	g.Run(":8080")

}
