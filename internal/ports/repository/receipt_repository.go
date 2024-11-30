package repository

import "go-receipt-processor/internal/domain"

// ReceiptStore defines the methods required for storing and retrieving receipts.
type ReceiptStore interface {
	Save(receipt domain.Receipt) (receiptID string, err error)
	Find(id string) (receipt domain.Receipt, err error)
}
