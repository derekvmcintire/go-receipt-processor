package http

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go-receipt-processor/internal/domain"
	internalHttp "go-receipt-processor/internal/ports/http" // Alias this import to avoid conflict
	netHttp "net/http"                                      // Alias net/http to netHttp
)

// Handler is a struct that holds the service for processing receipts.
type Handler struct {
	ReceiptService internalHttp.ReceiptService // Use the interface from internal/ports/http
}

// NewHandler creates a new Handler.
func NewHandler(service internalHttp.ReceiptService) *Handler {
	return &Handler{ReceiptService: service}
}

// ProcessReceipt handles the receipt processing route.
func (h *Handler) ProcessReceipt(c *gin.Context) {
	var receipt domain.Receipt
	if err := c.ShouldBindJSON(&receipt); err != nil {
		// Use the alias for net/http
		c.JSON(netHttp.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate a unique ID for the receipt.
	receipt.ID = uuid.New().String()

	// Call the ProcessReceipt method from the ReceiptService to calculate points.
	receiptID, points, err := h.ReceiptService.ProcessReceipt(receipt)
	if err != nil {
		// Handle error (optional, depending on your design)
		c.JSON(netHttp.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the generated ID and points
	c.JSON(netHttp.StatusOK, gin.H{
		"id":     receiptID,
		"points": points,
	})
}
