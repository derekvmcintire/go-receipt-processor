package memory

import (
	"github.com/google/uuid"
	"go-receipt-processor/internal/domain"
	"go-receipt-processor/internal/ports/repository"
)

// ReceiptStoreImpl is an in-memory implementation of the repository.ReceiptStore interface.
// It stores receipts in a map, keyed by unique IDs generated using UUIDs.
type ReceiptStoreImpl struct {
	receipts map[string]domain.Receipt // In-memory map of receipts, keyed by their unique ID.
}

// NewReceiptStore creates and returns a new in-memory ReceiptStoreImpl instance.
// The receipts map is initialized as an empty map.
//
// Returns:
//   - A new instance of repository.ReceiptStore (which is *ReceiptStoreImpl).
func NewReceiptStore() repository.ReceiptStore {
	return &ReceiptStoreImpl{
		receipts: make(map[string]domain.Receipt),
	}
}

// Save stores a receipt in the in-memory store and generates a unique ID for it.
//
// Parameters:
//   - receipt: The domain.Receipt object to be stored.
//
// Returns:
//   - receiptID: The unique identifier generated for the receipt.
//   - err: An error, if any. The current implementation does not produce errors.
func (r *ReceiptStoreImpl) Save(receipt domain.Receipt) (string, error) {
	// Generate a unique ID for the receipt.
	receiptID := uuid.New().String()
	// Save the receipt in the map using the generated ID as the key.
	r.receipts[receiptID] = receipt
	return receiptID, nil
}

// Find retrieves a receipt from the in-memory store using the provided unique ID.
//
// Parameters:
//   - id: The unique identifier of the receipt to be retrieved.
//
// Returns:
//   - The domain.Receipt associated with the provided ID.
//   - An error, if any. Currently, no errors are returned, but this may be adjusted in the future.
func (r *ReceiptStoreImpl) Find(id string) (domain.Receipt, error) {
	// Retrieve the receipt from the map using the provided ID.
	foundReceipt := r.receipts[id]
	return foundReceipt, nil
}
