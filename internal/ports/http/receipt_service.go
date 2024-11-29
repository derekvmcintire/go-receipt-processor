package http

import "go-receipt-processor/internal/domain"

// ReceiptService defines the interface for processing receipts and managing points.
// Implementations of this interface handle receipt processing, point calculation,
// and retrieval of points for a specific receipt.
type ReceiptService interface {
	// ProcessReceipt processes the given receipt and calculates the associated points.
	// It returns:
	//   - The unique ID of the receipt as a string.
	//   - An error if the receipt processing fails (e.g., invalid data).
	ProcessReceipt(receipt domain.Receipt) (receiptID string, err error)

	// GetPoints retrieves the points associated with the given receipt ID.
	// It returns:
	//   - The number of points as an integer.
	//   - An error if the retrieval fails (e.g., receipt ID not found).
	GetPoints(id string) (points int, err error)
}
