package http_test

import (
	"fmt"
	adaptersHttp "go-receipt-processor/internal/adapters/http"
	"go-receipt-processor/internal/domain"
	externalHttp "net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
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
	return "1", nil
}

func TestGetPointsHandler_Success(t *testing.T) {
	// Arrange
	mockService := new(MockReceiptService)
	mockService.On("GetPoints", "123").Return(100, nil) // Mock a successful return of 100 points

	handler := adaptersHttp.NewGetReceiptPointsHandler(mockService)
	router := gin.Default()
	router.GET("/receipts/:id/points", handler.GetPoints)

	// Act
	req, err := externalHttp.NewRequest("GET", "/receipts/123/points", nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, externalHttp.StatusOK, w.Code)
	expectedResponse := `{"points":100}`
	assert.JSONEq(t, expectedResponse, w.Body.String())

	// Verify that the mock service method was called as expected
	mockService.AssertExpectations(t)
}

func TestGetPointsHandler_Error(t *testing.T) {
	// Arrange
	mockService := new(MockReceiptService)
	mockService.On("GetPoints", "123").Return(0, fmt.Errorf("receipt not found")) // Mock an error return

	handler := adaptersHttp.NewGetReceiptPointsHandler(mockService)
	router := gin.Default()
	router.GET("/receipts/:id/points", handler.GetPoints)

	// Act
	req, err := externalHttp.NewRequest("GET", "/receipts/123/points", nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, externalHttp.StatusInternalServerError, w.Code)
	expectedResponse := `{"error":"receipt not found"}`
	assert.JSONEq(t, expectedResponse, w.Body.String())

	// Verify that the mock service method was called as expected
	mockService.AssertExpectations(t)
}
