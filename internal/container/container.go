// Package container provides dependency injection for the receipt processing application.
// It initializes and wires together services, repositories, and handlers to ensure proper
// separation of concerns and modular design.
package container

import (
	adaptersHttp "go-receipt-processor/internal/adapters/http"
	"go-receipt-processor/internal/adapters/memory"
	"go-receipt-processor/internal/application"
	portsHttp "go-receipt-processor/internal/ports/http"
	"go-receipt-processor/internal/ports/repository"
)

// Container holds the application's dependencies, including services, repositories,
// and any other components needed to handle business logic and request processing.
// It centralizes the creation and management of these dependencies for ease of testing and configuration.
type Container struct {
	ReceiptService   portsHttp.ReceiptService   // Service for processing receipts.
	PointsCalculator portsHttp.PointsCalculator // Service for calculating points based on receipts.
	ReceiptStore     repository.ReceiptStore    // Repository for storing and retrieving receipt data.
}

// NewContainer initializes and returns a new Container instance.
// It wires together the services and repositories by instantiating their dependencies.
func NewContainer() *Container {
	return &Container{
		ReceiptService: application.NewReceiptService(
			application.NewPointsCalculator(), // Initializes the points calculator service.
			memory.NewReceiptStore(),          // Uses an in-memory implementation of the receipt store.
		),
	}
}

// NewReceiptProcessHandler creates and returns a new handler for processing receipts.
// This handler is responsible for handling the `POST /receipt/process` route and interacts
// with the ReceiptService to handle business logic.
func (c *Container) NewReceiptProcessHandler() *adaptersHttp.ReceiptProcessHandler {
	return adaptersHttp.NewReceiptProcessHandler(c.ReceiptService)
}
