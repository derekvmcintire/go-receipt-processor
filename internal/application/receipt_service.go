package application

import (
	"fmt"
	"go-receipt-processor/internal/domain"
	http "go-receipt-processor/internal/ports/core"
	"go-receipt-processor/internal/ports/repository"
)

// ReceiptServiceImpl is an implementation of the ReceiptService interface.
type ReceiptServiceImpl struct {
	PointsCalculator http.PointsCalculator
	ReceiptStore     repository.ReceiptStore
}

// NewReceiptService
//
// Parameters:
//   - c: The PointsCalculator used to calculate points for a receipt.
//   - rs: The ReceiptStore used to store and retrieve receipts.
//
// Returns:
//   - A new instance of ReceiptServiceImpl with the provided dependencies.
func NewReceiptService(c http.PointsCalculator, rs repository.ReceiptStore) http.ReceiptService {
	return &ReceiptServiceImpl{
		PointsCalculator: c,
		ReceiptStore:     rs,
	}
}

// ProcessReceipt
//
// Parameters:
//   - receipt: The domain.Receipt object containing receipt details.
//
// Returns:
//   - receiptID: A unique identifier for the processed receipt.
//   - err: An error if processing or saving fails.
func (s *ReceiptServiceImpl) ProcessReceipt(receipt domain.Receipt) (string, error) {
	points, err := s.PointsCalculator.CalculatePoints(receipt)
	if err != nil {
		return "", fmt.Errorf("unable to process receipt: %v", err)
	}

	receipt.Points = points

	receiptID, err := s.ReceiptStore.Save(receipt)
	if err != nil {
		return "", fmt.Errorf("failed to insert receipt: %v", err)
	}

	return receiptID, nil
}

// GetPoints
//
// Parameters:
//   - id: The unique ID of the receipt whose points are being retrieved.
//
// Returns:
//   - points: The points associated with the receipt.
//   - err: An error if the receipt cannot be found or any other issue arises.
func (s *ReceiptServiceImpl) GetPoints(id string) (int, error) {
	receipt, err := s.ReceiptStore.Find(id)
	if err != nil {
		return 0, fmt.Errorf("failed to find receipt: %v", err)
	}

	points := receipt.Points

	return points, nil
}
