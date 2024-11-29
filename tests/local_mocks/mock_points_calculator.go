package local_mocks

import (
	"go-receipt-processor/internal/domain"

	"github.com/stretchr/testify/mock"
)

// Mock PointsCalculator
type MockPointsCalculator struct {
	mock.Mock
}

func (m *MockPointsCalculator) CalculatePoints(receipt domain.Receipt) (int, error) {
	args := m.Called(receipt)
	return args.Int(0), args.Error(1)
}
