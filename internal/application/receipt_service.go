package application

// Importing necessary packages
// domain: Contains the domain models, like `Receipt` and `Item`.
// http: The `http` package defines the interface for the receipt service.
// math: Provides mathematical functions, like `Floor` and `Ceil`.
// strings: Contains functions for string manipulation, such as `ReplaceAll` and `TrimSpace`.
// uuid: Used to generate unique identifiers (UUIDs).
import (
	"fmt"
	"go-receipt-processor/internal/domain"
	"go-receipt-processor/internal/ports/http"
	"go-receipt-processor/internal/ports/repository"
)

// ReceiptServiceImpl is the concrete implementation of the ReceiptService interface.
// It contains the logic for processing a receipt and calculating points based on various rules.
type ReceiptServiceImpl struct {
	PointsCalculator http.PointsCalculator // Add PointsCalculator as a field
	ReceiptStore     repository.ReceiptStore
}

// NewReceiptService creates a new instance of ReceiptServiceImpl.
// It is a factory function that returns a pointer to a new ReceiptServiceImpl object.
// It accepts c as a parameter, which it explicitly sets as PointsCalculator on the implementation
func NewReceiptService(c http.PointsCalculator, rs repository.ReceiptStore) http.ReceiptService {
	return &ReceiptServiceImpl{PointsCalculator: c, ReceiptStore: rs}
}

// ProcessReceipt is the method that processes a receipt and calculates points based on certain rules.
// It receives a `domain.Receipt` as input, which contains details like the retailer, total, items, and purchase info.
// It returns a unique receipt ID, the calculated points, and any errors that occur during processing.
func (s *ReceiptServiceImpl) ProcessReceipt(receipt domain.Receipt) (string, int, error) {

	points, err := s.PointsCalculator.CalculatePoints(receipt)
	if err != nil {
		return "", 0, fmt.Errorf("invalid purchase time format: %v", err)
	}

	// @TODO: create a repository
	// save receipt to repository with an id
	// find receipt by id
	receipt.Points = points
	receiptID, err := s.ReceiptStore.Save(receipt)
	if err != nil {
		return "", 0, fmt.Errorf("failed to insert receipt: %v", err)
	}

	// Return the generated receipt ID, the calculated points, and nil (since there is no error in this case).
	return receiptID, points, nil // Make sure the return matches the method signature from the ReceiptService interface.
}
