package repository

import "go-receipt-processor/internal/domain"

// ReceiptStore defines the methods required for storing and retrieving receipts.
// Implementations of this interface provide the ability to persist and fetch receipt data,
// supporting various storage backends such as in-memory or database-based solutions.
type ReceiptStore interface {
	// Save stores the given receipt and returns:
	//   - The unique ID of the saved receipt as a string.
	//   - An error if the receipt could not be saved (e.g., due to storage issues).
	Save(receipt domain.Receipt) (receiptID string, err error)

	// Find retrieves a receipt by its unique ID and returns:
	//   - The receipt corresponding to the given ID.
	//   - An error if the receipt is not found or another issue occurs (e.g., database failure).
	Find(id string) (receipt domain.Receipt, err error)
}
