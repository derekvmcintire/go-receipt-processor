package memory

import (
	"github.com/google/uuid"
	"go-receipt-processor/internal/domain"
	"go-receipt-processor/internal/ports/repository"
	"sync"
)

// ReceiptStoreImpl stores receipts in an in-memory map, keyed by unique IDs.
type ReceiptStoreImpl struct {
	receipts map[string]domain.Receipt
}

// Declare a private variable to hold the singleton instance
var instance *ReceiptStoreImpl
var once sync.Once

// NewReceiptStore returns the singleton instance of ReceiptStoreImpl
func NewReceiptStore() repository.ReceiptStore {
	once.Do(func() {
		// Only create the instance once
		instance = &ReceiptStoreImpl{
			receipts: make(map[string]domain.Receipt),
		}
	})
	return instance
}

// Save stores a receipt in memory and returns its ID
func (r *ReceiptStoreImpl) Save(receipt domain.Receipt) (string, error) {
	receiptID := uuid.New().String()
	r.receipts[receiptID] = receipt
	return receiptID, nil
}

// Find retrieves a receipt by ID
func (r *ReceiptStoreImpl) Find(id string) (domain.Receipt, error) {
	foundReceipt := r.receipts[id]
	return foundReceipt, nil
}
