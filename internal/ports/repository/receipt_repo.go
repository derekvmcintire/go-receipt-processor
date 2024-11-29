// Package repository defines interfaces for data storage operations.
// These interfaces abstract the underlying storage implementation, enabling
// flexibility and easier testing by allowing dependency injection of different
// storage backends (e.g., in-memory, database).
package repository

import "go-receipt-processor/internal/domain"

// ReceiptStore defines the methods required for storing and retrieving receipts.
// This interface allows different implementations, such as in-memory storage
// or database-backed storage, to be used interchangeably.
type ReceiptStore interface {
	// Save stores the given receipt and returns its unique ID or an error if the operation fails.
	Save(receipt domain.Receipt) (string, error)

	// Find retrieves a receipt by its unique ID. It returns the receipt's ID
	// and an error if the receipt is not found or another issue occurs.
	Find(id string) (domain.Receipt, error)
}
