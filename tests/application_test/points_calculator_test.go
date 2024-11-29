// points_calculator_impl_test.go
package application_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go-receipt-processor/internal/application"
	"go-receipt-processor/internal/domain"
	"go-receipt-processor/tests/mocks"
	"testing"
)

func TestCalculatePoints_MockedHelpers(t *testing.T) {
	// Create a mock helper
	mockHelper := new(mocks.MockPointsCalculatorHelpers)

	// Set up mock behavior for each helper function
	mockHelper.On("AddPointsForRetailerName", mock.Anything).Return(5)
	mockHelper.On("AddPointsForRoundDollarTotal", mock.Anything).Return(10, nil)
	mockHelper.On("AddPointsForMultipleOfQuarter", mock.Anything).Return(5, nil)
	mockHelper.On("AddPointsForItemCount", mock.Anything).Return(3)
	mockHelper.On("AddPointsForItemDescriptions", mock.Anything).Return(2, nil)
	mockHelper.On("AddPointsForOddDay", mock.Anything).Return(0)
	mockHelper.On("AddPointsForAfternoonPurchaseTime", mock.Anything).Return(3)

	// Create a new PointsCalculator instance with the mocked helpers
	calculator := application.NewPointsCalculator(mockHelper)

	// Create a sample receipt
	receipt := domain.Receipt{
		Retailer:     "StoreABC",
		PurchaseDate: "2024-11-29",
		PurchaseTime: "15:30",
		Items: []domain.Item{
			{ShortDescription: "Item 1", Price: "5.00"},
			{ShortDescription: "Item 2", Price: "5.00"},
		},
		Total:  "10.00",
		Points: 0,
	}

	// Calculate points
	points, err := calculator.CalculatePoints(receipt)

	// Assert the points and no error
	assert.NoError(t, err)
	assert.Equal(t, 28, points) // 5 + 10 + 5 + 3 + 2 + 0 + 3 = 28

	// Assert that all expected mock methods were called
	mockHelper.AssertExpectations(t)
}
