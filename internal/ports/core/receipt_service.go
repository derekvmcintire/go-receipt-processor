package http

import "go-receipt-processor/internal/domain"

// ReceiptService defines the interface for processing receipts and managing points.
type ReceiptService interface {
	ProcessReceipt(receipt domain.Receipt) (receiptID string, err error)
	GetPoints(id string) (points int, err error)
}
