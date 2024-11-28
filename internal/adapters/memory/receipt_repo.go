// In-memory implementation of repository
package memory

import (
	"go-receipt-processor/internal/domain"
	"go-receipt-processor/internal/ports/repository"
)

// define the struct
type ReceiptRepositoryImpl struct{}

func NewReceiptRepo() repository.ReceiptRepo {
	return &ReceiptRepositoryImpl{}
}

func (r *ReceiptRepositoryImpl) Save(receipt domain.Receipt) (string, error) {
	return "1", nil
}

func (r *ReceiptRepositoryImpl) Find(id int) (int, error) {
	return 1, nil
}
