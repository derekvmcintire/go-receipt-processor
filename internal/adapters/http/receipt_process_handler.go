// Package http provides the HTTP handlers for receipt processing.
// It integrates with the business logic layer and manages the API responses.
package http

import (
	"go-receipt-processor/internal/domain"                  // Package for domain models like Receipt
	internalHttp "go-receipt-processor/internal/ports/http" // Service interface for receipt processing (aliased to avoid conflict with Gin)
	"go-receipt-processor/internal/ports/http/response"     // Response interface for defining the shape of the return data
	netHttp "net/http"                                      // Standard Go package for HTTP-related constants (aliased to avoid naming conflicts)

	"github.com/gin-gonic/gin" // Web framework for building REST APIs in Go
)

// ReceiptProcessHandler manages HTTP requests for processing receipts.
// It interacts with the ReceiptService to handle the core business logic.
type ReceiptProcessHandler struct {
	ReceiptService internalHttp.ReceiptService // Service interface for processing receipts and calculating points
}

// NewReceiptProcessHandler is a constructor function for creating a ReceiptProcessHandler instance.
// It takes a ReceiptService as a dependency, enabling proper dependency injection.
func NewReceiptProcessHandler(service internalHttp.ReceiptService) *ReceiptProcessHandler {
	return &ReceiptProcessHandler{ReceiptService: service}
}

// ProcessReceipt handles HTTP POST requests to the `/receipt/process` route.
// It validates the input, processes the receipt, and returns the result.
func (h *ReceiptProcessHandler) ProcessReceipt(c *gin.Context) {
	var receipt domain.Receipt // Variable to hold the parsed receipt from the request body

	// Attempt to parse the JSON body into the `receipt` struct
	if err := c.ShouldBindJSON(&receipt); err != nil {
		// Respond with a 400 Bad Request if JSON binding fails
		c.JSON(netHttp.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the service to process the receipt and calculate points
	receiptID, err := h.ReceiptService.ProcessReceipt(receipt)
	if err != nil {
		// Respond with a 500 Internal Server Error if processing fails
		c.JSON(netHttp.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with a 200 OK and return the receipt ID in a structured response
	c.JSON(netHttp.StatusOK, response.ReceiptProcessResponse{
		ID: receiptID, // ID of the processed receipt
	})
}
