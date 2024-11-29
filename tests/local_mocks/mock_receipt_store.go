package local_mocks

import (
	"go-receipt-processor/internal/domain"

	"github.com/stretchr/testify/mock"
)

// Mock ReceiptStore
type MockReceiptStore struct {
	mock.Mock
}

func (m *MockReceiptStore) Save(receipt domain.Receipt) (string, error) {
	args := m.Called(receipt)
	return args.String(0), args.Error(1)
}

func (m *MockReceiptStore) Find(id string) (domain.Receipt, error) {
	args := m.Called(id)
	return args.Get(0).(domain.Receipt), args.Error(1)
}
