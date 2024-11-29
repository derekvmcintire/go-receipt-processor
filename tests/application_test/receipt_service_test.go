package application_test

import (
	"fmt"
	"go-receipt-processor/internal/application"
	"go-receipt-processor/internal/domain"
	"go-receipt-processor/tests/local_mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestReceiptService_ProcessReceipt(t *testing.T) {
	// Create mock objects
	mockPointsCalculator := new(local_mocks.MockPointsCalculator)
	mockReceiptStore := new(local_mocks.MockReceiptStore)

	// Create the ReceiptService instance with the mocked dependencies
	receiptService := application.NewReceiptService(mockPointsCalculator, mockReceiptStore)

	// Make a copy of the MockReceipt to avoid modifying the global value
	receipt := local_mocks.MockReceipt
	receipt.Points = 50 // Set the expected points value after calculation

	// Mock behavior for CalculatePoints to return 50 points
	mockPointsCalculator.On("CalculatePoints", receipt).Return(50, nil)

	// Mock behavior for Save to expect the receipt with Points set to 50
	mockReceiptStore.On("Save", mock.MatchedBy(func(r domain.Receipt) bool {
		// Ensure the Points are 50 when saving
		return r.Points == 50
	})).Return("12345", nil)

	// Call ProcessReceipt method
	receiptID, err := receiptService.ProcessReceipt(receipt)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, "12345", receiptID)
	mockPointsCalculator.AssertExpectations(t)
	mockReceiptStore.AssertExpectations(t)
}

func TestReceiptService_ProcessReceipt_ErrorInPointsCalculation(t *testing.T) {
	// Create mock objects
	mockPointsCalculator := new(local_mocks.MockPointsCalculator)
	mockReceiptStore := new(local_mocks.MockReceiptStore)

	// Create the ReceiptService instance with the mocked dependencies
	receiptService := application.NewReceiptService(mockPointsCalculator, mockReceiptStore)

	// Make a copy of the MockReceipt to avoid modifying the global value
	receipt := local_mocks.MockReceipt

	// Mock behavior for CalculatePoints to return an error
	mockPointsCalculator.On("CalculatePoints", receipt).Return(0, fmt.Errorf("invalid purchase time format"))

	// Call ProcessReceipt method
	receiptID, err := receiptService.ProcessReceipt(receipt)

	// Assertions
	assert.Error(t, err)
	assert.Equal(t, "", receiptID)
	mockPointsCalculator.AssertExpectations(t)
	mockReceiptStore.AssertExpectations(t)
}

func TestReceiptService_GetPoints(t *testing.T) {
	// Create mock objects
	mockPointsCalculator := new(local_mocks.MockPointsCalculator)
	mockReceiptStore := new(local_mocks.MockReceiptStore)

	// Create the ReceiptService instance with the mocked dependencies
	receiptService := application.NewReceiptService(mockPointsCalculator, mockReceiptStore)

	// Prepare the test data
	receipt := local_mocks.MockReceipt
	receipt.Points = 50 // Set the expected points for the receipt

	// Mock the behavior of Find method to return the receipt
	mockReceiptStore.On("Find", "12345").Return(receipt, nil)

	// Call the GetPoints method
	points, err := receiptService.GetPoints("12345")

	// Assertions
	assert.NoError(t, err)      // No error should be returned
	assert.Equal(t, 50, points) // Points should be 50
	mockReceiptStore.AssertExpectations(t)
}

func TestReceiptService_GetPoints_ReceiptNotFound(t *testing.T) {
	// Create mock objects
	mockPointsCalculator := new(local_mocks.MockPointsCalculator)
	mockReceiptStore := new(local_mocks.MockReceiptStore)

	// Create the ReceiptService instance with the mocked dependencies
	receiptService := application.NewReceiptService(mockPointsCalculator, mockReceiptStore)

	// Mock the behavior of Find method to return an error (receipt not found)
	mockReceiptStore.On("Find", "12345").Return(domain.Receipt{}, fmt.Errorf("receipt not found"))

	// Call the GetPoints method
	points, err := receiptService.GetPoints("12345")

	// Assertions
	assert.Error(t, err)       // Error should be returned
	assert.Equal(t, 0, points) // Points should be 0
	mockReceiptStore.AssertExpectations(t)
}
