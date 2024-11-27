package application

import (
    "go-receipt-processor/internal/domain"
    "go-receipt-processor/internal/ports/http"
    "math"
    "strings"
    "github.com/google/uuid"
)

// ReceiptServiceImpl is the concrete implementation of the ReceiptService interface.
type ReceiptServiceImpl struct{}

// NewReceiptService creates a new instance of ReceiptServiceImpl.
func NewReceiptService() http.ReceiptService {
    return &ReceiptServiceImpl{}
}

// ProcessReceipt processes a receipt and calculates points based on the receipt details.
func (s *ReceiptServiceImpl) ProcessReceipt(receipt domain.Receipt) (string, int, error) {
    points := 0

    // Rule 1: One point for each alphanumeric character in the retailer name.
    points += len(strings.ReplaceAll(receipt.Retailer, " ", ""))

    // Rule 2: 50 points if the total is a round dollar amount with no cents.
    if receipt.Total == math.Floor(receipt.Total) {
        points += 50
    }

    // Rule 3: 25 points if the total is a multiple of 0.25.
    if int(receipt.Total*100)%25 == 0 {
        points += 25
    }

    // Rule 4: 5 points for every two items on the receipt.
    points += (len(receipt.Items) / 2) * 5

    // Rule 5: For items with descriptions that are a multiple of 3 in length, apply the price multiplier rule.
    for _, item := range receipt.Items {
        if len(strings.TrimSpace(item.ShortDescription))%3 == 0 {
            points += int(math.Ceil(item.Price * 0.2))
        }
    }

    // Rule 6: 6 points if the purchase day is odd.
    if receipt.PurchaseDate.Day()%2 != 0 {
        points += 6
    }

    // Rule 7: 10 points if the purchase time is between 2:00pm and 4:00pm.
    if receipt.PurchaseTime.Hour() >= 14 && receipt.PurchaseTime.Hour() < 16 {
        points += 10
    }

    // Generate a unique ID for the receipt (for example purposes, this could be a UUID).
    receiptID := uuid.New().String()

    return receiptID, points, nil // Make sure this matches the interface's signature.
}
