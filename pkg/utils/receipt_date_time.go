package utils

import (
	"fmt"
	"go-receipt-processor/internal/domain"
	"time"
)

// ParseReceiptDateTime parses the purchase date and time from the receipt
// Note that time format must use a very specific date:
// Mon Jan 2 15:04:05 MST 2006 - so using "15:04" here is
// just a format and not a value, but using "15:03" would break this lol
//
// Parameters:
//   - receipt: The domain.Receipt object containing receipt details.
//
// Returns: A time.TIme value representing the the exact time of the purchase
//   - combinedTime: A unique identifier for the processed receipt.
//   - err: An error if parsing date or time fails
func ParseReceiptDateTime(receipt domain.Receipt) (time.Time, error) {
	// Parse the purchase date using the "YYYY-MM-DD" format
	parsedDate, err := time.Parse("2006-01-02", receipt.PurchaseDate)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid purchase date format: %v", err)
	}

	// Parse the purchase time using the "HH:MM" format
	parsedTime, err := time.Parse("15:04", receipt.PurchaseTime)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid purchase time format: %v", err)
	}

	// Combine the parsed date and time into a single time.Time value.
	// Note: We replace the time components of parsedDate with the time from parsedTime
	combinedTime := time.Date(parsedDate.Year(), parsedDate.Month(), parsedDate.Day(),
		parsedTime.Hour(), parsedTime.Minute(), 0, 0, parsedDate.Location())

	// Return the combined time
	return combinedTime, nil
}
