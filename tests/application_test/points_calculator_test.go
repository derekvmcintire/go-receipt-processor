package application_test

import (
	"go-receipt-processor/internal/application"
	"go-receipt-processor/internal/domain"
	"go-receipt-processor/tests/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCalculatePoints_MockedHelpers(t *testing.T) {
	// Create a mock helper
	mockRules := new(mocks.MockPointsCalculatorRules)

	// Set up mock point return values
	mockRetailerNamePoints := 5
	mockRoundDollarPoints := 10
	mockMultipleOfQuarterPoints := 5
	mockItemCountPoints := 3
	mockItemDescriptionPoints := 2
	mockOddDayPoints := 0
	mockPurchaseTimePoints := 3

	// Set up mock behavior for each helper function
	mockRules.On("AddPointsForRetailerName", mock.Anything).Return(mockRetailerNamePoints)
	mockRules.On("AddPointsForRoundDollarTotal", mock.Anything).Return(mockRoundDollarPoints, nil)
	mockRules.On("AddPointsForMultipleOfQuarter", mock.Anything).Return(mockMultipleOfQuarterPoints, nil)
	mockRules.On("AddPointsForItemCount", mock.Anything).Return(mockItemCountPoints)
	mockRules.On("AddPointsForItemDescriptions", mock.Anything).Return(mockItemDescriptionPoints, nil)
	mockRules.On("AddPointsForOddDay", mock.Anything).Return(mockOddDayPoints)
	mockRules.On("AddPointsForAfternoonPurchaseTime", mock.Anything).Return(mockPurchaseTimePoints)

	// Set up expected points
	expectedPoints := 0
	expectedPoints += mockRetailerNamePoints
	expectedPoints += mockRoundDollarPoints
	expectedPoints += mockMultipleOfQuarterPoints
	expectedPoints += mockItemCountPoints
	expectedPoints += mockItemDescriptionPoints
	expectedPoints += mockOddDayPoints
	expectedPoints += mockPurchaseTimePoints

	// Create a new PointsCalculator instance with the mocked helpers
	calculator := application.NewPointsCalculator(mockRules)

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
	assert.Equal(t, expectedPoints, points) // 5 + 10 + 5 + 3 + 2 + 0 + 3 = 28

	// Assert that all expected mock methods were called
	mockRules.AssertExpectations(t)
}
