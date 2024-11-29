// Package http defines the ReceiptService interface for processing receipts.
package http

import "go-receipt-processor/internal/domain"

// ReceiptService defines the interface for processing receipts.
// It includes a method to process a receipt and return its ID, points, and any errors.
type ReceiptService interface {
	// ProcessReceipt processes a receipt and calculates points.
	// It returns the receipt ID, the calculated points, and any error encountered.
	ProcessReceipt(receipt domain.Receipt) (string, int, error)
}
