package application

import (
	"fmt"
	"go-receipt-processor/internal/domain"           // Package containing domain models like Receipt
	http "go-receipt-processor/internal/ports/core"  // Package containing the PointsCalculator interface for calculating receipt points
	"go-receipt-processor/internal/ports/repository" // Package containing the ReceiptStore interface for saving and retrieving receipts
)

// ReceiptServiceImpl is an implementation of the ReceiptService interface.
type ReceiptServiceImpl struct {
	PointsCalculator http.PointsCalculator   // Interface for calculating points based on receipt data
	ReceiptStore     repository.ReceiptStore // Interface for managing storage and retrieval of receipts
}

// NewReceiptService creates and returns a new instance of ReceiptServiceImpl.
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

// ProcessReceipt processes a receipt by calculating points and storing it in the repository.
//
// Parameters:
//   - receipt: The domain.Receipt object containing receipt details.
//
// Returns:
//   - receiptID: A unique identifier for the processed receipt.
//   - err: An error if processing or saving fails.
func (s *ReceiptServiceImpl) ProcessReceipt(receipt domain.Receipt) (string, error) {
	// Calculate points using the PointsCalculator.
	points, err := s.PointsCalculator.CalculatePoints(receipt)
	if err != nil {
		// If an error occurs during points calculation, return an error with a message.
		return "", fmt.Errorf("unable to process receipt: %v", err)
	}

	// Assign the calculated points to the receipt.
	receipt.Points = points

	// Save the receipt to the repository and retrieve its unique ID.
	receiptID, err := s.ReceiptStore.Save(receipt)
	if err != nil {
		// If saving the receipt fails, return an error.
		return "", fmt.Errorf("failed to insert receipt: %v", err)
	}

	// Return the unique receipt ID and no error if everything is successful.
	return receiptID, nil
}

// GetPoints retrieves the points associated with a receipt by its unique ID.
//
// Parameters:
//   - id: The unique ID of the receipt whose points are being retrieved.
//
// Returns:
//   - points: The points associated with the receipt.
//   - err: An error if the receipt cannot be found or any other issue arises.
func (s *ReceiptServiceImpl) GetPoints(id string) (int, error) {
	// Retrieve the receipt from the repository using the provided ID.
	receipt, err := s.ReceiptStore.Find(id)
	if err != nil {
		// If the receipt cannot be found, return an error.
		return 0, fmt.Errorf("failed to find receipt: %v", err)
	}

	// Return the points associated with the found receipt.
	points := receipt.Points

	return points, nil
}
