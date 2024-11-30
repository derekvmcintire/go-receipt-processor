package http

import (
	internalHttp "go-receipt-processor/internal/ports/core"
	"go-receipt-processor/internal/ports/http/response"
	netHttp "net/http"

	"github.com/gin-gonic/gin"
)

// GetReceiptPointsHandler manages HTTP requests for getting receipt points by ID.
type GetReceiptPointsHandler struct {
	ReceiptService internalHttp.ReceiptService // Service interface for getting receipt points
}

// NewGetReceiptPointsHandler
//
// Parameters:
//   - service: The ReceiptService responsible for fetching receipt points.
//
// Returns:
//   - A new instance of GetReceiptPointsHandler with the provided ReceiptService.
func NewGetReceiptPointsHandler(service internalHttp.ReceiptService) *GetReceiptPointsHandler {
	return &GetReceiptPointsHandler{ReceiptService: service}
}

// GetPoints
//
// Parameters:
//   - c: The Gin context, which contains the HTTP request and response data.
//
// Returns:
//   - A JSON response with either a 200 OK status and the points, or a 500 Internal Server Error if an error occurs.
func (h *GetReceiptPointsHandler) GetPoints(c *gin.Context) {
	id := c.Param("id")

	points, err := h.ReceiptService.GetPoints(id)
	if err != nil {
		c.JSON(netHttp.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(netHttp.StatusOK, response.GetReceiptPointsResponse{
		Points: points,
	})
}
