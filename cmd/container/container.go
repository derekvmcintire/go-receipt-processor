package container

import (
	adaptersHttp "go-receipt-processor/internal/adapters/http"
	"go-receipt-processor/internal/adapters/memory"
	"go-receipt-processor/internal/application"
	portsHttp "go-receipt-processor/internal/ports/core"
)

// Container holds the application's dependencies
type Container struct {
	ReceiptService portsHttp.ReceiptService
}

// NewContainer
//
// Returns:
//   - A new instance of Container with all dependencies initialized.
func NewContainer() *Container {
	return &Container{
		ReceiptService: application.NewReceiptService(
			application.NewPointsCalculator(application.NewPointsCalculatorHelper()),
			memory.NewReceiptStore(),
		),
	}
}

// NewReceiptProcessHandler
//
// Returns:
//   - A new instance of ReceiptProcessHandler, which can handle receipt processing requests.
func (c *Container) NewReceiptProcessHandler() *adaptersHttp.ReceiptProcessHandler {
	return adaptersHttp.NewReceiptProcessHandler(c.ReceiptService)
}

// NewGetReceiptPointsHandler
//
// Returns:
//   - A new instance of GetReceiptPointsHandler, which can handle requests to get points for a specific receipt.
func (c *Container) NewGetReceiptPointsHandler() *adaptersHttp.GetReceiptPointsHandler {
	return adaptersHttp.NewGetReceiptPointsHandler(c.ReceiptService)
}
