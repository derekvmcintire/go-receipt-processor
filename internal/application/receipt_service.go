package application

// Package application contains the core business logic of the receipt processing service.
// It provides the ReceiptService, which handles receipt processing, points calculation, and interaction with repositories.

import (
	"fmt"
	"go-receipt-processor/internal/domain"
	"go-receipt-processor/internal/ports/http"
	"go-receipt-processor/internal/ports/repository"
)

// ReceiptServiceImpl is an implementation of the ReceiptService interface.
// It processes receipts, calculates points, and stores the receipts in a repository.
type ReceiptServiceImpl struct {
	PointsCalculator http.PointsCalculator   // Calculates points based on receipt data.
	ReceiptStore     repository.ReceiptStore // Manages storage and retrieval of receipts.
}

// NewReceiptService creates and returns a new instance of ReceiptServiceImpl.
// It takes a PointsCalculator for calculating points and a ReceiptStore for managing receipts.
func NewReceiptService(c http.PointsCalculator, rs repository.ReceiptStore) http.ReceiptService {
	return &ReceiptServiceImpl{
		PointsCalculator: c,
		ReceiptStore:     rs,
	}
}

// ProcessReceipt processes a receipt by calculating points and storing it in the repository.
// It returns the receipt's unique ID, the calculated points, and an error, if any.
//
// Parameters:
//   - receipt: The domain.Receipt object containing receipt details.
//
// Returns:
//   - receiptID: A unique identifier for the processed receipt.
//   - points: The calculated points for the receipt.
//   - err: An error if processing or saving fails.
func (s *ReceiptServiceImpl) ProcessReceipt(receipt domain.Receipt) (string, error) {

	// Calculate points using the PointsCalculator.
	points, err := s.PointsCalculator.CalculatePoints(receipt)
	if err != nil {
		return "", fmt.Errorf("invalid purchase time format: %v", err)
	}

	// Assign calculated points to the receipt.
	receipt.Points = points

	// Save the receipt to the repository and retrieve its unique ID.
	receiptID, err := s.ReceiptStore.Save(receipt)
	if err != nil {
		return "", fmt.Errorf("failed to insert receipt: %v", err)
	}

	return receiptID, nil
}
