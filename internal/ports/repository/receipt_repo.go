// Interface for storing receipts (e.g., database or in-memory)
package repository

import "go-receipt-processor/internal/domain"

// ReceiptStore defines the interface for storing receipts in a database or memory
type ReceiptStore interface {
	Save(receipt domain.Receipt) (string, error)
	Find(id int) (int, error)
}
