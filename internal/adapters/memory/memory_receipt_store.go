package memory

import (
	"github.com/google/uuid"
	"go-receipt-processor/internal/domain"
	"go-receipt-processor/internal/ports/repository"
)

// ReceiptStoreImpl stores receipts in an in-memory map, keyed by unique IDs.
type ReceiptStoreImpl struct {
	receipts map[string]domain.Receipt
}

// NewReceiptStore
//
// Returns:
//   - A new instance of repository.ReceiptStore (which is *ReceiptStoreImpl).
func NewReceiptStore() repository.ReceiptStore {
	return &ReceiptStoreImpl{
		receipts: make(map[string]domain.Receipt),
	}
}

// Save
//
// Parameters:
//   - receipt: The domain.Receipt object to be stored.
//
// Returns:
//   - receiptID: The unique identifier generated for the receipt.
//   - err: An error, if any. The current implementation does not produce errors.
func (r *ReceiptStoreImpl) Save(receipt domain.Receipt) (string, error) {
	receiptID := uuid.New().String()
	r.receipts[receiptID] = receipt
	return receiptID, nil
}

// Find
//
// Parameters:
//   - id: The unique identifier of the receipt to be retrieved.
//
// Returns:
//   - The domain.Receipt associated with the provided ID.
//   - An error, if any. Currently, no errors are returned, but this may be adjusted in the future.
func (r *ReceiptStoreImpl) Find(id string) (domain.Receipt, error) {
	foundReceipt := r.receipts[id]
	return foundReceipt, nil
}
