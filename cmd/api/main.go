package main

import (
	"github.com/gin-gonic/gin"
	"go-receipt-processor/internal/container"
)

func main() {
	// Create the container that manages all dependencies
	c := container.NewContainer()

	// Set up Gin router
	g := gin.Default()

	// Register routes and pass the handlers from the container
	g.POST("/receipt/process", c.NewReceiptProcessHandler().ProcesReceipt)

	// Run the Gin server
	g.Run(":8080")

}
