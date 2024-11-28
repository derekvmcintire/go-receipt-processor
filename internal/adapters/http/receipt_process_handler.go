package http

// Importing necessary packages
import (
	"go-receipt-processor/internal/domain"                  //Package that holds the domain models (like Receipt)
	internalHttp "go-receipt-processor/internal/ports/http" // The service interface for receipt processing (from the ports/http package) - Alias this import to avoid conflict with the gin package
	netHttp "net/http"                                      // Standard Go package for HTTP-related constants (like StatusOK, StatusBadRequest) - Alias net/http to avoid name conflicts

	"github.com/gin-gonic/gin" //Web framework for building REST APIs in Go
	"github.com/google/uuid"   // Package to generate unique identifiers (UUIDs)
)

/*
 * ReceiptProcessHandler is a struct that handles receipt processing routes.
 * It contains the service (ReceiptService) which will perform the business logic
 * for processing the receipt and calculating points.
 */
type ReceiptProcessHandler struct {
	// The ReceiptService is an interface that defines the methods needed for processing receipts.
	// The concrete implementation of this interface is provided when the handler is created.
	ReceiptService internalHttp.ReceiptService
}

/*
 * NewReceiptProcessHandler is a factory function that creates a new instance of ReceiptProcessHandler.
 * This function accepts the ReceiptService interface as a parameter, which will be injected into the handler.
 * The handler uses this service to perform the business logic of processing the receipt.
 */
func NewReceiptProcessHandler(service internalHttp.ReceiptService) *ReceiptProcessHandler {
	// Return a new ReceiptProcessHandler instance with the given ReceiptService.
	return &ReceiptProcessHandler{ReceiptService: service}
}

/*
 * ProcessReceipt is the method that handles the HTTP POST request for processing a receipt.
 * It's tied to the route `/receipt/process` and processes incoming JSON data for a receipt.
 */
func (h *ReceiptProcessHandler) ProcessReceipt(c *gin.Context) {
	// Declare a variable to hold the receipt data from the request.
	var receipt domain.Receipt

	// Attempt to bind the incoming JSON data to the `receipt` variable.
	// This will automatically parse the JSON body into a `domain.Receipt` struct.
	// c.ShouldBindJSON returns nil on success
	if err := c.ShouldBindJSON(&receipt); err != nil {
		// If there was an error binding the JSON (e.g., invalid format),
		// return a BadRequest (400) response with the error message.
		c.JSON(netHttp.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate a unique ID for the receipt. This could be useful for tracking purposes.
	// UUID ensures that the receipt ID is globally unique.
	receipt.ID = uuid.New().String()

	// Call the `ProcessReceipt` method of the ReceiptService to process the receipt.
	// This service method is responsible for business logic like calculating points.
	receiptID, points, err := h.ReceiptService.ProcessReceipt(receipt)
	if err != nil {
		// If there is an error during the processing (e.g., business rule violation),
		// return an InternalServerError (500) response with the error message.
		c.JSON(netHttp.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// If processing is successful, return the generated receipt ID and the calculated points.
	// The response is sent with a status code of OK (200) and a JSON object containing the ID and points.
	c.JSON(netHttp.StatusOK, gin.H{
		"id":     receiptID, // The unique ID of the processed receipt
		"points": points,    // The calculated points for the receipt
	})
}
