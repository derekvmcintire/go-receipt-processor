// mock_points_calculator_helpers.go
package mocks

import (
	"go-receipt-processor/internal/domain"
	"time"

	"github.com/stretchr/testify/mock"
)

// MockPointsCalculatorHelpers is a mock implementation of the PointsCalculatorHelpers interface.
type MockPointsCalculatorHelpers struct {
	mock.Mock
}

func (m *MockPointsCalculatorHelpers) AddPointsForRetailerName(receipt domain.Receipt) int {
	args := m.Called(receipt)
	return args.Int(0)
}

func (m *MockPointsCalculatorHelpers) AddPointsForRoundDollarTotal(receipt domain.Receipt) (int, error) {
	args := m.Called(receipt)
	return args.Int(0), args.Error(1)
}

func (m *MockPointsCalculatorHelpers) AddPointsForMultipleOfQuarter(receipt domain.Receipt) (int, error) {
	args := m.Called(receipt)
	return args.Int(0), args.Error(1)
}

func (m *MockPointsCalculatorHelpers) AddPointsForItemCount(receipt domain.Receipt) int {
	args := m.Called(receipt)
	return args.Int(0)
}

func (m *MockPointsCalculatorHelpers) AddPointsForItemDescriptions(receipt domain.Receipt) (int, error) {
	args := m.Called(receipt)
	return args.Int(0), args.Error(1)
}

func (m *MockPointsCalculatorHelpers) AddPointsForOddDay(parsedDateAndTime time.Time) int {
	args := m.Called(parsedDateAndTime)
	return args.Int(0)
}

func (m *MockPointsCalculatorHelpers) AddPointsForAfternoonPurchaseTime(parsedDateAndTime time.Time) int {
	args := m.Called(parsedDateAndTime)
	return args.Int(0)
}
