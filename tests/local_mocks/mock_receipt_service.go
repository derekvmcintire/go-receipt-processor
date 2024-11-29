package local_mocks

import (
	"go-receipt-processor/internal/domain"

	"github.com/stretchr/testify/mock"
)

// MockReceiptService is a mock of the ReceiptService interface for unit testing
type MockReceiptService struct {
	mock.Mock
}

func (m *MockReceiptService) GetPoints(receiptID string) (int, error) {
	args := m.Called(receiptID)
	return args.Int(0), args.Error(1)
}

func (m *MockReceiptService) ProcessReceipt(receipt domain.Receipt) (string, error) {
	args := m.Called(receipt)
	return args.String(0), args.Error(1)
}
