package http

import (
	internalHttp "go-receipt-processor/internal/ports/http"
	"go-receipt-processor/internal/ports/http/response"
	netHttp "net/http"

	"github.com/gin-gonic/gin"
)

// GetReceiptPointsHandler manages HTTP requests for getting receipt points.
// It interacts with the ReceiptService to handle the core business logic.
type GetReceiptPointsHandler struct {
	ReceiptService internalHttp.ReceiptService // Service interface for getting receipt points
}

// NewGetReceioptPointsHandler is a constructor function for creating a GetReceiptPointsHandler instance.
// It takes a ReceiptService as a dependency, enabling proper dependency injection.
func NewGetReceiptPointsHandler(service internalHttp.ReceiptService) *GetReceiptPointsHandler {
	return &GetReceiptPointsHandler{ReceiptService: service}
}

func (h *GetReceiptPointsHandler) GetPoints(c *gin.Context) {
	id := c.Param("id")

	// Attempt to parse the
	points, err := h.ReceiptService.GetPoints(id)
	if err != nil {
		// Respond with a 500 Internal Server Error if processing fails
		c.JSON(netHttp.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with a 200 OK and return the receipt ID in a structured response
	c.JSON(netHttp.StatusOK, response.GetReceiptPointsResponse{
		Points: points, // Points from the saved receipt
	})
}
