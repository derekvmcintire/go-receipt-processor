package container

// Importing necessary packages
// adaptersHttp: the HTTP-related handlers in the `adapters/http` package
// application: the business logic layer (the `application` package)
// portsHttp: defines the ports (interfaces) for services like `ReceiptService`
import (
	adaptersHttp "go-receipt-processor/internal/adapters/http"
	"go-receipt-processor/internal/application"
	portsHttp "go-receipt-processor/internal/ports/http"
)

// The Container struct is used to hold instances of services like `ReceiptService`
// and provide them to the various parts of the application (handlers, etc.)
// The Container is responsible for the creation and wiring of dependencies.
type Container struct {
	// ReceiptService is the service responsible for processing receipts.
	// It's an interface defined in the `portsHttp` package, and it is implemented
	// by a concrete service (e.g., `ReceiptServiceImpl` in the `application` package).
	ReceiptService   portsHttp.ReceiptService
	PointsCalculator portsHttp.PointsCalculator
}

// NewContainer is a function that creates and returns a new instance of the Container.
// It initializes the services (e.g., `ReceiptService`) by calling the factory functions
// from the `application` package.
func NewContainer() *Container {
	// Here we're creating a new instance of the Container and initializing the
	// ReceiptService using a factory function from the `application` package.
	// The factory function `application.NewReceiptService()` creates a concrete
	// instance of the service (e.g., `ReceiptServiceImpl`).
	return &Container{
		ReceiptService: application.NewReceiptService(application.NewPointsCalculator()),
	}
}

// NewReceiptProcessHandler is a method on the Container that creates a new handler
// for the `POST /receipt/process` route. This handler is responsible for processing
// the receipt and interacting with the `ReceiptService` to handle the business logic.
func (c *Container) NewReceiptProcessHandler() *adaptersHttp.ReceiptProcessHandler {
	// This method uses the `NewReceiptProcessHandler` function from the `adaptersHttp` package
	// to create a new `ReceiptProcessHandler`. It injects the `ReceiptService` dependency
	// into the handler when creating it.
	//
	// `ReceiptProcessHandler` is the handler that will handle the actual HTTP request
	// for processing a receipt. This is where the request data is passed to the service layer.
	return adaptersHttp.NewReceiptProcessHandler(c.ReceiptService)
}
