package http

import (
    "github.com/gin-gonic/gin"
    "go-receipt-processor/internal/domain"
    "go-receipt-processor/internal/application"
    internalHttp "go-receipt-processor/internal/ports/http" // Alias this import to avoid conflict
    netHttp "net/http" // Alias net/http to netHttp
    "github.com/google/uuid"
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

    // Calculate points based on the receipt details.
    points := application.CalculatePoints(&receipt)

    // Return the generated ID and points
    c.JSON(netHttp.StatusOK, gin.H{
        "id":     receipt.ID,
        "points": points,
    })
}
