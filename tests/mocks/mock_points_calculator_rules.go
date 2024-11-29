// mock_points_calculator_helpers.go
package mocks

import (
	"go-receipt-processor/internal/domain"
	"time"

	"github.com/stretchr/testify/mock"
)

// MockPointsCalculatorRules is a mock implementation of the PointsCalculatorHelpers interface.
type MockPointsCalculatorRules struct {
	mock.Mock
}

func (m *MockPointsCalculatorRules) AddPointsForRetailerName(receipt domain.Receipt) int {
	args := m.Called(receipt)
	return args.Int(0)
}

func (m *MockPointsCalculatorRules) AddPointsForRoundDollarTotal(receipt domain.Receipt) (int, error) {
	args := m.Called(receipt)
	return args.Int(0), args.Error(1)
}

func (m *MockPointsCalculatorRules) AddPointsForMultipleOfQuarter(receipt domain.Receipt) (int, error) {
	args := m.Called(receipt)
	return args.Int(0), args.Error(1)
}

func (m *MockPointsCalculatorRules) AddPointsForItemCount(receipt domain.Receipt) int {
	args := m.Called(receipt)
	return args.Int(0)
}

func (m *MockPointsCalculatorRules) AddPointsForItemDescriptions(receipt domain.Receipt) (int, error) {
	args := m.Called(receipt)
	return args.Int(0), args.Error(1)
}

func (m *MockPointsCalculatorRules) AddPointsForOddDay(parsedDateAndTime time.Time) int {
	args := m.Called(parsedDateAndTime)
	return args.Int(0)
}

func (m *MockPointsCalculatorRules) AddPointsForAfternoonPurchaseTime(parsedDateAndTime time.Time) int {
	args := m.Called(parsedDateAndTime)
	return args.Int(0)
}
