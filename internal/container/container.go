package container

import (
	adaptersHttp "go-receipt-processor/internal/adapters/http" // Package for HTTP handlers interacting with the ReceiptService
	"go-receipt-processor/internal/adapters/memory"            // In-memory implementation of the ReceiptStore for testing and demo purposes
	"go-receipt-processor/internal/application"                // Contains core business logic for receipt processing and points calculation
	portsHttp "go-receipt-processor/internal/ports/http"       // HTTP service interfaces for the application, including receipt processing
)

// Container holds the application's dependencies, including services, repositories,
// and any other components needed to handle business logic and request processing.
// It centralizes the creation and management of these dependencies, enabling modularity,
// easy configuration, and flexibility, especially in testing scenarios.
type Container struct {
	ReceiptService portsHttp.ReceiptService // Service for processing receipts, calculating points, and interacting with repositories
}

// NewContainer initializes and returns a new Container instance.
//
// Returns:
//   - A new instance of Container with all dependencies initialized.
func NewContainer() *Container {
	return &Container{
		// Initialize the ReceiptService with a PointsCalculator and a ReceiptStore implementation.
		ReceiptService: application.NewReceiptService(
			application.NewPointsCalculator(application.NewPointsCalculatorHelper()), // Initializes the points calculator service.
			memory.NewReceiptStore(), // Uses an in-memory implementation of the receipt store (for testing or demo).
		),
	}
}

// NewReceiptProcessHandler creates and returns a new handler for processing receipts.
//
// Returns:
//   - A new instance of ReceiptProcessHandler, which can handle receipt processing requests.
func (c *Container) NewReceiptProcessHandler() *adaptersHttp.ReceiptProcessHandler {
	return adaptersHttp.NewReceiptProcessHandler(c.ReceiptService)
}

// NewGetReceiptPointsHandler creates and returns a new handler for retrieving receipt points.
//
// Returns:
//   - A new instance of GetReceiptPointsHandler, which can handle requests to get points for a specific receipt.
func (c *Container) NewGetReceiptPointsHandler() *adaptersHttp.GetReceiptPointsHandler {
	return adaptersHttp.NewGetReceiptPointsHandler(c.ReceiptService)
}
