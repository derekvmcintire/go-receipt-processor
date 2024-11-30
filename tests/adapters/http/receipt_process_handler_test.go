package http_test

import (
	"fmt"
	adaptersHttp "go-receipt-processor/internal/adapters/http"
	"go-receipt-processor/internal/domain"
	"go-receipt-processor/tests/local_mocks"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestProcessReceipt_Success(t *testing.T) {
	// Arrange: Create a mock service and set expectations for a successful response
	mockService := new(local_mocks.MockReceiptService)
	mockService.On("ProcessReceipt", mock.Anything).Return("receipt123", nil) // Mock successful receipt processing

	// Create the handler and the router
	handler := adaptersHttp.NewReceiptProcessHandler(mockService)
	router := gin.Default()
	router.POST("/receipt/process", handler.ProcessReceipt)

	// Act: Send a valid JSON POST request to process the receipt
	receipt := domain.Receipt{
		Retailer:     "Store A",
		PurchaseDate: "2024-11-29",
		PurchaseTime: "14:30",
		Items: []domain.Item{
			{ShortDescription: "Item 1", Price: "50.00"},
			{ShortDescription: "Item 2", Price: "50.00"},
		},
		Total:  "100.00",
		Points: 100,
	}

	// Generate the request body for the receipt as JSON
	reqBody := fmt.Sprintf(`{
		"retailer": "%s",
		"purchaseDate": "%s",
		"purchaseTime": "%s",
		"items": [{"shortDescription": "%s", "price": "%s"}, {"shortDescription": "%s", "price": "%s"}],
		"total": "%s",
		"points": %d
	}`, receipt.Retailer, receipt.PurchaseDate, receipt.PurchaseTime,
		receipt.Items[0].ShortDescription, receipt.Items[0].Price,
		receipt.Items[1].ShortDescription, receipt.Items[1].Price,
		receipt.Total, receipt.Points)

	// Send the HTTP request
	req, err := http.NewRequest("POST", "/receipt/process", strings.NewReader(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert: Verify that the response status and body are correct
	assert.Equal(t, http.StatusOK, w.Code) // Should return 200 OK
	expectedResponse := `{"id":"receipt123"}`
	assert.JSONEq(t, expectedResponse, w.Body.String()) // Check that the receipt ID is returned correctly

	// Verify that the mock service method was called as expected
	mockService.AssertExpectations(t)
}

func TestProcessReceipt_InvalidJSON(t *testing.T) {
	// Arrange: Create a mock service (it won't be called because of invalid JSON)
	mockService := new(local_mocks.MockReceiptService)

	// Create the handler and the router
	handler := adaptersHttp.NewReceiptProcessHandler(mockService)
	router := gin.Default()
	router.POST("/receipt/process", handler.ProcessReceipt)

	// Act: Send an invalid JSON POST request (missing closing quote for "store")
	invalidJSON := `{"retailer": "Store A,""purchaseDate": "2024-11-29","purchaseTime":"14:30","items":[{"shortDescription":"Item 1","price":"50.00"},{"shortDescription":"Item 2","price":"50.00"}],"total":"100.00","points":100}`
	req, err := http.NewRequest("POST", "/receipt/process", strings.NewReader(invalidJSON))
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert: Verify that the response status is 400 Bad Request for invalid JSON
	assert.Equal(t, http.StatusBadRequest, w.Code) // Should return 400 Bad Request
	expectedResponse := `{"details":"invalid character '\"' after object key:value pair", "error":"Invalid request payload"}`
	assert.JSONEq(t, expectedResponse, w.Body.String()) // Check for proper error message

	// Verify that the mock service was not called
	mockService.AssertExpectations(t)
}

func TestProcessReceipt_ServiceError(t *testing.T) {
	// Arrange: Create a mock service and set expectations for an error case
	mockService := new(local_mocks.MockReceiptService)
	mockService.On("ProcessReceipt", mock.Anything).Return("", fmt.Errorf("processing failed")) // Mock a failure

	// Create the handler and the router
	handler := adaptersHttp.NewReceiptProcessHandler(mockService)
	router := gin.Default()
	router.POST("/receipt/process", handler.ProcessReceipt)

	// Act: Send a valid JSON POST request to process the receipt
	receipt := domain.Receipt{
		Retailer:     "Store A",
		PurchaseDate: "2024-11-29",
		PurchaseTime: "14:30",
		Items: []domain.Item{
			{ShortDescription: "Item 1", Price: "50.00"},
			{ShortDescription: "Item 2", Price: "50.00"},
		},
		Total:  "100.00",
		Points: 100,
	}

	// Generate the request body for the receipt as JSON
	reqBody := fmt.Sprintf(`{
		"retailer": "%s",
		"purchaseDate": "%s",
		"purchaseTime": "%s",
		"items": [{"shortDescription": "%s", "price": "%s"}, {"shortDescription": "%s", "price": "%s"}],
		"total": "%s",
		"points": %d
	}`, receipt.Retailer, receipt.PurchaseDate, receipt.PurchaseTime,
		receipt.Items[0].ShortDescription, receipt.Items[0].Price,
		receipt.Items[1].ShortDescription, receipt.Items[1].Price,
		receipt.Total, receipt.Points)

	// Send the HTTP request
	req, err := http.NewRequest("POST", "/receipt/process", strings.NewReader(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert: Verify that the response status and body are correct for service error
	assert.Equal(t, http.StatusInternalServerError, w.Code) // Should return 500 Internal Server Error
	expectedResponse := `{"error":"processing failed"}`
	assert.JSONEq(t, expectedResponse, w.Body.String()) // Check for error message

	// Verify that the mock service method was called as expected
	mockService.AssertExpectations(t)
}
