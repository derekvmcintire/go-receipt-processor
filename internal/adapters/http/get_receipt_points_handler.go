package http

import (
	internalHttp "go-receipt-processor/internal/ports/http" // Service interface for retrieving receipt points
	"go-receipt-processor/internal/ports/http/response"     // Response struct for returning receipt points in a standardized format
	netHttp "net/http"                                      // Standard Go package for HTTP-related constants (aliased to avoid naming conflicts)

	"github.com/gin-gonic/gin" // Web framework for building REST APIs in Go
)

// GetReceiptPointsHandler manages HTTP requests for getting receipt points by ID.
// It interacts with the ReceiptService to handle the business logic of fetching points for a specific receipt.
type GetReceiptPointsHandler struct {
	ReceiptService internalHttp.ReceiptService // Service interface for getting receipt points
}

// NewGetReceiptPointsHandler creates and returns a new instance of GetReceiptPointsHandler.
// It takes a ReceiptService as a dependency, enabling proper dependency injection.
//
// Parameters:
//   - service: The ReceiptService responsible for fetching receipt points.
//
// Returns:
//   - A new instance of GetReceiptPointsHandler with the provided ReceiptService.
func NewGetReceiptPointsHandler(service internalHttp.ReceiptService) *GetReceiptPointsHandler {
	return &GetReceiptPointsHandler{ReceiptService: service}
}

// GetPoints handles the HTTP request to retrieve the points for a specific receipt by its ID.
// It retrieves the receipt points from the ReceiptService and sends an appropriate HTTP response.
//
// Parameters:
//   - c: The Gin context, which contains the HTTP request and response data.
//
// Returns:
//   - A JSON response with either a 200 OK status and the points, or a 500 Internal Server Error if an error occurs.
func (h *GetReceiptPointsHandler) GetPoints(c *gin.Context) {
	// Retrieve the receipt ID from the URL parameters (e.g., /receipts/:id/points).
	id := c.Param("id")

	// Attempt to fetch the receipt points from the ReceiptService using the given ID.
	points, err := h.ReceiptService.GetPoints(id)
	if err != nil {
		// If there is an error (e.g., receipt not found), respond with a 500 Internal Server Error.
		c.JSON(netHttp.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// If no error occurs, respond with a 200 OK and return the points in a structured response.
	c.JSON(netHttp.StatusOK, response.GetReceiptPointsResponse{
		Points: points, // Points extracted from the saved receipt
	})
}
