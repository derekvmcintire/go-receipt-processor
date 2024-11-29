// In-memory implementation of repository
package memory

import (
	"go-receipt-processor/internal/domain"
	"go-receipt-processor/internal/ports/repository"

	"github.com/google/uuid"
)

// define the struct
type ReceiptStoreImpl struct {
	receipts map[string]domain.Receipt
}

func NewReceiptRepo() repository.ReceiptStore {
	return &ReceiptStoreImpl{
		receipts: make(map[string]domain.Receipt),
	}
}

func (r *ReceiptStoreImpl) Save(receipt domain.Receipt) (string, error) {
	// Generate a unique ID for the receipt. This is done using the `uuid.New().String()` method.
	// This ID could be used for tracking the receipt or saving it in a database.
	receiptID := uuid.New().String()
	r.receipts[receiptID] = receipt
	return receiptID, nil
}

func (r *ReceiptStoreImpl) Find(id int) (int, error) {
	return 1, nil
}
