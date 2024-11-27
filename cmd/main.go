package main

import (
    "go-receipt-processor/internal/adapters/http"
    "go-receipt-processor/internal/application"
    "github.com/gin-gonic/gin"
)

func main() {
    // Create an instance of the ReceiptService implementation
    receiptService := application.NewReceiptService()  // This returns a ReceiptServiceImpl

    // Create the handler with the receipt service
    handler := http.NewHandler(receiptService)

    // Set up Gin router
    r := gin.Default()

    // Register routes and bind handlers
    r.POST("/process", handler.ProcessReceipt)

    // Run the Gin server
    r.Run(":8080")
}
