package application

// Importing necessary packages
// domain: Contains the domain models, like `Receipt` and `Item`.
// http: The `http` package defines the interface for the receipt service.
// math: Provides mathematical functions, like `Floor` and `Ceil`.
// strings: Contains functions for string manipulation, such as `ReplaceAll` and `TrimSpace`.
// uuid: Used to generate unique identifiers (UUIDs).
import (
	"github.com/google/uuid"
	"go-receipt-processor/internal/domain"
	"go-receipt-processor/internal/ports/http"
	"math"
	"strings"
)

// ReceiptServiceImpl is the concrete implementation of the ReceiptService interface.
// It contains the logic for processing a receipt and calculating points based on various rules.
type ReceiptServiceImpl struct{}

// NewReceiptService creates a new instance of ReceiptServiceImpl.
// It is a factory function that returns a pointer to a new ReceiptServiceImpl object.
// This is used in the container to create the service and inject it where needed.
func NewReceiptService() http.ReceiptService {
	return &ReceiptServiceImpl{}
}

// ProcessReceipt is the method that processes a receipt and calculates points based on certain rules.
// It receives a `domain.Receipt` as input, which contains details like the retailer, total, items, and purchase info.
// It returns a unique receipt ID, the calculated points, and any errors that occur during processing.
func (s *ReceiptServiceImpl) ProcessReceipt(receipt domain.Receipt) (string, int, error) {
	points := 0 // Initialize points to 0. This will accumulate points based on various rules.

	// Rule 1: One point for each alphanumeric character in the retailer name.
	// First, we remove any spaces from the retailer name using `strings.ReplaceAll`.
	// Then, we count the length of the remaining string and add that to the points.
	points += len(strings.ReplaceAll(receipt.Retailer, " ", ""))

	// Rule 2: 50 points if the total amount is a round dollar value (no cents).
	// This is checked by comparing the total to its floor value (i.e., the value without decimals).
	if receipt.Total == math.Floor(receipt.Total) {
		points += 50
	}

	// Rule 3: 25 points if the total is a multiple of 0.25.
	// We multiply the total by 100 to check if it's a multiple of 25 (i.e., total in cents).
	if int(receipt.Total*100)%25 == 0 {
		points += 25
	}

	// Rule 4: 5 points for every two items on the receipt.
	// The number of items is divided by 2, and the result is multiplied by 5 to calculate the points.
	points += (len(receipt.Items) / 2) * 5

	// Rule 5: For items with descriptions that are a multiple of 3 in length, apply the price multiplier rule.
	// We check if the length of the description (after trimming spaces) is a multiple of 3.
	// If true, we add points based on the price of the item, applying a 20% multiplier.
	for _, item := range receipt.Items {
		if len(strings.TrimSpace(item.ShortDescription))%3 == 0 {
			points += int(math.Ceil(item.Price * 0.2)) // Use math.Ceil to round up the points.
		}
	}

	// Rule 6: 6 points if the purchase day is odd (i.e., 1st, 3rd, 5th, etc. day of the month).
	// We check if the day of the purchase (from `receipt.PurchaseDate`) is an odd number.
	if receipt.PurchaseDate.Day()%2 != 0 {
		points += 6
	}

	// Rule 7: 10 points if the purchase time is between 2:00pm and 4:00pm.
	// We check the hour of the `PurchaseTime` to see if it falls within the 2-4pm range.
	if receipt.PurchaseTime.Hour() >= 14 && receipt.PurchaseTime.Hour() < 16 {
		points += 10
	}

	// Generate a unique ID for the receipt. This is done using the `uuid.New().String()` method.
	// This ID could be used for tracking the receipt or saving it in a database.
	receiptID := uuid.New().String()

	// Return the generated receipt ID, the calculated points, and nil (since there is no error in this case).
	return receiptID, points, nil // Make sure the return matches the method signature from the ReceiptService interface.
}
