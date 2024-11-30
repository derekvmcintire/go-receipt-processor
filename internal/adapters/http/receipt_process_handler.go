package http

import (
	"go-receipt-processor/internal/domain"
	internalHttp "go-receipt-processor/internal/ports/core"
	"go-receipt-processor/internal/ports/http/response"
	netHttp "net/http"

	"github.com/gin-gonic/gin"
)

// ReceiptProcessHandler manages HTTP requests for processing receipts.
type ReceiptProcessHandler struct {
	ReceiptService internalHttp.ReceiptService
}

// NewReceiptProcessHandler
//
// Parameters:
//   - service: The ReceiptService responsible for processing the receipt and calculating points.
//
// Returns:
//   - A new instance of ReceiptProcessHandler with the provided ReceiptService.
func NewReceiptProcessHandler(service internalHttp.ReceiptService) *ReceiptProcessHandler {
	return &ReceiptProcessHandler{ReceiptService: service}
}

// ProcessReceipt
//
// Parameters:
//   - c: The Gin context, which contains the HTTP request and response data.
//
// Returns:
//   - A JSON response with either a 200 OK status and the receipt ID, or a 400 Bad Request if input validation fails,
//     or a 500 Internal Server Error if processing the receipt fails.
func (h *ReceiptProcessHandler) ProcessReceipt(c *gin.Context) {
	var receipt domain.Receipt

	if err := c.ShouldBindJSON(&receipt); err != nil {
		c.JSON(netHttp.StatusBadRequest, gin.H{"error": "Invalid request payload", "details": err.Error()})
		return
	}

	receiptID, err := h.ReceiptService.ProcessReceipt(receipt)
	if err != nil {
		c.JSON(netHttp.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(netHttp.StatusOK, response.ReceiptProcessResponse{
		ID: receiptID,
	})
}
