package http

import "go-receipt-processor/internal/domain"

// ReceiptService defines the interface for processing receipts.
type ReceiptService interface {
    ProcessReceipt(receipt domain.Receipt) (string, int, error)
}
